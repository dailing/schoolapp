package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
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
