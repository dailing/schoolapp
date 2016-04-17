package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
)

type UserUpdateController struct {
	beego.Controller
}

func (c *UserUpdateController) Post() {
	info := TypeUserReq{}
	body := c.Ctx.Input.CopyBody(beego.AppConfig.DefaultInt64("bodybuffer", 1024*1024))
	beego.Info("Post Body is:", string(body))
	err := json.Unmarshal(body, &info)
	ErrReport(err)
	if err != nil {
		c.Abort("500")
		return
	}
	Token := ParseToken(info.Token)
	if Token.UserName == "" {
		c.Abort("400")
		return
	}
	info.UserInfo.Username = Token.UserName
	retVal := TypeUserReq{
		MataData: GenMataData(),
		Status:   GenStatus(StatusCodeOK),
	}
	// perform update
	err = UpdateUserInfo(info.UserInfo)
	ErrReport(err)
	if err != nil {
		c.Abort("500")
		return
	}
	retVal.Status = GenStatus(StatusCodeOK)
	c.Data["json"] = retVal
	c.ServeJSON()

}
