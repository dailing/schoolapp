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
	ID       int    `json:"ID"       orm:"pk;auto;column(Uid)"`
	Username string `json:"username" orm:"unique;column(username)"`
	Password string `json:"password" orm:"column(password)"`
	NickName string `json:"nickname" orm:"column(nickname)"`
	Phone    string `json:"phone"    orm:"column(phone)"`
	Email    string `json:"email"    orm:"column(email)"`
	Coins    int    `josn:"coins"    orm:"column(coins)"`
}

/*
 * item related structure definition
 */
type TypeItemInfo struct {
	ID          int    `json:"ID"            orm:"pk;auto;colume(id)"`
	OwnerID     int    `json:"ownerID"       orm:"colume(ownerID)"`
	Description string `json:"description"   orm:"colume(description)"`
	Price       int    `json:"price"         orm:"colume(price)"`
	Images      string `json:"images"        orm:"colume(images)"`
	Status      int    `json:"status"        orm:"colume(status)"`
}

type TypeItemReqResp struct {
	MataData TypeMataData `json:"mataData"`
	Token    string       `json:"token"`
	ItemInfo TypeItemInfo `json:"itemInfo"`
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
	ID          int       `json:"ID"            orm:"pk;auto;colume(id)"`
	Content     string    `json:"content"       orm:"colume(content)"`
	ItemId      int       `json:"itemID"        orm:"colume(itemID)"`
	PublisherID int       `json:"publisherID"   orm:"colume(publisherID)"`
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
	ID          int       `json:"ID"            orm:"pk;auto;colume(id)"`
	Content     string    `json:"content"       orm:"colume(content)"`
	ItemID      int       `json:"itemID"        orm:"colume(item_id)"`
	BuyerID     int       `json:"buyerID"       orm:"colume(buyer_id)"`
	PublisherID int       `json:"publisherID"   orm:"colume(publisher_id)"`
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
}
type TypeItemAixinwuReq struct {
	MataData TypeMataData        `json:"mataData"`
	Token    string              `json:"token"`
	Item     TypeItemAixinwuInfo `json:"itemInfo"`
	Status   TypeStatus          `json:"status"`
}

// ln_jacount_info
type TypeLcnJacountInfo struct {
	Id          int    `json:"id"            orm:"pk;auto;colume(id)"`
	Customer_id int    `json:"customer_id"   orm:"colume(customer_id)"`
	Jaccount_id string `json:"jaccount_id"   orm:"colume(jaccount_id)"`
	Citizenid   string `json:"citizenid"     orm:"colume(citizenid)"`
	Realname    string `json:"realname"      orm:"colume(realname)"`
	Dept        string `json:"dept"          orm:"colume(dept)"`
	Tel         string `json:"tel"           orm:"colume(tel)"`
	Snum        string `json:"snum"          orm:"colume(snum)"`
	Is_student  string `json:"is_student"    orm:"colume(is_student)"`
}

func (u *TypeLcnJacountInfo) TableName() string {
	return "lcn_jaccount_info"
}

type TypeLcnDonateBatch struct {
	Id          int       `json:"id"             orm:"pk;auto;colume(id)"`
	User_id     int       `json:"user_id"        orm:"colume(user_id)"`
	Snum        string    `json:"snum"           orm:"colume(snum)"`
	Produced_at time.Time `json:"produced_at"    orm:"auto_now_add;type(datetime);colume(produced_at)"`
	Desc        string    `json:"desc"    orm:"  colume(desc)"`
	Donation_sn string    `json:"donation_sn"    orm:"colume(donation_sn)"`
	Barcode     string    `json:"barcode"        orm:"colume(barcode)"`
	Status      int       `json:"status"         orm:"colume(status)"`
}

func (u *TypeLcnDonateBatch) TableName() string {
	return "lcn_donation_batch"
}

/*
 * 	Aixinwu product databese
 */
//Table: lcn_product
type TypeAixinwuProduct struct {
	Id                     int       `json:"id"                      orm:"colume(id)"`
	Cat_id                 int       `json:"cat_id"                  orm:"colume(cat_id)"`
	Brand_id               int       `json:"brand_id"                orm:"colume(brand_id)"`
	Attr_set_id            int       `json:"attr_set_id"             orm:"colume(attr_set_id)"`
	Price                  int       `json:"price"                   orm:"colume(price)"`
	Market_price           int       `json:"market_price"            orm:"colume(market_price)"`
	Special_price          int       `json:"special_price"           orm:"colume(special_price)"`
	Name                   string    `json:"name"                    orm:"colume(name)"`
	Short_name             string    `json:"short_name"              orm:"colume(short_name)"`
	Url_alias              string    `json:"url_alias"               orm:"colume(url_alias)"`
	Short_desc             string    `json:"short_desc"              orm:"colume(short_desc)"`
	Desc                   string    `json:"desc"                    orm:"colume(desc)"`
	Weight                 int       `json:"weight"                  orm:"colume(weight)"`
	Stock                  int       `json:"stock"                   orm:"colume(stock)"`
	Limit                  int       `json:"limit"                   orm:"colume(limit)"`
	Is_on_sale             int       `json:"is_on_sale"              orm:"colume(is_on_sale)"`
	On_sale_at             time.Time `json:"on_sale_at"              orm:"colume(on_sale_at)"`
	Tag                    string    `json:"tag"                     orm:"colume(tag)"`
	Meta_title             string    `json:"meta_title"              orm:"colume(meta_title)"`
	Meta_keywords          string    `json:"meta_keywords"           orm:"colume(meta_keywords)"`
	Meta_desc              string    `json:"meta_desc"               orm:"colume(meta_desc)"`
	Is_new                 int       `json:"is_new"                  orm:"colume(is_new)"`
	Is_hot                 int       `json:"is_hot"                  orm:"colume(is_hot)"`
	Is_special_price       int       `json:"is_special_price"        orm:"colume(is_special_price)"`
	Special_price_start_at time.Time `json:"special_price_start_at"  orm:"colume(special_price_start_at)"`
	Special_price_end_at   time.Time `json:"special_price_end_at"    orm:"colume(special_price_end_at)"`
	Is_commend             int       `json:"is_commend"              orm:"colume(is_commend)"`
	Is_delete              int       `json:"is_delete"               orm:"colume(is_delete)"`
	Created_at             time.Time `json:"created_at"              orm:"colume(created_at)"`
	Updated_at             time.Time `json:"updated_at"              orm:"colume(updated_at)"`
	Code                   string    `json:"code"                    orm:"colume(code)"`
	Barcode                string    `json:"barcode"                 orm:"colume(barcode)"`
	Is_borrow              int       `json:"is_borrow"               orm:"colume(is_borrow)"`
	Is_cash                int       `json:"is_cash"                 orm:"colume(is_cash)"`
}

func (u *TypeAixinwuProduct) TableName() string {
	return "lcn_product"
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

const (
	StatusCodeOK             = iota
	StatusCodeErrorLoginInfo = iota
)

var ErrorDesp = map[int]string{
	StatusCodeOK:             "OK",
	StatusCodeErrorLoginInfo: "Wrong username or password",
}
