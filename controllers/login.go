package controllers

import (
	"encoding/json"
	"time"

	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

type TypeLoginInfo struct {
	MataData TypeMataData `json:"mataData"`
	UserInfo TypeUserInfo `json:"userinfo"`
}

type TypeLoginResp struct {
	MataData TypeMataData `json:"mataData"`
	Token    string       `json:"token"`
	Status   TypeStatus   `json:"status"`
}

func (c *LoginController) Get() {
	c.Data["json"] = TypeLoginInfo{
		MataData: TypeMataData{
			TimeStamp: int(time.Now().UnixNano()),
			Device:    "test",
		},
		UserInfo: TypeUserInfo{
			Username: "test",
			Password: "psw",
		},
	}
	c.ServeJSON()
}

func (c *LoginController) Post() {
	info := TypeLoginInfo{}
	beego.Trace(string(c.Ctx.Input.RequestBody))
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &info)
	ErrReport(err)
	if err != nil {
		c.Abort("500")
		return
	}
	retval := TypeLoginResp{
		MataData: TypeMataData{
			TimeStamp: GetTimeStamp(),
			Device:    "Server",
		},
		Token: "",
	}
	// check username and psw
	if !checkLogIn(info.UserInfo.Username, info.UserInfo.Password) {
		retval.Status.Code = StatusCodeErrorLoginInfo
		retval.Status.Description = ErrorDesp[StatusCodeErrorLoginInfo]
		c.Data["json"] = retval
		c.ServeJSON()
		return
	}
	retval.Status.Code = StatusCodeOK
	retval.Status.Description = ErrorDesp[StatusCodeOK]
	retval.Token = GenToken(info.UserInfo.Username, info.UserInfo.Password)
	c.Data["json"] = retval
	c.ServeJSON()
}

func checkLogIn(username, psw string) bool {
	return true
}
