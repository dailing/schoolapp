package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"math/rand"
	"time"
)

type ItemAddAixinwuController struct {
	beego.Controller
}

func getDonationSN() string {
	t := time.Now()
	ret := t.Format("20060122")
	ret += fmt.Sprintf("%05d", rand.Int()%100000)
	return ret
}

func (c *ItemAddAixinwuController) Post() {
	beego.Debug("Aixinwu item")
	request := TypeItemAixinwuReq{}
	body := c.Ctx.Input.CopyBody(beego.AppConfig.DefaultInt64("bodybuffer", 1024*1024))
	beego.Info("Post Body is:", string(body))
	err := json.Unmarshal(body, &request)
	tbody, err := json.Marshal(request)
	beego.Trace(string(tbody))
	ErrReport(err)
	if err != nil {
		c.Abort("500")
		return
	}
	response := TypeRegularResp{
		MataData: GenMataData(),
	}
	// check token
	tInfo := ParseToken(request.Token)
	if tInfo.UserID <= 0 {
		c.Abort("400")
		return
	}
	// get jacount info
	o := orm.NewOrm()
	jinfo := TypeLcnJacountInfo{
		Jaccount_id: request.Item.JAcountID,
	}
	err = o.Read(&jinfo, "jaccount_id")
	ErrReport(err)

	// set parameters
	dinfo := TypeLcnDonateBatch{
		Donation_sn: getDonationSN(),
		User_id:     jinfo.Customer_id,
		Snum:        jinfo.Snum,
		Desc:        request.Item.Desc,
		Status:      1,
	}
	_, err = o.Insert(&dinfo)
	ErrReport(err)
	response.Status = GenStatus(StatusCodeOK)
	c.Data["json"] = response
	c.ServeJSON()
}
