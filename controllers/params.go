package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strconv"
	"strings"
)

type ParamsGet struct {
	beego.Controller
}

func (c *ParamsGet) Post() {
	beego.Debug("add user")
	request := TypeParametersRwqResp{}
	//body := c.Ctx.Input.CopyBody(beego.AppConfig.DefaultInt64("bodybuffer", 1024*1024))
	//beego.Info("Post Body is:", string(body))
	//err := json.Unmarshal(body, &request)
	//ErrReport(err)
	//if err != nil {
	//	c.Abort("400")
	//	return
	//}
	response := TypeGetItemsResp{
		MataData: GenMataData(),
		Status:   GenStatus(StatusCodeOK),
	}

	o := orm.NewOrm()
	request.Parameters.ID = 1
	o.Read(&request.Parameters)

	// ser parameters
	strs := strings.Split(request.Parameters.HomePageItem, ",")
	beego.Info(request.Parameters.HomePageItem)
	response.Items = make([]TypeItemInfo, 0)
	for _, str := range strs {
		id, err := strconv.ParseInt(str, 10, 32)
		ErrReport(err)
		beego.Debug("Adding ", id)
		if err == nil {
			item, err := GetItemByID(int(id))
			ErrReport(err)
			if err == nil {
				response.Items = append(response.Items, item)
			}
			beego.Trace(item)
		}
	}
	b, _ := json.Marshal(response)
	beego.Trace(string(b))
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
