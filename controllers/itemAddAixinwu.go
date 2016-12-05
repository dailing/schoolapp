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
		info.Barcode = fmt.Sprintf("54%07d", code)
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
	response := TypeRegularResp{}
	// check token
	tInfo := ParseToken(request.Token)
	if tInfo.UserID <= 0 {
		c.Abort("400")
		return
	}
	var itemInfo TypeAixinwuItem
	for {
		// get jacount info
		o := orm.NewOrm()
		jinfo := TypeAixinwuJaccountInfo{
			Jaccount_id: request.Item.JAcountID,
		}
		err = o.Read(&jinfo, "jaccount_id")
		ErrReport(err)
		if err != nil {
			break
		}

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
		if err != nil {
			break
		}
		itemInfo = TypeAixinwuItem{
			Barcode:     dinfo.Barcode,
			Status:      1,
			Donation_id: int(donationID),
			Valuation:   request.Item.Valuation,
			//Valuation:   123,
			Name:        request.Item.Desc,
			Description: request.Item.Desc,
			Is_delete:   0,
			Quantity:    1,
			Validity:    time.Now(),
		}
		_, err = o.Insert(&itemInfo)
		ErrReport(err)
		if err != nil {
			break
		}
		break
	}
	response.Status = GenStatus(StatusCodeOK)
	if err != nil {
		response.Status = GenStatus(StatusCodeDatabaseErr)
	}
	c.Data["json"] = struct {
		Status   TypeStatus      `json:"status"`
		ItemInfo TypeAixinwuItem `json:"item_info"`
	}{
		Status:   response.Status,
		ItemInfo: itemInfo,
	}
	c.ServeJSON()
}

type DonateRecordGetController struct {
	beego.Controller
}

func (c *DonateRecordGetController) Post() {
	beego.Debug("Aixinwu item")
	request := TypeAixinwuDonateRecordGetReq{}
	body := c.Ctx.Input.CopyBody(beego.AppConfig.DefaultInt64("bodybuffer", 1024*1024))
	beego.Info("Post Body is:", string(body))
	err := json.Unmarshal(body, &request)
	ErrReport(err)
	if err != nil {
		c.Abort("400")
	}
	tokenInfo := ParseToken(request.Token)
	if tokenInfo.UserID <= 0 {
		c.Abort("400")
	}
	aixinwuID := TransferLocalIDtoAixinwuID(tokenInfo.UserID)
	o := orm.NewOrm()
	records := make([]TypeLcnDonateBatch, 0)
	qs := o.QueryTable(&TypeLcnDonateBatch{})
	_, err = qs.Filter("user_id", aixinwuID).
		Offset(request.Offset).
		Limit(request.Length).
		All(&records)
	ErrReport(err)
	response := TypeAixinwuDonateRecordGetResp{
		Records: records,
	}
	if err == nil {
		response.Status = GenStatus(StatusCodeOK)
	} else {
		response.Status = GenStatus(StatusCodeDatabaseErr)
	}
	c.Data["json"] = response
	c.ServeJSON()
}
