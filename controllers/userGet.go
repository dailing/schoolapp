package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
)

type UserGetController struct {
	beego.Controller
}

func (c *UserGetController) Post() {
	info := TypeRegularReq{}
	body := c.Ctx.Input.CopyBody(beego.AppConfig.DefaultInt64("bodybuffer", 1024*1024))
	beego.Info("Post Body is:", string(body))
	err := json.Unmarshal(body, &info)
	ErrReport(err)
	if err != nil {
		c.Abort("500")
		return
	}
	tInfo := ParseToken(info.Token)
	if tInfo.UserID <= 0 {
		c.Abort("400")
		return
	}
	userInfo, err := GetUserInfo(tInfo.UserName)
	ErrReport(err)
	retVal := TypeUserReq{
		MataData: GenMataData(),
		Status:   GenStatus(StatusCodeOK),
		UserInfo: userInfo,
	}
	// check username and psw
	retVal.Status = GenStatus(StatusCodeOK)
	c.Data["json"] = retVal
	c.ServeJSON()

}
