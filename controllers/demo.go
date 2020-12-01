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
	c.Ctx.WriteString("Hello, " + c.GetString("name", "Kitty"))
}

func (c *DemoController) Hi() {
	c.Ctx.WriteString("Hi, " + c.GetString("name", "Kitty"))
}

func loadConfig() {
	err := beego.LoadAppConfig("ini", ConfigFile)
	if err != nil {
		logs.Critical("An error occurred:", err)
		panic(err)
	}
	logs.Info("load config name is", beego.AppConfig.String("name"))
}
