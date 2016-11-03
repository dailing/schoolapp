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
	userInfo.VerificationCode = ""
	userInfo.Password = ""
	ErrReport(err)
	retVal := TypeUserReq{
		MataData: GenMataData(),
		Status:   GenStatus(StatusCodeOK),
		UserInfo: userInfo,
	}
	// check username and psw
	retVal.Status = GenStatus(StatusCodeOK)
	retVal.UserInfo.NickName = BaseDecode(retVal.UserInfo.NickName)
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
	err := json.Unmarshal(body, &info)
	ErrReport(err)
	tokeninfo := ParseToken(info.Token)
	if tokeninfo.UserID <= 0 {
		c.Abort("400")
	}
	o := orm.NewOrm()
	localuserino, err := GetUserInfoByID(fmt.Sprint(tokeninfo.UserID))
	ErrReport(err)
	aixinwuUserinfo := TypeAixinwuJaccountInfo{
		Jaccount_id: localuserino.JAccount,
	}
	ErrReport(o.Read(&aixinwuUserinfo, "jaccount_id"))

	strId := fmt.Sprint(aixinwuUserinfo.Customer_id)
	qs := o.QueryTable(&TypeAixinwuAddress{})
	qs = qs.Filter("customer_id", strId).Filter("is_deleted", 0)
	retval := make([]TypeAixinwuAddress, 0)
	_, err = qs.All(&retval)
	beego.Trace("UserID :", strId, " ", retval)
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
	c.Data["json"] = struct {
		Username string `json:"username" orm:"type(text);unique;column(username)"`
		NickName string `json:"nickname" orm:"type(text);column(nickname)"`
	}{
		Username: userinfo.Username,
		NickName: userinfo.NickName,
	}
	c.ServeJSON()
}
