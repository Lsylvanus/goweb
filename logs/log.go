package logs

import "github.com/astaxie/beego/logs"

func InitLogs() {
	logs.SetLogger("console")
	logs.EnableFuncCallDepth(true)
	logs.Async(1e3)
	logs.SetLogger(logs.AdapterFile,`{"filename":"debug.log", "level":7, "maxlines":0, "maxsize":0, "daily":true, "maxdays":10}`)
}