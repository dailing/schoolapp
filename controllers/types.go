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
	Produced_at time.Time `json:"produced_at"    orm:"auto_now_add;colume(produced_at)"`
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
	id                     int
	cat_id                 int
	brand_id               int
	attr_set_id            int
	price                  int
	market_price           int
	special_price          int
	name                   string
	short_name             string
	url_alias              string
	short_desc             string
	desc                   string
	weight                 int
	stock                  int
	limit                  int
	is_on_sale             int
	on_sale_at             time.Time
	tag                    string
	meta_title             string
	meta_keywords          string
	meta_desc              string
	is_new                 int
	is_hot                 int
	is_special_price       int
	special_price_start_at time.Time
	special_price_end_at   time.Time
	is_commend             int
	is_delete              int
	created_at             time.Time
	updated_at             time.Time
	code                   string
	barcode                string
	is_borrow              int
	is_cash                int
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
