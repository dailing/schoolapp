package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
)

type ChatGetController struct {
	beego.Controller
}

func (c *ChatGetController) Post() {
	beego.Debug("chart get")
	request := TypeChatReq{}
	body := c.Ctx.Input.CopyBody(beego.AppConfig.DefaultInt64("bodybuffer", 1024*1024))
	beego.Info("Post Body is:", string(body))
	err := json.Unmarshal(body, &request)
	ErrReport(err)
	if err != nil {
		c.Abort("400")
		return
	}
	response := TypeChatResp{
		MataData: GenMataData(),
	}
	// check token
	tInfo := ParseToken(request.Token)
	if tInfo.UserID <= 0 {
		c.Abort("401")
		return
	}
	// ser parameters
	response.Chat = GetChat(request.Chat.ItemID, tInfo.UserID)
	ErrReport(err)
	response.Status = GenStatus(StatusCodeOK)
	//response.ItemInfo = itemInfo
	c.Data["json"] = response
	c.ServeJSON()
}
