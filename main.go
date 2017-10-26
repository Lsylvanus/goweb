package main

import (
	log "goweb/logs"
	"goweb/models"
	_ "goweb/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/redis.v5"

	"flag"
	"runtime"
	"goweb/job"
)

func init() {
	dbCon := flag.String("db", "root:chenghai3c@/scheduler?charset=utf8&loc=Local", "Mysql db paht")
	flag.Parse()

	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", *dbCon)
	orm.RegisterModel(&models.JobInfo{}, &models.JobInfoHistory{}, &models.JobSnapshot{})
}

func Cache() *redis.Client {
	cache := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	if pong, err := cache.Ping().Result(); err != nil {
		logs.Error(pong, err)
	}
	return cache
}

func main() {
	log.InitLogs()

	// set CPU
	runtime.GOMAXPROCS(runtime.NumCPU())
	orm.Debug = true
	jobManager := job.NewJobMnager()
	jobManager.PushAllJob()
	// TODO Init jobList

	// set static resource
	beego.SetStaticPath("static","static")
	beego.SetStaticPath("public","static")

	// start web app
	beego.Run()

}
