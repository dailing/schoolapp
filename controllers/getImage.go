package controllers

import "github.com/astaxie/beego"

type GetImageController struct {
	beego.Controller
}

func (c *GetImageController) Post() {
	beego.Info("GetImage")
}
