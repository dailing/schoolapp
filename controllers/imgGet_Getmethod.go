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
	imgID := c.Ctx.Input.Param(":imgid")
	if imgID == "" {
		beego.Warning("No image")
		c.Abort("500")
	}
	img, err := ioutil.ReadFile(imgPath + imgID)
	ErrReport(err)
	err = c.Ctx.Output.Body(img)
	ErrReport(err)
}
