package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
)

type ItemAddController struct {
	beego.Controller
}

func (c *ItemAddController) Post() {
	beego.Debug("add user")
	request := TypeItemReqResp{}
	body := c.Ctx.Input.CopyBody(beego.AppConfig.DefaultInt64("bodybuffer", 1024*1024))
	beego.Info("Post Body is:", string(body))
	err := json.Unmarshal(body, &request)
	ErrReport(err)
	if err != nil {
		c.Abort("500")
		return
	}
	response := TypeRegularResp{
		MataData: GenMataData(),
	}
	// check token
	tInfo := ParseToken(request.Token)
	if tInfo.UserID <= 0 {
		c.Abort("400")
		return
	}
	// set parameters
	request.ItemInfo.OwnerID = tInfo.UserID
	_, err = AddItem(request.ItemInfo)
	response.Status = GenStatus(StatusCodeOK)
	c.Data["json"] = response
	c.ServeJSON()
}
