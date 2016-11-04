package controllers

import (
	"encoding/json"

	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/dailing/levlog"
	"html/template"
	"io/ioutil"
	"strconv"
)

type AixintuItemGetController struct {
	beego.Controller
}

func (c *AixintuItemGetController) Post() {
	beego.Debug("get product")
	request := TypeAixinwuItemReqResp{
		Category: -1,
	}
	body := c.Ctx.Input.CopyBody(beego.AppConfig.DefaultInt64("bodybuffer", 1024*1024))
	beego.Info("Post Body is:", string(body))
	err := json.Unmarshal(body, &request)
	beego.Trace("body is :", string(body))
	ErrReport(err)
	if err != nil {
		c.Abort("400")
		return
	}
	response := TypeAixinwuItemReqResp{
		MataData: GenMataData(),
	}
	response.AixinwuItems = GetAixintuItems(request.StartAt, request.Length, request.Category, request.Type)
	c.Data["json"] = response
	c.ServeJSON()
}

type AixintuProductDescriptionRestfulRController struct {
	beego.Controller
}

func (c *AixintuProductDescriptionRestfulRController) Get() {
	beego.Trace("desp for : ", c.Ctx.Input.Param(":productID"))
	c.TplName = "desp.tpl"
	c.Data["id"] = c.Ctx.Input.Param(":productID")

	o := orm.NewOrm()
	id, err := strconv.ParseInt(c.Ctx.Input.Param(":productID"), 10, 64)
	ErrReport(err)
	product := TypeAixinwuProduct{
		Id: int(id),
	}
	o.Read(&product)
	c.Data["desp"] = template.HTML(product.Desc)
	c.Data["baseurl"] = "originalSite"
	c.Data["product"] = product

}

type AixintuProductDescriptionImgRestfulRController struct {
	beego.Controller
}

func (c *AixintuProductDescriptionImgRestfulRController) Get() {
	beego.Trace("desp img for : ", c.Ctx.Input.Param(":productID"), "path : ", c.Ctx.Input.Param(":imgpath"))
	pathPrefix := "/home/sjtu/Sites/test.aixinwu.sjtu.edu.cn/"
	img, err := ioutil.ReadFile(pathPrefix +
		c.Ctx.Input.Param(":imgpath"))
	c.Ctx.Output.Body(img)
	ErrReport(err)
}

type AixintuItemGetRestfulRController struct {
	beego.Controller
}

func (c *AixintuItemGetRestfulRController) Get() {
	itemID, err := strconv.ParseInt(c.Ctx.Input.Param(":productID"), 10, 64)
	levlog.E(err)
	items := GetAixintuItemsByID(int(itemID))
	resp := TypeAixinwuItemReqResp{
		AixinwuItems: items,
		Status:       GenStatus(StatusCodeOK),
	}
	c.Data["json"] = resp
	c.ServeJSON()
}

func (c *AixintuItemGetRestfulRController) Post() {
	beego.Debug("get product")
	request := TypeRegularReq{}
	body := c.Ctx.Input.CopyBody(beego.AppConfig.DefaultInt64("bodybuffer", 1024*1024))
	beego.Info("Post Body is:", string(body))
	err := json.Unmarshal(body, &request)
	beego.Trace("body is :", string(body))
	ErrReport(err)
	if err != nil {
		c.Abort("400")
		return
	}
	tokenInfo := ParseToken(request.Token)
	if tokenInfo.UserID <= 0 {
		c.Abort("400")
	}

	itemID, err := strconv.ParseInt(c.Ctx.Input.Param(":productID"), 10, 64)
	aixinwuID := TransferLocalIDtoAixinwuID(tokenInfo.UserID)
	levlog.E(err)
	items := GetAixintuItemsByID(int(itemID))
	o := orm.NewOrm()
	qs := o.QueryTable(&TypeAixinwuOrder{})
	orders := make([]TypeAixinwuOrder, 0)
	qs.Filter("customer_id", fmt.Sprint(aixinwuID)).All(&orders)
	count := 0
	for _, order := range orders {
		qs := o.QueryTable(&TypeAixinwuOrderItem{})
		orderItems := make([]TypeAixinwuOrderItem, 0)
		num, err := qs.Filter("Order_id", order.Id).Filter("Product_id", items[0].Id).All(&orderItems)
		ErrReport(err)
		for _, i := range orderItems {
			count += i.Quantity
		}
		count += int(num)
	}
	items[0].AlreadyBuy = count
	resp := TypeAixinwuItemReqResp{
		AixinwuItems: items,
		Status:       GenStatus(StatusCodeOK),
	}
	c.Data["json"] = resp
	c.ServeJSON()
}
