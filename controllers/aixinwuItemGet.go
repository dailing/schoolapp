package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"html/template"
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
	ErrReport(err)
	if err != nil {
		c.Abort("400")
		return
	}
	response := TypeAixinwuItemReqResp{
		MataData: GenMataData(),
	}
	response.AixinwuItems = GetAixintuItems(request.StartAt, request.Length, request.Category)
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
}
