package controllers

import (
	"encoding/json"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
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

func (c *LoginController) _Get() {
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
	succ, err := checkLogIn(info.UserInfo.Username, info.UserInfo.Password)
	if !succ {
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

func checkLogIn(username, psw string) (bool, error) {
	o := orm.NewOrm()
	user := SQLuserinfo{
		Username: username,
	}
	err := o.Read(&user, "username")
	ErrReport(err)
	if err != nil {
		return false, err
	}
	beego.Trace("Read User Psw:", user.Password)
	if psw == user.Password {
		return true, nil
	}
	return false, nil
}
