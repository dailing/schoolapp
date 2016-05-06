package controllers

import (
	"github.com/astaxie/beego"
)

type StaticsGetController struct {
	beego.Controller
}

func (c *StaticsGetController) Post() {
	beego.Debug("static get controller")
	response := TypeStaticInfoResp{
		MataData: GenMataData(),
	}
	// get info
	

	// return
	response.Status = GenStatus(StatusCodeOK)
	c.Data["json"] = response
	c.ServeJSON()
}
