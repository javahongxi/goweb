package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

var (
	ConfigFile = "conf/app.conf"
)

type DemoController struct {
	beego.Controller
}

func (c *DemoController) Hello() {
	loadConfig()
	c.Ctx.WriteString("Hello demo!")
}

func loadConfig() {
	err := beego.LoadAppConfig("ini", ConfigFile)
	if err != nil {
		logs.Critical("An error occurred:", err)
		panic(err)
	}
	logs.Info("load config name is", beego.AppConfig.String("name"))
}
