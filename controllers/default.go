package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"

	"strings"
	"net/http"
	"time"
	"io/ioutil"
	"encoding/json"
	"goweb/models"
)

type MainController struct {
	beego.Controller
}

type UserController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (u *UserController) Index()  {
	u.Data["Title"] = "Express Search"
	u.Data["Info"] = "Personal Authorization Information"
	u.TplName = "exp/express.tpl"
	logs.Info("--------------- express info ---------------")
}

func (u *UserController) ExpressSearch()  {
	comType := u.GetString("com_type")
	postId := u.GetString("post_id")
	if comType == "" && postId == "" {
		u.Ctx.WriteString("快递公司与单号不能为空")
		return
	}
	urlStr := "http://www.kuaidi100.com/query?type=" + strings.TrimSpace(comType) + "&postid=" + strings.TrimSpace(postId)
	req, err := http.NewRequest("GET", urlStr, nil)
	if err != nil {
		u.Ctx.WriteString("创建请求失败")
		return
	}
	client := http.Client{Timeout:time.Second * 30}
	resp, err := client.Do(req)
	if err != nil {
		u.Ctx.WriteString("请求失败")
		return
	}
	defer resp.Body.Close()
	if resp != nil {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			u.Ctx.WriteString("读取页面内容失败")
			return
		}
		delivery := new(models.Delivery)
		err = json.Unmarshal(body, delivery)
		if err != nil {
			u.Ctx.WriteString("JSON解析失败")
			return
		}
		results := delivery.Data
		exps := make([]models.Express, 0)
		for _, result := range results {
			logs.Info("time :", result.Time)
			logs.Info("context :", result.Context)
			exps = append(exps, result)
		}
		u.Data["Exps"] = exps
	}
	u.TplName = "exp/success.tpl"
}