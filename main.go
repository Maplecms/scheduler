package main

import (
	"github.com/astaxie/beego"
	"scheduler/controller"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"scheduler/entity"
	"scheduler/job"
	"runtime"
)


func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:123456@/scheduler?charset=utf8&loc=Local")
	orm.RegisterModel(&entity.JobInfo{},&entity.JobInfoHistory{},&entity.JobSnapshot{})

}
func main()  {

	// set CPU
	runtime.GOMAXPROCS(runtime.NumCPU())
	orm.Debug = true
	jobManager := job.NewJobMnager()
	jobManager.PushAllJob()
	// TODO Init jobList

	// set home  path
	beego.Router("/",&controller.IndexController{},"get:Index")

	// jobinfo
	beego.Router("/jobinfo/list",&controller.JobInfoManagerController{},"*:List")
	beego.Router("/jobinfo/add",&controller.JobInfoManagerController{},"get:ToAdd")
	beego.Router("/jobinfo/add",&controller.JobInfoManagerController{},"post:Add")
	beego.Router("/jobinfo/edit",&controller.JobInfoManagerController{},"get:ToEdit")
	beego.Router("/jobinfo/edit",&controller.JobInfoManagerController{},"post:Edit")
	beego.Router("/jobinfo/delete",&controller.JobInfoManagerController{},"post:Delete")
	beego.Router("/jobinfo/info",&controller.JobInfoManagerController{},"get:Info")
	beego.Router("/jobinfo/active",&controller.JobInfoManagerController{},"*:Active")
	// jobsnapshot
	beego.Router("/jobsnapshot/list",&controller.JobSnapshotController{},"*:List")
	beego.Router("/jobsnapshot/info",&controller.JobSnapshotController{},"get:Info")

	// jobinfohistory
	beego.Router("/jobinfohistory/list",&controller.JobInfoHistoryController{},"*:List")

	//about
	beego.Router("/about",&controller.AboutController{},"*:Index")

	//monitor
	beego.Router("/monitor/",&controller.MonitorController{},"*:Index")

	// set static resource
	beego.SetStaticPath("static","static")
	beego.SetStaticPath("public","static")


	// start web app
	beego.Run()

}

