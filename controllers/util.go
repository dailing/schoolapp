package controllers

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/astaxie/beego"
)

func init() {
	rand.Seed(int64(time.Now().UnixNano()))
}

func Hash(str string) string {
	return str
}

func GenToken(usrname, psw string) string {
	return fmt.Sprintf("%s.%s.%016d", usrname, psw, GetTimeStamp())
}

func GenRandToken() string {
	return fmt.Sprintf("%016X", rand.Int63())
}

func GetTimeStamp() int {
	return int(time.Now().UnixNano())
}

func ErrReport(v interface{}) {
	if v != nil {
		beego.BeeLogger.Error("%v", v)
	}
}

func GenMataData() TypeMataData {
	return TypeMataData{
		TimeStamp: GetTimeStamp(),
		Device:    "server",
	}
}

func GenStatus(code int) TypeStatus {
	return TypeStatus{
		Code:        code,
		Description: ErrorDesp[code],
	}
}
