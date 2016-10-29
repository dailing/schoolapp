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

type AinxinwuAddressSetController struct {
	beego.Controller
}

func (c *AinxinwuAddressSetController) Post() {
	response := TypeRegularResp{
		Status: GenStatus(StatusCodeOK),
	}
	beego.Debug("get address")
	request := TypeSetAddressReq{}
	body := c.Ctx.Input.CopyBody(beego.AppConfig.DefaultInt64("bodybuffer", 1024*1024))
	beego.Info("Post Body is:", string(body))
	err := json.Unmarshal(body, &request)
	ErrReport(err)
	if err != nil {
		response.Status = GenStatus(StatusCodeErrorLoginInfo)
	}

	tokeninfo := ParseToken(request.Token)
	uid := TransferLocalIDtoAixinwuID(tokeninfo.UserID)
	address := TypeAixinwuAddress{
		Customer_id: fmt.Sprint(uid),
		Is_default:  1,
	}
	o := orm.NewOrm()
	err = o.Read(&address, "customer_id", "is_default")
	ErrReport(err)
	if err != nil {
		response.Status = GenStatus(StatusCodeDatabaseErr)
	}
	address.Is_default = 0
	o.Update(&address, "is_default")
	address.Id = 0
	address.Is_default = 1
	address.Consignee = request.Consignee
	address.Mobile = request.Mobile
	address.Snum = request.Snum
	_, err = o.Insert(&address)
	ErrReport(err)
	if err != nil {
		response.Status = GenStatus(StatusCodeDatabaseErr)
	}
	c.Data["json"] = response
	c.ServeJSON()
}
