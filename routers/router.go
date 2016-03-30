package routers

import (
	"git.oschina.net/dddailing/schoolapp/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
