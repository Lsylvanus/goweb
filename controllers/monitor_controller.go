package controllers

import (
	"goweb/common"
	"goweb/job"
)

type MonitorController struct {
	BaseController
}

func (this *MonitorController) Index() {

	jobManger := job.NewJobMnager()
	jobList, err := jobManger.GetJobSnapshotList()
	if err != nil {
		common.PanicIf(err)

	}

	this.TplName = "monitor/index.html"
	this.Data["jobList"] = jobList
	this.Render()
}
