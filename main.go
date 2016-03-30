package main

import (
	"git.oschina.net/dddailing/schoolapp/controllers"
	_ "git.oschina.net/dddailing/schoolapp/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.RESTRouter("/login", &controllers.LoginController{})
	beego.Run()
}
