package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
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
				_product := TypeAixinwuProduct{
					Id: item.ProductID,
				}
				beego.Trace("adding ", item.ProductID)
				err := o.Read(&_product)
				ErrReport(err)
				product = &_product
			} else {
				// TODO fix this
				response.Status = GenStatus(StatusCodeNotImplemented)
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
		order := TypeAixinwuOrder{
			Customer_id:         jaccountInfo.Customer_id,
			Total_price:         total_price,
			Total_product_price: total_price,
			Consignee_id:        jaccountInfo.Customer_id,
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
