package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type StaticsGetController struct {
	beego.Controller
}

func getStaticInfo() TypeStaticInfo {
	o := orm.NewOrm()
	// get price
	var totalPrice string
	err := o.Raw("select SUM(total_price) from `lcn_order` where status in(3,4,9)").QueryRow(&totalPrice)
	ErrReport(err)
	if err != nil {
		totalPrice = ""
	}

	var itemCount string
	err = o.Raw("SELECT sum(quantity) as total FROM `lcn_order_item` WHERE order_id in (select id from lcn_order where status in (3,4,9) and is_delete=0)").QueryRow(&itemCount)
	ErrReport(err)
	if err != nil {
		itemCount = ""
	}

	var UserCout string
	err = o.Raw("select count(*) as total from lcn_customer").QueryRow(&UserCout)
	ErrReport(err)
	if err != nil {
		UserCout = ""
	}

	var visit string
	err = o.Raw("SELECT SUM(`value`) FROM lcn_site_config WHERE `key`='visit'").QueryRow(&visit)
	ErrReport(err)
	if err != nil {
		fmt.Println(err)
		visit = ""
	}

	return TypeStaticInfo{
		Money: totalPrice,
		Item:  itemCount,
		User:  UserCout,
		Visit: visit,
	}
}

func (c *StaticsGetController) Post() {
	beego.Debug("static get controller")
	response := TypeStaticInfoResp{
		MataData:   GenMataData(),
		StaticInfo: getStaticInfo(),
	}
	// return
	response.Status = GenStatus(StatusCodeOK)
	c.Data["json"] = response
	c.ServeJSON()
}
