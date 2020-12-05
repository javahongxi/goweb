package routers

import (
	"github.com/astaxie/beego"
	"github.com/javahongxi/goweb/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.Router("/hello", &controllers.DemoController{}, "get:Hello;post:Hi")

	beego.Router("/user/hello", &controllers.UserController{}, "post:Hello")
	beego.Router("/user/add", &controllers.UserController{}, "get,post:Add")
	beego.Router("/user/delete", &controllers.UserController{}, "post:Delete")
	beego.Router("/user/update", &controllers.UserController{}, "get,post:Update")
	beego.Router("/user/find", &controllers.UserController{}, "*:Find")
}
