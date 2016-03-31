package controllers

import (
	"fmt"
	"time"

	"github.com/astaxie/beego"
)

func Hash(str string) string {
	return str
}

func GenToken(usrname, psw string) string {
	return fmt.Sprintf("%s.%s.%016d", usrname, psw, GetTimeStamp())
}

func GetTimeStamp() int {
	return int(time.Now().UnixNano())
}

func ErrReport(v interface{}) {
	if v != nil {
		beego.BeeLogger.Error("%v", v)
	}
}
