package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
)

type ItemGetController struct {
	beego.Controller
}

func (c *ItemGetController) Post() {
	beego.Debug("add user")
	request := TypeItemReqResp{}
	body := c.Ctx.Input.CopyBody(beego.AppConfig.DefaultInt64("bodybuffer", 1024*1024))
	beego.Info("Post Body is:", string(body))
	err := json.Unmarshal(body, &request)
	ErrReport(err)
	if err != nil {
		c.Abort("400")
		return
	}
	response := TypeItemReqResp{
		MataData: GenMataData(),
	}
	// check token
	//tInfo := ParseToken(request.Token)
	//if tInfo.UserID <= 0 {
	//	c.Abort("401")
	//	return
	//}
	// ser parameters
	itemInfo, err := GetItemByID(request.ItemInfo.ID)
	ErrReport(err)
	response.Status = GenStatus(StatusCodeOK)
	response.ItemInfo = itemInfo
	c.Data["json"] = response
	c.ServeJSON()
}
