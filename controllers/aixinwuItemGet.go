package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
)

type AixintuItemGetController struct {
	beego.Controller
}

func (c *AixintuItemGetController) Post() {
	beego.Debug("add user")
	request := TypeAixinwuItemReqResp{
		Category: -1,
	}
	body := c.Ctx.Input.CopyBody(beego.AppConfig.DefaultInt64("bodybuffer", 1024*1024))
	beego.Info("Post Body is:", string(body))
	err := json.Unmarshal(body, &request)
	ErrReport(err)
	if err != nil {
		c.Abort("400")
		return
	}
	response := TypeAixinwuItemReqResp{
		MataData: GenMataData(),
	}
	//	// check token
	//	tInfo := ParseToken(request.Token)
	//	if tInfo.UserID <= 0 {
	//		c.Abort("401")
	//		return
	//	}
	//	//
	//		ErrReport(err)
	response.AixinwuItems = GetAixintuItems(request.StartAt, request.Length, request.Category)
	c.Data["json"] = response
	c.ServeJSON()
}
