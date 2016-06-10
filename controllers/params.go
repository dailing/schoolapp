package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type ParamsGet struct {
	beego.Controller
}

func (c *ParamsGet) Post() {
	beego.Debug("add user")
	request := TypeParametersRwqResp{}
	body := c.Ctx.Input.CopyBody(beego.AppConfig.DefaultInt64("bodybuffer", 1024*1024))
	beego.Info("Post Body is:", string(body))
	err := json.Unmarshal(body, &request)
	ErrReport(err)
	if err != nil {
		c.Abort("400")
		return
	}
	response := TypeParametersRwqResp{
		MataData: GenMataData(),
	}
	// ser parameters
	o := orm.NewOrm()
	response.Parameters.ID = 1
	o.Read(&response.Parameters)
	c.Data["json"] = response
	c.ServeJSON()
}

type ParamsSet struct {
	beego.Controller
}

func (c *ParamsSet) Post() {
	beego.Debug("add user")
	request := TypeParametersRwqResp{}
	body := c.Ctx.Input.CopyBody(beego.AppConfig.DefaultInt64("bodybuffer", 1024*1024))
	beego.Info("Post Body is:", string(body))
	err := json.Unmarshal(body, &request)
	ErrReport(err)
	if err != nil {
		c.Abort("400")
		return
	}
	//　TODO Check the user authorization
	response := TypeParametersRwqResp{
		MataData: GenMataData(),
	}
	// ser parameters
	o := orm.NewOrm()
	request.Parameters.ID = 1
	o.Update(request)
	c.Data["json"] = response
	c.ServeJSON()
}
