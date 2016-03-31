package main

import (
	_ "git.oschina.net/dddailing/schoolapp/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.SetLogger("file", `{"filename":"test.log"}`)
	beego.Run()
}
