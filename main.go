package main

import (
	_ "goweb/routers"
	"goweb/logs"

	"github.com/astaxie/beego"
)

func main() {
	logs.InitLogs()

	beego.Run()
}