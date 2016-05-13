package controllers

import (
	"github.com/astaxie/beego"
	"io/ioutil"
)

type SearchRestfulRController struct {
	beego.Controller
}

func (c *SearchRestfulRController) Get() {
	beego.Trace("search resuful, the type is:", c.Ctx.Input.Param(":searchField"))

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
