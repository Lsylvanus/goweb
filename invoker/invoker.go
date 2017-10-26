package invoker

import (
	"goweb/common"
	"goweb/models"
	"goweb/policy"

	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// 执行
type Invoker struct {
}

// 执行任务
func (this *Invoker) Execute(jobInfo *models.JobInfo, nextTime time.Time, params string) error {
	snapshot, err := this.Init(jobInfo, nextTime)
	log.Println("snapshot = ", snapshot)
	fac := &policy.Factory{}
	np := fac.FindPolicy(jobInfo)
	count := 0
	for {

		url := np.GetNextUrl()

		if url == "" {
			this.executeJobResult(snapshot, "所有目标服务器地址都不可用"+time.Now().Local().Format("2006-01-02 15:04:05"), models.ERROR)
			break
		}
		snapshot.Url = url
		// 准备执行

		err = this.invoke(snapshot)
		if err == nil {
			break
		} else if count < 3 {

			this.executeJobResult(snapshot, "目标服务器地址:"+url+"不可用正在尝试"+time.Now().Local().Format("2006-01-02 15:04:05"), models.ERROR)
		}
		count++

	}
	return err
}

// 执行任务
func (this *Invoker) invoke(jobSnapshot *models.JobSnapshot) error {

	this.executeJobResult(jobSnapshot, "准备任务提交至目标服务器地址:"+jobSnapshot.Url+time.Now().Local().Format("2006-01-02 15:04:05"), models.INVOKING)
	startTime := time.Now()

	jobRequest := &common.JobRequest{JobSnapshot: jobSnapshot.Id, Params: jobSnapshot.Params, Status: models.INVOKING}

	content, err := json.Marshal(jobRequest)

	if err != nil {
		this.executeJobResult(jobSnapshot, "解析job请求参数出错"+time.Now().Local().Format("2006-01-02 15:04:05"), models.ERROR)
		return err
	}

	resp, err := http.Post(jobSnapshot.Url, "application/json;charset=utf-8", bytes.NewBuffer(content))
	if err != nil {
		this.executeJobResult(jobSnapshot, "目标服务器地址:"+jobSnapshot.Url+"不可用"+time.Now().Local().Format("2006-01-02 15:04:05"), models.ERROR)
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	result := &common.JobResponse{}
	log.Println("body = ", string(body))
	err = json.Unmarshal(body, result)

	if err != nil {
		this.executeJobResult(jobSnapshot, "目标服务器地址:"+jobSnapshot.Url+"非法的响应"+time.Now().Local().Format("2006-01-02 15:04:05"), models.ERROR)
		log.Println("err = ", err)
		return err
	}

	addr := resp.Request.RemoteAddr

	log.Println("serveraddr=", addr)
	jobSnapshot.ServerAddress = addr
	if result.Success == true {
		nowTime := time.Now()
		d := nowTime.Sub(startTime)
		timeConsume := int64(d.Seconds())
		jobSnapshot.TimeConsume = timeConsume
		jobSnapshot.Result = result.Content
		jobSnapshot.Ip = common.GetIPFromUrl(jobSnapshot.Url)
		this.executeJobResult(jobSnapshot, "目标服务器地址:"+jobSnapshot.Url+"执行任务已经成功提交 "+time.Now().Local().Format("2006-01-02 15:04:05"), models.INVOKING)
		go this.processCheckJobResult(jobSnapshot)

	} else {
		this.executeJobResult(jobSnapshot, "目标服务器地址:"+jobSnapshot.Url+"执行失败:"+result.Message+time.Now().Local().Format("2006-01-02 15:04:05"), models.ERROR)

	}
	return nil

}

func (this *Invoker) executeJobResult(snapshot *models.JobSnapshot, detail, status string) {

	snapshot.Detail = snapshot.Detail + detail + "\n"
	snapshot.Status = status
	snapshot.ModifyTime = time.Now().Local()
	snapshot.UpdateSnapshot()

}
func (this *Invoker) Init(jobInfo *models.JobInfo, nextTime time.Time) (*models.JobSnapshot, error) {

	serverAddr := common.GetLocalAddr()
	snapshot := &models.JobSnapshot{
		JobInfoId:     jobInfo.Id,
		Name:          jobInfo.Name,
		Group:         jobInfo.Group,
		Status:        models.INIT,
		Url:           jobInfo.Urls,
		TimeConsume:   0,
		ServerAddress: serverAddr,
		NextTime:      nextTime,
		CreateTime:    time.Now().Local(),
		Detail:        "初始化" + time.Now().Local().Format("2006-01-02 15:04:05") + "\n",
		Params:        jobInfo.Param,
	}
	err := snapshot.InsertJobSnapshot()
	return snapshot, err
}

// 执行更新状态
func (this *Invoker) processCheckJobResult(jobSnapshot *models.JobSnapshot) {
	var quit bool = false
	var tick = time.Tick(time.Second * 5)
	for !quit {

		select {

		case <-tick:

			jobRequest := &common.JobRequest{JobSnapshot: jobSnapshot.Id, Params: jobSnapshot.Params, Status: jobSnapshot.Status}

			content, err := json.Marshal(jobRequest)

			if err != nil {
				this.executeJobResult(jobSnapshot, "解析job请求参数出错"+time.Now().Local().Format("2006-01-02 15:04:05"), models.ERROR)
				continue
			}

			resp, err := http.Post(jobSnapshot.Url, "application/json;charset=utf-8", bytes.NewBuffer(content))
			if err != nil {
				this.executeJobResult(jobSnapshot, "目标服务器地址:"+jobSnapshot.Url+"不可用"+time.Now().Local().Format("2006-01-02 15:04:05"), models.ERROR)
				continue
			}

			body, err := ioutil.ReadAll(resp.Body)

			if err != nil {
				resp.Body.Close()
				continue
			}
			result := &common.JobResponse{}
			log.Println("body = ", string(body))
			err = json.Unmarshal(body, result)
			if err != nil {
				resp.Body.Close()
				continue
			}
			log.Println("result= ", result)
			if result.Status == models.EXECUTING {
				this.executeJobResult(jobSnapshot, "目标服务器地址:"+jobSnapshot.Url+" 正在执行中..."+time.Now().Local().Format("2006-01-02 15:04:05"), models.EXECUTING)
			} else if result.Status == models.COMPLETED {
				this.executeJobResult(jobSnapshot, "目标服务器地址:"+jobSnapshot.Url+"任务执行完成..."+time.Now().Local().Format("2006-01-02 15:04:05"), models.COMPLETED)
				quit = true
			} else {
				this.executeJobResult(jobSnapshot, "目标服务器地址:"+jobSnapshot.Url+"任务执行失败..."+time.Now().Local().Format("2006-01-02 15:04:05"), models.ERROR)
				quit = true
			}

		}
	}
}
