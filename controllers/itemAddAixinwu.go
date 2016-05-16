package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"math/rand"
	"time"
	"unicode/utf8"
)

type ItemAddAixinwuController struct {
	beego.Controller
}

func getDonationSN() string {
	t := time.Now()
	ret := ""
	ret += fmt.Sprintf("%04d%02d%02d", t.Year(), t.Month(), t.Day())
	ret += fmt.Sprintf("%05d", rand.Int()%100000)
	return ret
}

/*
 * The barcode should be a string of length 9
 * YYMMiiiii
 * 34.""
 */
func genBarcode() string {
	const modVal = 10000000
	o := orm.NewOrm()
	var info TypeLcnDonateBatch
	var code = 0
	for {
		code = rand.Int() % modVal
		//code += 1
		info.Barcode = fmt.Sprintf("34%07d", code)
		err := o.Read(&info, "barcode")
		beego.Debug("trying bar code ", info.Barcode)
		ErrReport(err)
		if err == orm.ErrNoRows {
			break
		}
	}
	beego.Debug("succeed with barcode ", info.Barcode)
	return info.Barcode
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
	jinfo := TypeAixinwuJaccountInfo{
		Jaccount_id: request.Item.JAcountID,
	}
	err = o.Read(&jinfo, "jaccount_id")
	ErrReport(err)

	// set parameters
	succ := utf8.Valid([]byte(request.Item.Desc))
	beego.Info(succ)
	//request.Item.Desc = baseEncode(request.Item)
	dinfo := TypeLcnDonateBatch{
		Donation_sn: getDonationSN(),
		Barcode:     genBarcode(),
		User_id:     jinfo.Customer_id,
		Snum:        jinfo.Snum,
		Desc:        request.Item.Desc,
		Status:      1,
	}
	donationID, err := o.Insert(&dinfo)
	ErrReport(err)
	// TODO also add this information to item database
	itemInfo := TypeAixinwuItem{
		Barcode:     dinfo.Barcode,
		Status:      1,
		Donation_id: int(donationID),
		// TODO  change this back
		//Valuation:   request.Item.Valuation,
		Valuation:   123,
		Name:        request.Item.Desc,
		Description: request.Item.Desc,
		Is_delete:   0,
		Quantity:    1,
		Validity:    time.Now(),
	}
	_, err = o.Insert(&itemInfo)
	ErrReport(err)
	response.Status = GenStatus(StatusCodeOK)
	c.Data["json"] = response
	c.ServeJSON()
}
