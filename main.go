package main

import (
	"git.oschina.net/dddailing/schoolapp/controllers"
	_ "git.oschina.net/dddailing/schoolapp/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.SetLogger("file", `{"filename":"test.log"}`)
	beego.Info("This is an info log")
	controllers.SysInit()
	beego.Run()
}
