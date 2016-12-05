package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type UserUpdateController struct {
	beego.Controller
}

func (c *UserUpdateController) Post() {
	info := TypeUserReq{}
	body := c.Ctx.Input.CopyBody(beego.AppConfig.DefaultInt64("bodybuffer", 1024*1024))
	beego.Info("Post Body is:", string(body))
	err := json.Unmarshal(body, &info)
	beego.Trace(string(body))
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

type UserChangePswController struct {
	beego.Controller
}

func (c *UserChangePswController) Post() {
	info := TypeUserInfo{}
	body := c.Ctx.Input.CopyBody(beego.AppConfig.DefaultInt64("bodybuffer", 1024*1024))
	beego.Info("Post Body is:", string(body))
	err := json.Unmarshal(body, &info)
	beego.Trace(string(body))
	ErrReport(err)
	if err != nil {
		c.Abort("400")
		return
	}
	if CheckVerificationCode(info.Username, info.VerificationCode) {
		i, err := GetUserInfo(info.Username)
		ErrReport(err)
		i.Password = info.Password
		o := orm.NewOrm()
		o.Update(&i)
	}
	response := TypeRegularResp{}
	response.Status = GenStatus(StatusCodeOK)
	if err != nil {
		response.Status = GenStatus(StatusCodeDatabaseErr)
	}
	c.Data["json"] = response
	c.ServeJSON()
}
