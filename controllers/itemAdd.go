package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
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
		request.ItemInfo.OwnerID = tInfo.UserID
	}
	beego.Trace("description", request.ItemInfo.Description)
	request.ItemInfo.OwnerID = tInfo.UserID
	// add coins to user
	o := orm.NewOrm()
	info := TypeItemInfo{
		OwnerID: tInfo.UserID,
	}
	err = o.Read(&info, "ownerID")
	beego.Trace(info)
	ErrReport(err)
	userinfo, err2 := GetUserInfoByID(fmt.Sprint(tInfo.UserID))
	ErrReport(err2)
	if userinfo.JAccount != "" && err == orm.ErrNoRows {
		beego.Trace("Adding bonus")
		aixinwuid := TransferLocalIDtoAixinwuID(tInfo.UserID)
		aiUser := TypeAixinwuCustomCash{
			User_id: aixinwuid,
		}
		err = o.Read(&aiUser, "user_id")
		ErrReport(err)
		aiUser.Total += 25
		_, err = o.Update(&aiUser, "total")
		ErrReport(err)
	}

	itemID, err := AddItem(request.ItemInfo)
	response.Status = GenStatus(StatusCodeOK)
	response.ItemInfo = request.ItemInfo
	response.ItemInfo.OwnerID = tInfo.UserID
	response.ItemInfo.ID = itemID

	c.Data["json"] = response
	c.ServeJSON()
}
