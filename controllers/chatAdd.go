package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
)

type ChatAddController struct {
	beego.Controller
}

func (c *ChatAddController) Post() {
	beego.Debug("add user")
	request := TypeChatReq{}
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
	request.Chat.PublisherID = tInfo.UserID
	// do insert
	_, err = AddChat(request.Chat)
	ErrReport(err)
	//request.ItemInfo.OwnerID = tInfo.UserID
	//itemID, err := AddItem(request.ItemInfo)
	response.Status = GenStatus(StatusCodeOK)
	//response.ItemInfo = request.ItemInfo
	//response.ItemInfo.OwnerID = tInfo.UserID
	//response.ItemInfo.ID = itemID
	c.Data["json"] = response
	c.ServeJSON()
}
