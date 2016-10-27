package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type UserGetController struct {
	beego.Controller
}

func (c *UserGetController) Post() {
	info := TypeRegularReq{}
	body := c.Ctx.Input.CopyBody(beego.AppConfig.DefaultInt64("bodybuffer", 1024*1024))
	beego.Info("Post Body is:", string(body))
	err := json.Unmarshal(body, &info)
	ErrReport(err)
	if err != nil {
		c.Abort("500")
		return
	}
	tInfo := ParseToken(info.Token)
	if tInfo.UserID <= 0 {
		c.Abort("400")
		return
	}
	userInfo, err := GetUserInfo(tInfo.UserName)
	ErrReport(err)
	retVal := TypeUserReq{
		MataData: GenMataData(),
		Status:   GenStatus(StatusCodeOK),
		UserInfo: userInfo,
	}
	// check username and psw
	retVal.Status = GenStatus(StatusCodeOK)
	c.Data["json"] = retVal
	c.ServeJSON()

}

type UserGetAddressController struct {
	beego.Controller
}

func (c *UserGetAddressController) Post() {
	//strId := c.Ctx.Input.Param(":uid")
	//id, err := strconv.ParseInt(strId, 10, 64)
	info := TypeRegularReq{}
	body := c.Ctx.Input.CopyBody(beego.AppConfig.DefaultInt64("bodybuffer", 1024*1024))
	beego.Info("Post Body is:", string(body))
	tokeninfo := ParseToken(info.Token)
	if tokeninfo.UserID <= 0 {
		c.Abort("400")
	}
	strId := fmt.Sprint(tokeninfo.UserID)
	o := orm.NewOrm()
	qs := o.QueryTable("lcn_customer_address")
	qs = qs.Filter("customer_id", strId).Filter("is_deleted", 0)
	retval := make([]TypeAixinwuAddress, 0)
	_, err := qs.All(&retval)
	ErrReport(err)
	if err != nil {
		c.Abort("500")
	}
	response := TypeAixinwuAddressResp{
		Status:  GenStatus(StatusCodeOK),
		Address: retval,
	}
	c.Data["json"] = response
	c.ServeJSON()
}

type UserGetRestfulController struct {
	beego.Controller
}

func (c *UserGetRestfulController) Get() {
	strid := c.Ctx.Input.Param(":uid")
	userinfo, err := GetUserInfoByID(strid)
	if err != nil {
		ErrReport(err)
		c.Abort("400")
	}
	c.Data["json"] = userinfo
	c.ServeJSON()
}
