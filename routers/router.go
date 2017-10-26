package routers

import (
	"goweb/controllers"

	"github.com/astaxie/beego"
)

func init() {
	//beego.Router("/", &controllers.MainController{})
	//
	//beego.Router("/index", &controllers.UserController{}, "*:Index")
	////beego.Router("/express/search", &controllers.UserController{}, "POST:ExpressSearch")
	//
	//beego.Router("/express", &controllers.ExpressController{}, "GET:Express")
	//beego.Router("/express/search", &controllers.ExpressController{}, "POST:Search")

	// set home  path
	beego.Router("/",&controllers.IndexController{},"get:Index")

	// jobinfo
	beego.Router("/jobinfo/list",&controllers.JobInfoManagerController{},"*:List")
	beego.Router("/jobinfo/add",&controllers.JobInfoManagerController{},"get:ToAdd")
	beego.Router("/jobinfo/add",&controllers.JobInfoManagerController{},"post:Add")
	beego.Router("/jobinfo/edit",&controllers.JobInfoManagerController{},"get:ToEdit")
	beego.Router("/jobinfo/edit",&controllers.JobInfoManagerController{},"post:Edit")
	beego.Router("/jobinfo/delete",&controllers.JobInfoManagerController{},"post:Delete")
	beego.Router("/jobinfo/info",&controllers.JobInfoManagerController{},"get:Info")
	beego.Router("/jobinfo/active",&controllers.JobInfoManagerController{},"*:Active")
	// jobsnapshot
	beego.Router("/jobsnapshot/list",&controllers.JobSnapshotController{},"*:List")
	beego.Router("/jobsnapshot/info",&controllers.JobSnapshotController{},"get:Info")

	// jobinfohistory
	beego.Router("/jobinfohistory/list",&controllers.JobInfoHistoryController{},"*:List")

	//about
	beego.Router("/about",&controllers.AboutController{},"*:Index")

	//monitor
	beego.Router("/monitor/",&controllers.MonitorController{},"*:Index")
}
