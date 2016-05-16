package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
)

/*
 *  Notice that this controller only updates the status codes.
 */

type ItemSetController struct {
	beego.Controller
}

func (c *ItemSetController) Post() {
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
	tInfo := ParseToken(request.Token)
	if tInfo.UserID <= 0 {
		c.Abort("401")
		return
	}
	item, err := GetItemByID(request.ItemInfo.ID)
	ErrReport(err)
	if err != nil {
		c.Abort("500")
		beego.Trace()
		return
	}
	originalItem, err := GetItemByID(request.ItemInfo.ID)
	if originalItem.OwnerID != tInfo.UserID {
		c.Abort("403")
		return
	}
	err = SetItem(item)
	// ser parameters
	ErrReport(err)
	if err != nil {
		response.Status.Code = 200
		response.Status.Description = err.Error()
	}
	if response.Status.Code == 0 {
		response.Status = GenStatus(StatusCodeOK)
	}
	//response.ItemInfo = itemInfo
	c.Data["json"] = response
	c.ServeJSON()
}
