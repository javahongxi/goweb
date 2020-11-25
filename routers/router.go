package routers

import (
	"github.com/astaxie/beego"
	"goweb/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/hello", &controllers.DemoController{}, "get:Hello")
}
