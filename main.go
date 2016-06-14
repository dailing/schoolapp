package main

import (
	"git.oschina.net/dddailing/schoolapp/controllers"
	"git.oschina.net/dddailing/schoolapp/natsChatServer"
	_ "git.oschina.net/dddailing/schoolapp/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.SetLogger("file", `{"filename":"test.log"}`)
	beego.Info("This is an info log")
	controllers.SysInit()
	// start chat server
	server := natsChatServer.NewServer()
	server.Start()

	beego.Run()
}
