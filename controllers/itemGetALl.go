package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
)

type ItemGetAllController struct {
	beego.Controller
}

func (c *ItemGetAllController) Post() {
	beego.Debug("add user")
	request := TypeRegularReq{}
	body := c.Ctx.Input.CopyBody(beego.AppConfig.DefaultInt64("bodybuffer", 1024*1024))
	beego.Info("Post Body is:", string(body))
	err := json.Unmarshal(body, &request)
	ErrReport(err)
	if err != nil {
		c.Abort("400")
		return
	}
	response := TypeGetItemsResp{
		MataData: GenMataData(),
	}
	// check token
	tInfo := ParseToken(request.Token)
	if tInfo.UserID <= 0 {
		c.Abort("401")
		return
	}
	// ser parameters
	itemInfo := GetAllItem()
	response.Status = GenStatus(StatusCodeOK)
	response.Items = itemInfo
	c.Data["json"] = response
	c.ServeJSON()
}
