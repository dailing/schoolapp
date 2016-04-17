package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

type TypeLoginResp struct {
	MataData TypeMataData `json:"mataData"`
	Token    string       `json:"token"`
	Status   TypeStatus   `json:"status"`
}

func (c *LoginController) Post() {
	info := TypeUserReq{}
	body := c.Ctx.Input.CopyBody(1024 * 1024)
	beego.Trace(string(body))
	err := json.Unmarshal(body, &info)
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
	id, err := checkLogIn(info.UserInfo.Username, info.UserInfo.Password)
	if id <= 0 {
		retval.Status.Code = StatusCodeErrorLoginInfo
		retval.Status.Description = ErrorDesp[StatusCodeErrorLoginInfo]
		c.Data["json"] = retval
		c.ServeJSON()
		return
	}
	// gentoken and setup redisDB
	retval.Status.Code = StatusCodeOK
	retval.Status.Description = ErrorDesp[StatusCodeOK]
	tInfo := TypeTokenInfo{
		UserName: info.UserInfo.Username,
		UserID:   id,
	}
	retval.Token = GenToken(tInfo)
	c.Data["json"] = retval
	c.ServeJSON()
}

func checkLogIn(username, psw string) (int, error) {
	user, err := GetUserInfo(username)
	ErrReport(err)
	if err != nil {
		return -1, err
	}
	beego.Trace("Read User Psw:", user.Password)
	if psw == user.Password {
		return user.ID, nil
	}
	return -1, nil
}
