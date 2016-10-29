package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type AixinwuVolunteerActGetController struct {
	beego.Controller
}

func (c *AixinwuVolunteerActGetController) Post() {
	beego.Debug("Volunteer join")
	body := c.Ctx.Input.CopyBody(beego.AppConfig.DefaultInt64("bodybuffer", 1024*1024))
	beego.Info("Post Body is:", string(body), "Length: ", len(body))
	request := TypeRegularReq{}
	err := json.Unmarshal(body, &request)
	ErrReport(err)
	tokeninfo := ParseToken(request.Token)

	response := make([]TypeAixinwuVolunteerAct, 0)
	o := orm.NewOrm()
	qs := o.QueryTable(&TypeAixinwuVolunteerAct{})
	qs = qs.Filter("status", 1) // 1 for open dactivity
	qs.All(&response)

	for index, _ := range response {
		if response[index].Poster != "" {
			response[index].Image = "http://test.aixinwu.info/uploads/" +
				response[index].Poster
		}
		volunteer := TypeAixinwuVolunteer{
			Uid:        TransferLocalIDtoAixinwuID(tokeninfo.UserID),
			Project_id: response[index].Id,
		}
		err := o.Read(&volunteer, "uid", "project_id")
		beego.Trace("checking ", volunteer.Uid,
			"  ", volunteer.Project_id,
		)
		if err == orm.ErrNoRows {
			response[index].Joined = 0
		} else {
			response[index].Joined = 1
			ErrReport(err)
		}
	}
	c.Data["json"] = response
	c.ServeJSON()
}

type AixinwuVolunteerActJoinController struct {
	beego.Controller
}

func (c *AixinwuVolunteerActJoinController) Post() {
	beego.Debug("Volunteer join")
	body := c.Ctx.Input.CopyBody(beego.AppConfig.DefaultInt64("bodybuffer", 1024*1024))
	beego.Info("Post Body is:", string(body), "Length: ", len(body))
	request := TypeAixinwuVolunteerJoinReq{}
	response := TypeRegularResp{
		Status: GenStatus(StatusCodeOK),
	}
	err := json.Unmarshal(body, &request)
	ErrReport(err)
	if err != nil {
		c.Abort("400")
	}
	tokenInfo := ParseToken(request.Token)

	o := orm.NewOrm()
	act := TypeAixinwuVolunteerAct{
		Id: request.Project_id,
	}
	err = o.Read(&act)
	ErrReport(err)
	if err != nil {
		c.Abort("400")
	}
	for {
		if act.Num_signed >= act.Num_needed {
			response.Status = GenStatus(StatusCodeNotEnoughMoney)
			break
		}
		if tokenInfo.UserID <= 0 {
			response.Status = GenStatus(StatusCodeErrorLoginInfo)
			break
		}
		record := TypeAixinwuVolunteer{
			Uid:        TransferLocalIDtoAixinwuID(tokenInfo.UserID),
			Project_id: request.Project_id,
			Project:    act.Name,
			Work_date:  act.Work_date,
			Workload:   float64(act.Workload),
			Content:    act.Content,
			Tel:        request.Tel,
			Pay_cash:   act.Pay_cash,
		}
		act.Num_signed += 1
		_, err = o.Insert(&record)
		ErrReport(err)
		if err != nil {
			response.Status = GenStatus(StatusCodeDatabaseErr)
			break
		}
		_, err = o.Update(&act)
		if err != nil {
			response.Status = GenStatus(StatusCodeDatabaseErr)
			break
		}
		break
	}
	c.Data["json"] = response
	c.ServeJSON()
}
