package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
)

type AixinwuAddressGetController struct {
	beego.Controller
}

func (c *AixinwuAddressGetController) Get() {
	beego.Debug("get address")
	request := TypeAixinwuItemReqResp{
		Category: -1,
	}
	body := c.Ctx.Input.CopyBody(beego.AppConfig.DefaultInt64("bodybuffer", 1024*1024))
	beego.Info("Post Body is:", string(body))
	err := json.Unmarshal(body, &request)
	beego.Trace("body is :", string(body))
	ErrReport(err)
	if err != nil {
		c.Abort("400")
		return
	}
	response := TypeAixinwuItemReqResp{
		MataData: GenMataData(),
	}
	response.AixinwuItems = GetAixintuItems(request.StartAt, request.Length, request.Category, request.Type)
	c.Data["json"] = response
	c.ServeJSON()
}
