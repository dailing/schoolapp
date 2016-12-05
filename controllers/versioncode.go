package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
)

type VersionCodeController struct {
	beego.Controller
}

const versionCodeKey = "AppVersionCode"

func GetDespKey(versionCode string) string {
	return "__VersionDesp__" + versionCode
}

type TypeVersionCode struct {
	VersionCode string `json:"version_code"`
	Desp        string `json:"desp"`
}

func (c *VersionCodeController) Get() {
	resp := TypeVersionCode{
		VersionCode: "0.0.0",
	}
	if ServerParameterHas(versionCodeKey) {
		resp.VersionCode = ServerParameterGet(versionCodeKey)
		resp.Desp = ServerParameterGet(GetDespKey(resp.VersionCode))
	}
	if resp.Desp == "" {
		beego.Error("No desp")
		resp.Desp = "sf"
	}
	c.Data["json"] = resp
	c.ServeJSON()
}

func (c *VersionCodeController) Post() {
	beego.Debug("add user")
	request := TypeVersionCode{}
	body := c.Ctx.Input.CopyBody(beego.AppConfig.DefaultInt64("bodybuffer", 1024*1024))
	beego.Info("Post Body is:", string(body))
	err := json.Unmarshal(body, &request)
	ErrReport(err)
	if err != nil {
		c.Abort("400")
		return
	}
	ServerParameterSet(versionCodeKey, request.VersionCode)
	ServerParameterSet(GetDespKey(request.VersionCode), request.Desp)
	resp := TypeRegularResp{
		Status: GenStatus(StatusCodeOK),
	}
	c.Data["json"] = resp
	c.ServeJSON()
}
