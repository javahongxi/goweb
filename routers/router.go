package routers

import (
	"github.com/astaxie/beego"
	"goweb/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.Router("/hello", &controllers.DemoController{}, "get:Hello")

	beego.Router("/user/add", &controllers.UserController{}, "get:Add")
	beego.Router("/user/delete", &controllers.UserController{}, "get:Delete")
	beego.Router("/user/update", &controllers.UserController{}, "get:Update")
	beego.Router("/user/find", &controllers.UserController{}, "get:Find")
}
