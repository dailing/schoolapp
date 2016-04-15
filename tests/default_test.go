package test

import (
	_ "git.oschina.net/dddailing/schoolapp/routers"
	"path/filepath"
	"runtime"

	"github.com/astaxie/beego"
)

func init() {
	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
	beego.SetLogger("file", `{"filename":"testlog_test.log"}`)
	beego.SetLogFuncCall(true)
	beego.Info("appPath:", apppath)
}
