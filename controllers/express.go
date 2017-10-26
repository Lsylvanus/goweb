package controllers

import (
	"goweb/models"
	"goweb/common"

	"github.com/astaxie/beego/logs"

	"strings"
	"encoding/json"
)

type (
	ExpressController struct {
		common.BeeControllers
	}
)

func (e *ExpressController) Express() {
	e.Data["Title"] = "Express Search"
	e.Data["Info"] = "Personal Authorization Information"
	e.TplName = "exp/express.tpl"
}

func (e *ExpressController) Search() {
	comType := e.GetString("com_type")
	postId := e.GetString("post_id")
	if comType == "" && postId == "" {
		e.Ctx.WriteString("快递公司与单号不能为空")
		return
	}
	urlStr := "http://www.kuaidi100.com/query?type=" + strings.TrimSpace(comType) + "&postid=" + strings.TrimSpace(postId)
	data, err := common.DoReq("GET", urlStr, nil)
	if err != nil {
		e.Ctx.WriteString(err.Error())
		return
	}
	delivery := new(models.Delivery)
	err = json.Unmarshal(data, delivery)
	if err != nil {
		e.Ctx.WriteString("json unmarshal fail : " + err.Error())
		return
	}
	results := delivery.Data
	exps := make([]models.Express, 0)
	for _, result := range results {
		logs.Info("time :", result.Time)
		logs.Info("context :", result.Context)
		exps = append(exps, result)
	}
	e.Data["Exps"] = exps
	e.TplName = "exp/success.tpl"
}