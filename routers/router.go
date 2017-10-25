package routers

import (
	"goweb/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})

    beego.Router("/index", &controllers.UserController{}, "*:Index")
	//beego.Router("/express/search", &controllers.UserController{}, "POST:ExpressSearch")

	beego.Router("/express", &controllers.ExpressController{}, "GET:Express")
	beego.Router("/express/search", &controllers.ExpressController{}, "POST:Search")
}
