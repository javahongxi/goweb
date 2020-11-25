package routers

import (
	"github.com/astaxie/beego"
	"goweb/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
