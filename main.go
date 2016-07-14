package main

import (
	"io/ioutil"
	"os"
	"os/exec"
	"time"

	"git.oschina.net/dddailing/schoolapp/controllers"
	"git.oschina.net/dddailing/schoolapp/natsChatServer"
	_ "git.oschina.net/dddailing/schoolapp/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.SetLogger("file", `{"filename":"test.log"}`)
	beego.Info("This is an info log")
	controllers.SysInit()
	// start gnatsd
	if _, err := os.Stat("gnatsd"); os.IsNotExist(err) {
		beego.Warning("No gnatsd found")
	} else {
		cmd := exec.Command("./gnatsd", "-l", "gnatsd.log", "-V")
		if cmd == nil {
			beego.Error()
		}
		cmd.Stderr = ioutil.Discard
		cmd.Stdout = ioutil.Discard
		controllers.ErrReport(cmd.Start())
		defer cmd.Process.Kill()
	}
	// start chat server
	time.Sleep(time.Second)
	server := natsChatServer.NewServer()
	server.Start()
	beego.Run()
}
