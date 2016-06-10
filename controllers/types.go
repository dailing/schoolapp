package controllers

import "time"

type TypeMataData struct {
	TimeStamp int
	Device    string
}

type TypeRegularResp struct {
	MataData TypeMataData `json:"mataData"`
	Status   TypeStatus   `json:"status"`
}

type TypeRegularReq struct {
	MataData TypeMataData `json:"mataData"`
	Token    string       `json:"token"`
}

// regular user request
type TypeUserReq struct {
	MataData TypeMataData `json:"mataData"`
	UserInfo TypeUserInfo `json:"userinfo"`
	Token    string       `json:"token"`
	Status   TypeStatus   `json:"status"`
}

// image related request parameters
type TypeImgResp struct {
	MataData TypeMataData `json:"mataData"`
	ImageID  string       `json:"imageID"`
	Token    string       `json:"token"`
	Status   TypeStatus   `json:"status"`
}

type TypeUserInfo struct {
	ID       int     `json:"ID"       orm:"pk;auto;column(Uid)"`
	Username string  `json:"username" orm:"type(text);unique;column(username)"`
	Password string  `json:"password" orm:"column(password)"`
	NickName string  `json:"nickname" orm:"type(text);column(nickname)"`
	Phone    string  `json:"phone"    orm:"column(phone)"`
	Email    string  `json:"email"    orm:"column(email)"`
	Coins    float64 `json:"coins"    orm:"column(coins)"`
	JAccount string  `json:"jaccount" orm:"column(jaccount)"`
}

/*
 * item related structure definition
 * Status :
 *		000: 初始状态
 *		200:已审批
 */
type TypeItemInfo struct {
	ID                      int       `json:"ID"                       orm:"pk;auto;(id)"`
	Caption                 string    `json:"caption"                  orm:"(caption);type(text);null"`
	BoughtAt                time.Time `json:"boughtAt"                 orm:"(boughtAt);null"`
	ItemCondition           int       `json:"itemCondition"            orm:"(itemCondition);null"`
	EstimatedPriceByUser    int       `json:"estimatedPriceByUser"     orm:"(estimatedPriceByUser);null"`
	EstimatedPriceByAiXinWu int       `json:"estimatedPriceByAiXinWu"  orm:"(estimatedPriceByAiXinWu);null"`
	Location                string    `json:"location"                 orm:"(location);type(text);null"`
	Category                int       `json:"category"                 orm:"(category);null"`
	PublishedAt             time.Time `json:"publishedAt"              orm:"(publishedAt);null"`
	UnShelveReason          string    `json:"unShelveReason"           orm:"(unShelveReason);type(text);null"`
	Description             string    `json:"description"              orm:"(description);type(text);null"`
	OwnerID                 int       `json:"ownerID"                  orm:"(ownerID);null"`
	Price                   int       `json:"price"                    orm:"(price);null"`
	Images                  string    `json:"images"                   orm:"(images);type(text);null"`
	Status                  int       `json:"status"                   orm:"(status);null"`
}

type TypeItemReqResp struct {
	MataData TypeMataData `json:"mataData"`
	Token    string       `json:"token"`
	ItemInfo TypeItemInfo `json:"itemInfo"`
	Status   TypeStatus   `json:"status"`
}

type TypeItemGetAllReq struct {
	MataData TypeMataData `json:"mataData"`
	Token    string       `json:"token"`
	StartAt  int          `json:"startAt"`
	Length   int          `json:"length"`
	Status   TypeStatus   `json:"status"`
}

type TypeGetItemsResp struct {
	MataData TypeMataData   `json:"mataData"`
	Items    []TypeItemInfo `json:"items"`
	Status   TypeStatus     `json:"status"`
}

/*
 *  	Comments related structure definition
 */
type TypeItemComments struct {
	ID          int       `json:"ID"            orm:"pk;auto;(id)"`
	Content     string    `json:"content"       orm:"(content);type(text)"`
	ItemId      int       `json:"itemID"        orm:"(itemID)"`
	PublisherID int       `json:"publisherID"   orm:"(publisherID)"`
	Created     time.Time `json:"created"       orm:"auto_now_add;type(datetime)"`
}

type TypeCommentReq struct {
	MataData TypeMataData     `json:"mataData"`
	Token    string           `json:"token"`
	Comment  TypeItemComments `json:"comment"`
	Status   TypeStatus       `json:"status"`
}

type TypeCommentResp struct {
	MataData TypeMataData       `json:"mataData"`
	Token    string             `json:"token"`
	Comments []TypeItemComments `json:"comment"`
	Status   TypeStatus         `json:"status"`
}

/*
 *	Chat related definition
 */
type TypeChatInfo struct {
	ID          int       `json:"ID"            orm:"pk;auto;(id)"`
	Content     string    `json:"content"       orm:"(content);type(text)"`
	ItemID      int       `json:"itemID"        orm:"(item_id)"`
	BuyerID     int       `json:"buyer_id"      orm:"(buyer_id)"`
	PublisherID int       `json:"publisher_id"  orm:"(publisher_id)"`
	Created     time.Time `json:"created"       orm:"auto_now_add;type(datetime)"`
}

type TypeChatReq struct {
	MataData TypeMataData `json:"mataData"`
	Token    string       `json:"token"`
	Chat     TypeChatInfo `json:"chat"`
}

type TypeChatResp struct {
	MataData TypeMataData `json:"mataData"`
	//Token    string         `json:"token"`
	Chat   []TypeChatInfo `json:"chat"`
	Status TypeStatus     `json:"status"`
}

/*
 * 	Item add aixinwu
 */
type TypeItemAixinwuInfo struct {
	JAcountID string `json:"jacount_id"`
	Desc      string `json:"desc"`
	Valuation int    `json:"valuation"`
}
type TypeItemAixinwuReq struct {
	MataData TypeMataData        `json:"mataData"`
	Token    string              `json:"token"`
	Item     TypeItemAixinwuInfo `json:"itemInfo"`
	Status   TypeStatus          `json:"status"`
}

// ln_jacount_info
//type TypeLcnJacountInfo struct {
//	Id          int    `json:"id"            orm:"pk;auto;(id)"`
//	Customer_id int    `json:"customer_id"   orm:"(customer_id)"`
//	Jaccount_id string `json:"jaccount_id"   orm:"(jaccount_id)"`
//	Citizenid   string `json:"citizenid"     orm:"(citizenid)"`
//	Realname    string `json:"realname"      orm:"(realname)"`
//	Dept        string `json:"dept"          orm:"(dept)"`
//	Tel         string `json:"tel"           orm:"(tel)"`
//	Snum        string `json:"snum"          orm:"(snum)"`
//	Is_student  string `json:"is_student"    orm:"(is_student)"`
//}
//
//func (u *TypeLcnJacountInfo) TableName() string {
//	return "lcn_jaccount_info"
//}

type TypeLcnDonateBatch struct {
	Id          int       `json:"id"             orm:"pk;auto;(id)"`
	User_id     int       `json:"user_id"        orm:"(user_id)"`
	Snum        string    `json:"snum"           orm:"(snum)"`
	Produced_at time.Time `json:"produced_at"    orm:"auto_now_add;type(datetime);(produced_at)"`
	Desc        string    `json:"desc"    orm:"  (desc)"`
	Donation_sn string    `json:"donation_sn"    orm:"(donation_sn)"`
	Barcode     string    `json:"barcode"        orm:"(barcode)"`
	Status      int       `json:"status"         orm:"(status)"`
}

func (u *TypeLcnDonateBatch) TableName() string {
	return "lcn_donation_batch"
}

/*
 * 	Aixinwu product databese
 */
//Table: lcn_product
type TypeAixinwuProduct struct {
	Id                     int       `json:"id"                      orm:"pk;auto;(id)"`
	Cat_id                 int       `json:"cat_id"                  orm:"(cat_id)"`
	Brand_id               int       `json:"brand_id"                orm:"(brand_id)"`
	Attr_set_id            int       `json:"attr_set_id"             orm:"(attr_set_id)"`
	Price                  int       `json:"price"                   orm:"(price)"`
	Market_price           int       `json:"market_price"            orm:"(market_price)"`
	Special_price          int       `json:"special_price"           orm:"(special_price)"`
	Name                   string    `json:"name"                    orm:"(name)"`
	Short_name             string    `json:"short_name"              orm:"(short_name)"`
	Url_alias              string    `json:"url_alias"               orm:"(url_alias)"`
	Short_desc             string    `json:"short_desc"              orm:"(short_desc)"`
	Desc                   string    `json:"desc"                    orm:"(desc)"`
	Weight                 int       `json:"weight"                  orm:"(weight)"`
	Stock                  int       `json:"stock"                   orm:"(stock)"`
	Limit                  int       `json:"limit"                   orm:"(limit)"`
	Is_on_sale             int       `json:"is_on_sale"              orm:"(is_on_sale)"`
	On_sale_at             time.Time `json:"on_sale_at"              orm:"(on_sale_at)"`
	Tag                    string    `json:"tag"                     orm:"(tag)"`
	Meta_title             string    `json:"meta_title"              orm:"(meta_title)"`
	Meta_keywords          string    `json:"meta_keywords"           orm:"(meta_keywords)"`
	Meta_desc              string    `json:"meta_desc"               orm:"(meta_desc)"`
	Is_new                 int       `json:"is_new"                  orm:"(is_new)"`
	Is_hot                 int       `json:"is_hot"                  orm:"(is_hot)"`
	Is_special_price       int       `json:"is_special_price"        orm:"(is_special_price)"`
	Special_price_start_at time.Time `json:"special_price_start_at"  orm:"(special_price_start_at)"`
	Special_price_end_at   time.Time `json:"special_price_end_at"    orm:"(special_price_end_at)"`
	Is_commend             int       `json:"is_commend"              orm:"(is_commend)"`
	Is_delete              int       `json:"is_delete"               orm:"(is_delete)"`
	Created_at             time.Time `json:"created_at"              orm:"(created_at)"`
	Updated_at             time.Time `json:"updated_at"              orm:"(updated_at)"`
	Code                   string    `json:"code"                    orm:"(code)"`
	Barcode                string    `json:"barcode"                 orm:"(barcode)"`
	Is_borrow              int       `json:"is_borrow"               orm:"(is_borrow)"`
	Is_cash                int       `json:"is_cash"                 orm:"(is_cash)"`
}

func (u *TypeAixinwuProduct) TableName() string {
	return "lcn_product"
}

type TypeAixinwuItem struct {
	Id             int       `json:"id"               orm:"pk;auto;(id)"`
	Barcode        string    `json:"barcode"          orm:"(barcode)"`
	Name           string    `json:"name"             orm:"(name)"`
	Valuation      int       `json:"valuation"        orm:"(valuation)"`
	Status         int       `json:"status"           orm:"(status)"`
	Quantity       int       `json:"quantity"         orm:"(quantity)"`
	Quantity_saled int       `json:"quantity_saled"   orm:"(quantity_saled)"`
	Category       int       `json:"category"         orm:"(category)"`
	Donation_id    int       `json:"donation_id"      orm:"(donation_id)"`
	Product_id     int       `json:"product_id"       orm:"(product_id)"`
	Order_id       string    `json:"order_id"         orm:"(order_id)"`
	Description    string    `json:"description"      orm:"(description)"`
	Create_time    time.Time `json:"create_time"      orm:"(create_time)"`
	Validity       time.Time `json:"validity"         orm:"(validity)"`
	Is_delete      int       `json:"is_delete"        orm:"(is_delete)"`
	Image_name     string    `json:"image_name"       orm:"(image_name)"`
}

func (u *TypeAixinwuItem) TableName() string {
	return "lcn_item"
}

type TypeAixinwuJaccountInfo struct {
	Id          int    `json:"id"            orm:"pk;auto;(id)"`
	Customer_id int    `json:"customer_id"   orm:"(customer_id)"`
	Jaccount_id string `json:"jaccount_id"   orm:"(jaccount_id)"`
	Citizenid   string `json:"citizenid"     orm:"(citizenid)"`
	Realname    string `json:"realname"      orm:"(realname)"`
	Dept        string `json:"dept"          orm:"(dept)"`
	Tel         string `json:"tel"           orm:"(tel)"`
	Snum        string `json:"snum"          orm:"(snum)"`
	Is_student  string `json:"is_student"    orm:"(is_student)"`
}

func (u *TypeAixinwuJaccountInfo) TableName() string {
	return "lcn_jaccount_info"
}

type TypeAixinwuCustomCash struct {
	User_id int     `json:"user_id"  orm:"pk;(user_id)"`
	Total   float64 `json:"total"    orm:"(total)"`
}

func (u *TypeAixinwuCustomCash) TableName() string {
	return "lcn_customer_cash"
}

/*
 * Static infomationo
 */
type TypeStaticInfo struct {
	Visit string `json:"visitCounter"`
	Money string `json:"money"`
	User  string `json:"user"`
	Item  string `json:"item"`
}
type TypeStaticInfoResp struct {
	MataData   TypeMataData   `json:"mataData"`
	StaticInfo TypeStaticInfo `json:"staticInfo"`
	Status     TypeStatus     `json:"status"`
}

/*
 *  The six picture in the main page of app
 */
type TypeMainPageItems struct {
	ID            int `json:"id"             orm:"(id)"`
	ItemID        int `json:"itemID"         orm:"(itemID)"`
	IsAiXinWuItem int `json:"isAixinwuItem"  orm:"(isAixinwuItem)"`
}

/*
 * 	Error code and other definitions
 */
type TypeStatus struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
}

type TypeTokenInfo struct {
	UserName string
	UserID   int
}

type TypeParameters struct {
	ID           int    `json:"homePageItem"        orm:"pk;(homePageItem)"`
	HomePageItem string `json:"homePageItem"        orm:"(homePageItem)"`
}

type TypeParametersRwqResp struct {
	MataData   TypeMataData   `json:"mataData"`
	Parameters TypeParameters `json:"parameter"`
	Status     TypeStatus     `json:"status"`
}

const (
	StatusCodeOK             = iota
	StatusCodeErrorLoginInfo = iota
)

var ErrorDesp = map[int]string{
	StatusCodeOK:             "OK",
	StatusCodeErrorLoginInfo: "Wrong username or password",
}
