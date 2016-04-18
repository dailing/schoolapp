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
	tbody, err := json.Marshal(request)
	beego.Trace(string(tbody))
	ErrReport(err)
	if err != nil {
		c.Abort("500")
		return
	}
	response := TypeItemReqResp{
		MataData: GenMataData(),
	}
	// check token
	tInfo := ParseToken(request.Token)
	if tInfo.UserID <= 0 {
		c.Abort("400")
		return
	}
	// set parameters
	if request.ItemInfo.Description == "" {
		request.ItemInfo.Description = "No Request"
	}
	beego.Trace("description", request.ItemInfo.Description)
	request.ItemInfo.OwnerID = tInfo.UserID
	itemID, err := AddItem(request.ItemInfo)
	response.Status = GenStatus(StatusCodeOK)
	response.ItemInfo = request.ItemInfo
	response.ItemInfo.OwnerID = tInfo.UserID
	response.ItemInfo.ID = itemID
	c.Data["json"] = response
	c.ServeJSON()
}
