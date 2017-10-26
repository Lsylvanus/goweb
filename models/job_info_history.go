package models

import (
	"github.com/astaxie/beego/orm"

	"time"
)

type JobInfoHistory struct {
	Id                  int    `orm:"pk;auto"`
	Name                string `orm:"notnull"`
	Group               string `orm:"notnull"`
	Type                string `orm:"notnull"`
	Time                time.Time
	Cron                string
	Urls                string `orm:"notnull"`
	ClassPath           string
	InvokePolicy        string
	IsActivity          int `orm:"notnull"`
	Desc                string
	CreateTime          time.Time
	ModifyTime          time.Time
	Param               string
	LatestTriggerTime   time.Time
	LatestServerAddress string
	OwnerPhone          string
}

// save
func (this *JobInfoHistory) SaveJobInfoHistory() error {
	_, err := orm.NewOrm().Insert(this)
	return err
}

// find list
func (this *JobInfoHistory) FindAllJobInfoList() ([]*JobInfoHistory, error) {

	var historys []*JobInfoHistory
	o := orm.NewOrm()
	qs := o.QueryTable("job_info_history")
	if this.Name != "" {
		qs = qs.Filter("name", this.Name)
	}

	if this.Group != "" {

		qs = qs.Filter("group", this.Group)
	}

	if this.Group == "" && this.Name == "" {

		qs = qs.Limit(100)
	}

	_, err := qs.All(&historys)
	return historys, err

}
