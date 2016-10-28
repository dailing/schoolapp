package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type AixinwuAddressGetController struct {
	beego.Controller
}

func (c *AixinwuAddressGetController) getAdddress(userid string) []TypeAixinwuAddress {
	o := orm.NewOrm()
	q := o.QueryTable(&TypeAixinwuAddress{})
	q = q.Filter("Customer_id", userid)
	retval := make([]TypeAixinwuAddress, 0)
	_, err := q.All(&retval)
	ErrReport(err)
	if err != nil {
		return nil
	}
	return retval
}

func (c *AixinwuAddressGetController) Get() {
	beego.Debug("get address")
	request := TypeRegularReq{}
	body := c.Ctx.Input.CopyBody(beego.AppConfig.DefaultInt64("bodybuffer", 1024*1024))
	beego.Info("Post Body is:", string(body))
	err := json.Unmarshal(body, &request)
	ErrReport(err)
	if err != nil {
		c.Abort("400")
		return
	}
	tokenInfo := ParseToken(request.Token)
	if tokenInfo.UserID <= 0 {
		beego.Warn("Error token:", request.Token)
		c.Abort("400")
	}
	response := TypeAixinwuAddressResp{
		MataData: GenMataData(),
		Status:   GenStatus(StatusCodeOK),
		Address:  c.getAdddress(fmt.Sprint(tokenInfo.UserID)),
	}
	c.Data["json"] = response
	c.ServeJSON()
}
