package main

import (
	_ "goweb/routers"
	"github.com/astaxie/beego"
	"goweb/logs"
)

func main() {
	logs.InitLogs()

	beego.Run()
}