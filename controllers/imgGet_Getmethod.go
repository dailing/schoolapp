package controllers

import (
	"github.com/astaxie/beego"
	"io/ioutil"
)

type ImgGetRestfulRController struct {
	beego.Controller
}

func (c *ImgGetRestfulRController) Get() {
	beego.Trace("img get resuful, the id is:", c.Ctx.Input.Param(":imgid"))
	var img []byte
	var err error
	for {
		if c.Ctx.Input.Param(":productid") != "" {
			beego.Trace("Img get product image, product ID : " + c.Ctx.Input.Param(":productid"))
			img, err = ioutil.ReadFile("/home/d/test.aixinwu.sjtu.edu.cn/uploads/product/" +
				c.Ctx.Input.Param(":productid") +
				"/" +
				c.Ctx.Input.Param(":imgid"))
			ErrReport(err)
			break
		}
		imgID := c.Ctx.Input.Param(":imgid")
		if imgID == "" {
			beego.Warning("No image")
			c.Abort("500")
		}
		for _, ipath := range imgPath {
			img, err = ioutil.ReadFile(ipath + imgID)
			if err == nil {
				break
			}
		}
		ErrReport(err)
		break
	}
	err = c.Ctx.Output.Body(img)
	ErrReport(err)
}
