package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strconv"
	"time"
)

type OrderProductController struct {
	beego.Controller
}

func (c *OrderProductController) Post() {
	beego.Debug("Make Order")
	request := TypeOrderProductReqResp{}
	body := c.Ctx.Input.CopyBody(beego.AppConfig.DefaultInt64("bodybuffer", 1024*1024))
	beego.Info("Post Body is:", string(body), "Length: ", len(body))
	err := json.Unmarshal(body, &request)
	ErrReport(err)
	if err != nil {
		c.Abort("500")
		//return
	}
	response := TypeOrderProductReqResp{
		MataData: GenMataData(),
		Status:   GenStatus(StatusCodeOK),
	}
	// check token
	tInfo := ParseToken(request.Token)
	if tInfo.UserID <= 0 {
		c.Abort("401")
		return
	}
	usrInfo, err := GetUserInfoByID(fmt.Sprint(tInfo.UserID))
	if err != nil {
		ErrReport(err)
		c.Abort("500")
	}
	// delete stock in product
	o := orm.NewOrm()
	err = o.Begin()
	ErrReport(err)
	if err != nil {
		c.Abort("500")
	}
	var total_price float64
	total_price = 0
	// reserve information fo future usage
	ItemsInfo := make([]InterfaceAixinwuProduct, 0)
	for {
		if usrInfo.JAccount == "" {
			response.Status = GenStatus(StatusNoJaccountInfo)
			break
		}
		jaccountInfo := TypeAixinwuJaccountInfo{
			Jaccount_id: usrInfo.JAccount,
		}
		err = o.Read(&jaccountInfo, "jaccount_id")
		ErrReport(err)
		if err != nil {
			response.Status = GenStatus(StatusCodeDatabaseErr)
		}

		for _, item := range request.OrderInfo {
			var product *TypeAixinwuProduct
			if !item.IsBook {
				//_product := TypeAixinwuProduct{
				//	Id: item.ProductID,
				//}
				//beego.Trace("adding ", item.ProductID)
				//err := o.Read(&_product)
				//ErrReport(err)
				product = &(GetAixintuItemsByID(item.ProductID)[0])
				beego.Error("order", " original ", product.OriginalPrice, " now ", product.GetPrice())
			} else {
				// TODO fix this
				response.Status = GenStatus(StatusCodeNotImplemented)
				beego.Error("Not implementled")
				//product = &TypeAixinwuBook{
				//	ISBN: item.ProductID,
				//}
			}
			ItemsInfo = append(ItemsInfo, product)
			if err != nil {
				ErrReport(err)
				response.Status.Code = StatusCodeUndefinedError
				response.Status.Description = err.Error()
				break
			}
			beego.Trace("Product:", product.GetName(), " id:", product.GetID(), " stock ", product.GetStock(), " querying:", item.Quantity)
			if product.GetStock() < item.Quantity {
				response.Status.Code = StatusCodeUndefinedError
				response.Status.Description = fmt.Sprint("Product:",
					product.GetName(), " id:", product.GetID(),
					" short of stock ", product.GetStock())
				break
			}
			// update stock
			product.SetStock(product.GetStock() - item.Quantity)
			_, err = o.Update(product, "stock")
			if err != nil {
				ErrReport(err)
				response.Status.Code = StatusCodeUndefinedError
				response.Status.Description = err.Error()
				break
			}
			total_price += product.GetPrice() * float64(item.Quantity)
		}
		if usrInfo.Coins < total_price {
			response.Status = GenStatus(StatusCodeNotEnoughMoney)
		} else {
			beego.Trace("custom　ID: ", jaccountInfo.Customer_id)
			cash := TypeAixinwuCustomCash{
				User_id: jaccountInfo.Customer_id,
			}
			err = o.Read(&cash, "user_id")
			ErrReport(err)
			cash.Total -= total_price
			if cash.Total < 0 {
				response.Status = GenStatus(StatusCodeNotEnoughMoney)
				break
			}
			_, err = o.Update(&cash, "total")
			ErrReport(err)
			if err != nil {
				response.Status = GenStatus(StatusCodeDatabaseErr)
				break
			}
		}
		if response.Status.Code != StatusCodeOK {
			break
		}
		// make a record
		address := getAddress(jaccountInfo.Customer_id)
		order := TypeAixinwuOrder{
			Customer_id:         jaccountInfo.Customer_id,
			Total_price:         total_price,
			Total_product_price: total_price,
			Consignee_id:        address.Id,
			Status:              3,
			Place_at:            time.Now(),
			Order_sn:            GenerateRandSN(&TypeAixinwuOrder{}),
		}
		orderID, err := o.Insert(&order)
		response.OrderID = int(orderID)
		if err != nil {
			ErrReport(err)
			response.Status.Code = StatusCodeUndefinedError
			response.Status.Description = err.Error()
			break
		}
		if orderID <= 0 {
			ErrReport("Invilid order id" + fmt.Sprint(orderID))
			response.Status.Code = StatusCodeUndefinedError
			response.Status.Description = "Invilid ID, DB ERROR"
			break
		}

		// make cash record
		cashrecord := TypeAixinwuCashLog{
			Datetime:    time.Now(),
			Customer_id: order.Customer_id,
			Admin_id:    0,
			Change_num:  float64(-order.Total_price),
			Reason:      "支付订单，订单号：" + order.Order_sn,
		}
		_, err = o.Insert(&cashrecord)
		ErrReport(err)

		// add order items to order
		orderIDs := ""
		for index, item := range request.OrderInfo {
			orderItem := TypeAixinwuOrderItem{
				Order_id:     int(orderID),
				Product_id:   fmt.Sprint(ItemsInfo[index].GetID()),
				Product_name: ItemsInfo[index].GetName(),
				Quantity:     item.Quantity,
				Price:        ItemsInfo[index].GetPrice(),
				Weight:       ItemsInfo[index].GetWeight(),
				Category:     ItemsInfo[index].GetCategory(),
			}
			itemID, err := o.Insert(&orderItem)
			if err != nil {
				ErrReport(err)
				response.Status.Code = StatusCodeUndefinedError
				response.Status.Description = err.Error()
				break
			}
			if orderItem.Category == 0 {
				orderIDs += fmt.Sprint(itemID)
			}
			if response.Status.Code != StatusCodeOK {
				break
			}
			// set orderID in item database
			if len(orderIDs) > 0 {
				// the mysql here is copied from original file
				//barcodes := make(map[string]int)
				//for _, product := range ItemsInfo {
				//	pp, succ := product.(*TypeAixinwuProduct)
				//	if !succ {
				//		ErrReport("Things are strange here")
				//		continue
				//	}
				//	barcodes[pp.Barcode] = 1
				//}
				//for barcode, _ := range barcodes {
				pp, succ := ItemsInfo[index].(*TypeAixinwuProduct)
				if !succ {
					ErrReport("Things are strange here")
					continue
				}
				barcode := pp.Barcode
				if barcode == "" {
					continue
				}
				aixinwuItem := TypeAixinwuItem{}
				//beego.Error("select * from lcn_item where barcode =  " + barcode)
				err = o.Raw("select * from lcn_item where barcode =  " + barcode).QueryRow(&aixinwuItem)
				ErrReport(err)
				_iteminfo, err := json.Marshal(&aixinwuItem)
				beego.Trace(string(_iteminfo))
				//err = o.Read(&aixinwuItem, "barcode")
				if err != nil {
					beego.Error("querying lcn_item where barcode = " + aixinwuItem.Barcode)
					ErrReport(err)
					response.Status.Code = StatusCodeUndefinedError
					response.Status.Description = err.Error()
					break
				}
				if aixinwuItem.Order_id == "" {
					aixinwuItem.Order_id = orderIDs
				} else {
					aixinwuItem.Order_id += "," + orderIDs
				}
				_, err = o.Update(&aixinwuItem, "order_id")
				if err != nil {
					ErrReport(err)
					response.Status.Code = StatusCodeUndefinedError
					response.Status.Description = err.Error()
					break
				}
				//}
			}
		}
		// make sure there is a break at the end of this for loop
		// this for loop is only executed once, for convenience of jumping out
		break
	}
	if response.Status.Code != StatusCodeOK {
		beego.Trace("Rolling back")
		o.Rollback()
	} else {
		o.Commit()
	}
	// TODO add item to order
	// set parameters
	c.Data["json"] = response
	c.ServeJSON()
}

type AixinwuOrderGetController struct {
	beego.Controller
}

func (c *AixinwuOrderGetController) Get() {
	strId := c.Ctx.Input.Param(":uid")
	strSart := c.Ctx.Input.Param(":start")
	strLength := c.Ctx.Input.Param(":len")
	id, _ := strconv.ParseInt(strId, 10, 64)
	start, _ := strconv.ParseInt(strSart, 10, 64)
	length, _ := strconv.ParseInt(strLength, 10, 64)
	retval := c.GetList(int(id), int(length), int(start))
	c.Data["json"] = retval
	c.ServeJSON()
}

func (c *AixinwuOrderGetController) Post() {
	beego.Debug("get Order")
	body := c.Ctx.Input.CopyBody(beego.AppConfig.DefaultInt64("bodybuffer", 1024*1024))
	beego.Info("Post Body is:", string(body), "Length: ", len(body))
	request := TypeAixinwuOrderGetReq{}
	err := json.Unmarshal(body, &request)
	ErrReport(err)
	if err != nil {
		c.Abort("400")
	}
	tokeninfo := ParseToken(request.Token)
	if tokeninfo.UserID <= 0 {
		beego.Warn("token err:", request.Token)
		c.Abort("400")
	}
	response := TypeAixinwuOrderGetResp{
		Status: GenStatus(StatusCodeOK),
		Orders: c.GetList(TransferLocalIDtoAixinwuID(tokeninfo.UserID), request.Length, request.Offset),
	}
	c.Data["json"] = response
	c.ServeJSON()
}

func (c *AixinwuOrderGetController) GetList(id int,
	length int,
	offset int) []TypeAixinwuOrder {
	beego.Trace("Getting Order List len:",
		length, " offset ",
		offset,
	)
	o := orm.NewOrm()
	qs := o.QueryTable(&TypeAixinwuOrder{})
	qs = qs.OrderBy("-id").
		Filter("is_delete", 0).
		Filter("customer_id", id).
		Offset(offset).
		Limit(length)
	retval := make([]TypeAixinwuOrder, 0)
	_, err := qs.All(&retval)
	if err != nil {
		ErrReport(err)
		return nil
	}
	for index, _ := range retval {
		retval[index].Items = getItems(retval[index].Id)
	}
	return retval
}

type AixinwuOrderItemGetController struct {
	beego.Controller
}

func (c *AixinwuOrderItemGetController) Post() {
	beego.Debug("get Order")
	body := c.Ctx.Input.CopyBody(beego.AppConfig.DefaultInt64("bodybuffer", 1024*1024))
	beego.Info("Post Body is:", string(body), "Length: ", len(body))
	request := TypeAixinwuOrderItemGetReq{}
	response := TypeAixinwuOrderItemResp{
		Status: GenStatus(StatusCodeOK),
	}
	err := json.Unmarshal(body, &request)
	ErrReport(err)
	if err != nil {
		c.Abort("400")
	}
	for {
		o := orm.NewOrm()
		tokeninfo := ParseToken(request.Token)

		orderInfo := TypeAixinwuOrder{
			Id: request.OrderID,
		}
		err = o.Read(&orderInfo)
		if err != nil {
			ErrReport(err)
			response.Status = GenStatus(StatusCodeDatabaseErr)
			break
		}
		if tokeninfo.UserID != orderInfo.Customer_id {
			beego.Warn("custumos id and order id not match:",
				tokeninfo.UserID, " ,, ",
				orderInfo.Customer_id,
			)
			response.Status = GenStatus(StatusCodeUndefinedError)
			break
		}
		response.Items = getItems(request.OrderID)
		break
	}
	c.Data["json"] = response
	c.ServeJSON()
}

func getItems(orderid int) []TypeAixinwuOrderItem {
	o := orm.NewOrm()
	qs := o.QueryTable(&TypeAixinwuOrderItem{})
	qs = qs.Filter("order_id", orderid)
	retval := make([]TypeAixinwuOrderItem, 0)
	qs.All(&retval)
	for index := range retval {
		images := make([]TypeAixinwuProductImage, 0)
		_, err := o.QueryTable("lcn_product_image").
			Filter("product_id", retval[index].Product_id).
			All(&images)
		ErrReport(err)
		if err != nil || len(images) == 0 {
			continue
		}
		imageStr := ""
		for _, imgs := range images {
			imageStr += "img/" + imgs.File + ","
		}
		imageStr = imageStr[:len(imageStr)-1]
		retval[index].Image = imageStr
	}
	return retval
}
