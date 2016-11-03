package controllers

import (
	"encoding/json"
	"time"
)

type AWXtime struct {
	time.Time
}

const ctLayout = "2006-01-02 15:04:05"

//func (ct *AWXtime) UnmarshalJSON(b []byte) (err error) {
//	if b[0] == '"' && b[len(b)-1] == '"' {
//		b = b[1 : len(b)-1]
//	}
//	ct.Time, err = time.Parse(ctLayout, string(b))
//	return
//}

func (ct AWXtime) MarshalJSON() ([]byte, error) {
	return []byte(`"` + ct.Time.Local().Format(ctLayout) + `"`), nil
}

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
	Image    string       `json:"image"`
}

// image related request parameters
type TypeImgResp struct {
	MataData TypeMataData `json:"mataData"`
	ImageID  string       `json:"imageID"`
	Token    string       `json:"token"`
	Status   TypeStatus   `json:"status"`
}

type TypeUserInfo struct {
	ID               int    `json:"ID"       orm:"pk;auto;column(Uid)"`
	Username         string `json:"username" orm:"type(text);column(username)"`
	Image            string `json:"image"    orm:"type(text);unique;column(image)"`
	Password         string `json:"password" orm:"column(password)"`
	VerificationCode string `json:"verification_code"`
	NickName         string `json:"nickname" orm:"type(text);column(nickname)"`
	//Phone            string  `json:"phone"    orm:"column(phone)"`
	Email    string  `json:"email"    orm:"column(email)"`
	Coins    float64 `json:"coins"    orm:"column(coins)"`
	JAccount string  `json:"jaccount" orm:"column(jaccount)"`
}

func (t *TypeUserInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Username string  `json:"username" orm:"type(text);column(username)"`
		Image    string  `json:"image"    orm:"type(text);unique;column(image)"`
		NickName string  `json:"nickname" orm:"type(text);column(nickname)"`
		Email    string  `json:"email"    orm:"column(email)"`
		Coins    float64 `json:"coins"    orm:"column(coins)"`
		JAccount string  `json:"jaccount" orm:"column(jaccount)"`
	}{
		Username: t.Username,
		Image:    t.Image,
		NickName: t.NickName,
		Email:    t.Email,
		Coins:    t.Coins,
		JAccount: t.JAccount,
	})
}

func (u *TypeUserInfo) TableName() string {
	return "app_user_info"
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
	BoughtAt                time.Time `json:"boughtAt"                 orm:"(boughtAt);type(datetime);null"`
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
	UserSuggestedPrice      int       `json:"user_suggested_price"     orm:"(user_suggested_price);null"`
	Images                  string    `json:"images"                   orm:"(images);type(text);null"`
	Status                  int       `json:"status"                   orm:"(status);null"`
}

func (u *TypeItemInfo) TableName() string {
	return "app_item_info"
}

func (t *TypeItemInfo) MarshalJSON() ([]byte, error) {
	type Alias TypeItemInfo
	return json.Marshal(&struct {
		*Alias
		BoughtAt    AWXtime `json:"boughtAt"`
		PublishedAt AWXtime `json:"publishedAt"              orm:"(publishedAt);null"`
	}{
		BoughtAt:    AWXtime{t.BoughtAt},
		PublishedAt: AWXtime{t.PublishedAt},
		Alias:       (*Alias)(t),
	})
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

func (u *TypeItemComments) TableName() string {
	return "app_item_comments"
}

func (t *TypeItemComments) MarshalJSON() ([]byte, error) {
	type Alias TypeItemComments
	return json.Marshal(&struct {
		*Alias
		Created AWXtime `json:"created"       orm:"auto_now_add;type(datetime)"`
	}{
		Created: AWXtime{t.Created},
		Alias:   (*Alias)(t),
	})
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

func (u *TypeChatInfo) TableName() string {
	return "app_chat_info"
}

func (t *TypeChatInfo) MarshalJSON() ([]byte, error) {
	type Alias TypeChatInfo
	return json.Marshal(&struct {
		*Alias
		Created AWXtime `json:"created"       orm:"auto_now_add;type(datetime)"`
	}{
		Created: AWXtime{t.Created},
		Alias:   (*Alias)(t),
	})
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

func (t *TypeLcnDonateBatch) MarshalJSON() ([]byte, error) {
	type Alias TypeLcnDonateBatch
	return json.Marshal(&struct {
		*Alias
		Produced_at AWXtime `json:"produced_at"    orm:"auto_now_add;type(datetime);(produced_at)"`
	}{
		Produced_at: AWXtime{t.Produced_at},
		Alias:       (*Alias)(t),
	})
}

func (u *TypeLcnDonateBatch) TableName() string {
	return "lcn_donation_batch"
}

/*
 * 	Aixinwu product databese
 */

type InterfaceAixinwuProduct interface {
	SetID(id int)
	GetID() int
	GetStock() int
	SetStock(int)
	GetName() string
	GetPrice() float64
	GetCategory() int
	GetWeight() int
}

//Table: lcn_product
type TypeAixinwuProduct struct {
	Id                     int       `json:"id"                      orm:"pk;auto;column(id)"`
	Cat_id                 int       `json:"cat_id"                  orm:"column(cat_id)"`
	Brand_id               int       `json:"brand_id"                orm:"column(brand_id)"`
	Attr_set_id            int       `json:"attr_set_id"             orm:"column(attr_set_id)"`
	Price                  float64   `json:"price"                   orm:"column(price)"`
	Market_price           float64   `json:"market_price"            orm:"column(market_price)"`
	Special_price          float64   `json:"special_price"           orm:"column(special_price)"`
	Name                   string    `json:"name"                    orm:"column(name)"`
	Short_name             string    `json:"short_name"              orm:"column(short_name)"`
	Url_alias              string    `json:"url_alias"               orm:"column(url_alias)"`
	Short_desc             string    `json:"short_desc"              orm:"column(short_desc)"`
	Desc                   string    `json:"desc"                    orm:"column(desc)"`
	Weight                 int       `json:"weight"                  orm:"column(weight)"`
	Stock                  int       `json:"stock"                   orm:"column(stock)"`
	Limit                  int       `json:"limit"                   orm:"column(limit)"`
	Is_on_sale             int       `json:"is_on_sale"              orm:"column(is_on_sale)"`
	On_sale_at             time.Time `json:"on_sale_at"              orm:"column(on_sale_at)"`
	Tag                    string    `json:"tag"                     orm:"column(tag)"`
	Meta_title             string    `json:"meta_title"              orm:"column(meta_title)"`
	Meta_keywords          string    `json:"meta_keywords"           orm:"column(meta_keywords)"`
	Meta_desc              string    `json:"meta_desc"               orm:"column(meta_desc)"`
	Is_new                 int       `json:"is_new"                  orm:"column(is_new)"`
	Is_hot                 int       `json:"is_hot"                  orm:"column(is_hot)"`
	Is_special_price       int       `json:"is_special_price"        orm:"column(is_special_price)"`
	Special_price_start_at time.Time `json:"special_price_start_at"  orm:"column(special_price_start_at)"`
	Special_price_end_at   time.Time `json:"special_price_end_at"    orm:"column(special_price_end_at)"`
	Is_commend             int       `json:"is_commend"              orm:"column(is_commend)"`
	Is_delete              int       `json:"is_delete"               orm:"column(is_delete)"`
	Created_at             time.Time `json:"created_at"              orm:"column(created_at)"`
	Updated_at             time.Time `json:"updated_at"              orm:"column(updated_at)"`
	Code                   string    `json:"code"                    orm:"column(code)"`
	Barcode                string    `json:"barcode"                 orm:"column(barcode)"`
	Is_borrow              int       `json:"is_borrow"               orm:"column(is_borrow)"`
	Is_cash                int       `json:"is_cash"                 orm:"column(is_cash)"`
	Image                  string    `json:"image"                   orm:"-"` // ignore this field in database
	DespUrl                string    `json:"desp_url"                orm:"-"` // ignore field
	OriginalPrice          float64   `json:"original_price"          orm:"-"`
}

func (t *TypeAixinwuProduct) MarshalJSON() ([]byte, error) {
	type Alias TypeAixinwuProduct
	return json.Marshal(&struct {
		*Alias
		On_sale_at             AWXtime `json:"on_sale_at"              orm:"column(on_sale_at)"`
		Special_price_start_at AWXtime `json:"special_price_start_at"  orm:"column(special_price_start_at)"`
		Special_price_end_at   AWXtime `json:"special_price_end_at"    orm:"column(special_price_end_at)"`
		Created_at             AWXtime `json:"created_at"              orm:"column(created_at)"`
		Updated_at             AWXtime `json:"updated_at"              orm:"column(updated_at)"`
	}{
		On_sale_at:             AWXtime{t.On_sale_at},
		Special_price_start_at: AWXtime{t.Special_price_start_at},
		Special_price_end_at:   AWXtime{t.Special_price_end_at},
		Created_at:             AWXtime{t.Created_at},
		Updated_at:             AWXtime{t.Updated_at},
		Alias:                  (*Alias)(t),
	})
}

func (u *TypeAixinwuProduct) TableName() string {
	return "lcn_product"
}
func (u *TypeAixinwuProduct) SetID(id int) {
	u.Id = id
}
func (u *TypeAixinwuProduct) GetID() int {
	return u.Id
}
func (u *TypeAixinwuProduct) GetStock() int {
	return u.Stock
}
func (u *TypeAixinwuProduct) SetStock(stock int) {
	u.Stock = stock
}
func (u *TypeAixinwuProduct) GetName() string {
	return u.Name
}
func (u *TypeAixinwuProduct) GetPrice() float64 {
	return float64(u.Price)
}
func (u *TypeAixinwuProduct) GetCategory() int {
	return 0
}
func (u *TypeAixinwuProduct) GetWeight() int {
	return u.Weight
}

type TypeAixinwuProductImage struct {
	Id         int       `json:"id"           orm:"auto;pk;column(id)"`
	Product_id int       `json:"product_id"   orm:"column(product_id)"`
	Image_name string    `json:"image_name"   orm:"column(image_name)"`
	File       string    `json:"file"         orm:"column(file)"`
	File_ext   string    `json:"file_ext"     orm:"column(file_ext)"`
	File_mime  string    `json:"file_mime"    orm:"column(file_mime)"`
	Width      int       `json:"width"        orm:"column(width)"`
	Height     int       `json:"height"       orm:"column(height)"`
	File_size  int       `json:"file_size"    orm:"column(file_size)"`
	Is_base    int       `json:"is_base"      orm:"column(is_base)"`
	Sort_order int       `json:"sort_order"   orm:"column(sort_order)"`
	Created_at time.Time `json:"created_at"   orm:"column(created_at)"`
	Updated_at time.Time `json:"updated_at"   orm:"column(updated_at)"`
}

func (t *TypeAixinwuProductImage) MarshalJSON() ([]byte, error) {
	type Alias TypeAixinwuProductImage
	return json.Marshal(&struct {
		*Alias
		Created_at AWXtime `json:"created_at"   orm:"column(created_at)"`
		Updated_at AWXtime `json:"updated_at"   orm:"column(updated_at)"`
	}{
		Created_at: AWXtime{t.Created_at},
		Updated_at: AWXtime{t.Updated_at},
		Alias:      (*Alias)(t),
	})
}

func (u *TypeAixinwuProductImage) TableName() string {
	return "lcn_product_image"
}

type TypeAixinwuBook struct {
	ISBN       string  `json:"ISBN"        orm:"pk;column(ISBN)"`
	Image      string  `json:"image"       orm:"column(image)"`
	Title      string  `json:"title"       orm:"column(title)"`
	Author     string  `json:"author"      orm:"column(author)"`
	Press      string  `json:"press"       orm:"column(press)"`
	Pubyear    string  `json:"pubyear"     orm:"column(pubyear)"`
	Pagecnt    int     `json:"pagecnt"     orm:"column(pagecnt)"`
	Price      float64 `json:"price"       orm:"column(price)"`
	Sold       int     `json:"sold"        orm:"column(sold)"`
	Discard    int     `json:"discard"     orm:"column(discard)"`
	Stock      int     `json:"stock"       orm:"column(stock)"`
	Sale_price float64 `json:"sale_price"  orm:"column(sale_price)"`
}

func (u *TypeAixinwuBook) TableName() string {
	return "lcn_booktrade_books"
}

func (u *TypeAixinwuBook) SetID(id int) {
	ErrReport("Not Implemented")
}
func (u *TypeAixinwuBook) GetID() int {
	ErrReport("Not Implemented")
	return 0
}
func (u *TypeAixinwuBook) GetStock() int {
	return u.Stock
}
func (u *TypeAixinwuBook) SetStock(stock int) {
	u.Stock = stock
}
func (u *TypeAixinwuBook) GetName() string {
	return u.Title
}
func (u *TypeAixinwuBook) GetPrice() float64 {
	return u.Price
}
func (u *TypeAixinwuBook) GetCategory() int {
	return 2
}
func (u *TypeAixinwuBook) GetWeight() int {
	return 0
}

type TypeAixinwuOrder struct {
	Id                  int                    `json:"id"                    orm:"auto;pk;column(id)"`
	Order_sn            string                 `json:"order_sn"              orm:"column(order_sn)"`
	Customer_id         int                    `json:"customer_id"           orm:"column(customer_id)"`
	Payment_id          int                    `json:"payment_id"            orm:"column(payment_id)"`
	Shipping_id         int                    `json:"shipping_id"           orm:"column(shipping_id)"`
	Total_product_price float64                `json:"total_product_price"   orm:"column(total_product_price)"`
	Total_weight        int                    `json:"total_weight"          orm:"column(total_weight)"`
	Auto_freight_fee    float64                `json:"auto_freight_fee"      orm:"column(auto_freight_fee)"`
	Actual_freight_fee  float64                `json:"actual_freight_fee"    orm:"column(actual_freight_fee)"`
	Payment_fee         float64                `json:"payment_fee"           orm:"column(payment_fee)"`
	Total_cost          float64                `json:"total_cost"            orm:"column(total_cost)"`
	Total_price         float64                `json:"total_price"           orm:"column(total_price)"`
	Need_pay            float64                `json:"need_pay"              orm:"column(need_pay)"`
	Already_pay         float64                `json:"already_pay"           orm:"column(already_pay)"`
	Is_need_invoice     int                    `json:"is_need_invoice"       orm:"column(is_need_invoice)"`
	Customer_remark     string                 `json:"customer_remark"       orm:"column(customer_remark)"`
	Status              int                    `json:"status"                orm:"column(status)"`
	Is_delete           int                    `json:"is_delete"             orm:"column(is_delete)"`
	Barcode             string                 `json:"barcode"               orm:"column(barcode)"`
	Consignee_id        int                    `json:"consignee_id"          orm:"column(consignee_id)"`
	Place_at            time.Time              `json:"place_at"              orm:"column(place_at)"`
	Update_at           time.Time              `json:"update_at"             orm:"column(update_at)"`
	Items               []TypeAixinwuOrderItem `json:"items"                 orm:"-"`
}

func (t *TypeAixinwuOrder) MarshalJSON() ([]byte, error) {
	type Alias TypeAixinwuOrder
	return json.Marshal(&struct {
		*Alias
		Place_at  AWXtime `json:"place_at"              orm:"column(place_at)"`
		Update_at AWXtime `json:"update_at"             orm:"column(update_at)"`
	}{
		Place_at:  AWXtime{t.Place_at},
		Update_at: AWXtime{t.Update_at},
		Alias:     (*Alias)(t),
	})
}

func (u *TypeAixinwuOrder) TableName() string {
	return "lcn_order"
}
func (u *TypeAixinwuOrder) GetSN() string {
	return u.Order_sn
}
func (u *TypeAixinwuOrder) SetSN(sn string) {
	u.Order_sn = sn
}

type TypeAixinwuOrderItem struct {
	Id           int     `json:"id"               orm:"auto;pk;column(id)"`
	Order_id     int     `json:"order_id"         orm:"column(order_id)"`
	Product_id   string  `json:"product_id"       orm:"column(product_id)"`
	Product_name string  `json:"product_name"     orm:"column(product_name)"`
	Quantity     int     `json:"quantity"         orm:"column(quantity)"`
	Price        float64 `json:"price"            orm:"column(price)"`
	Weight       int     `json:"weight"           orm:"column(weight)"`
	Category     int     `json:"category"         orm:"column(category)"`
	Image        string  `json:"image"            orm:"-"`
}

func (u *TypeAixinwuOrderItem) TableName() string {
	return "lcn_order_item"
}

type TypeAixinwuItem struct {
	Id             int       `json:"id"               orm:"pk;auto;column(id)"`
	Barcode        string    `json:"barcode"          orm:"column(barcode)"`
	Name           string    `json:"name"             orm:"column(name)"`
	Valuation      int       `json:"valuation"        orm:"column(valuation)"`
	Status         int       `json:"status"           orm:"column(status)"`
	Quantity       int       `json:"quantity"         orm:"column(quantity)"`
	Quantity_saled int       `json:"quantity_saled"   orm:"column(quantity_saled)"`
	Category       int       `json:"category"         orm:"column(category)"`
	Donation_id    int       `json:"donation_id"      orm:"column(donation_id)"`
	Product_id     int       `json:"product_id"       orm:"column(product_id)"`
	Order_id       string    `json:"order_id"         orm:"column(order_id)"`
	Description    string    `json:"description"      orm:"column(description)"`
	Create_time    time.Time `json:"create_time"      orm:"column(create_time)"`
	Validity       time.Time `json:"validity"         orm:"column(validity)"`
	Is_delete      int       `json:"is_delete"        orm:"column(is_delete)"`
	Image_name     string    `json:"image_name"       orm:"column(image_name)"`
}

func (t *TypeAixinwuItem) MarshalJSON() ([]byte, error) {
	type Alias TypeAixinwuItem
	return json.Marshal(&struct {
		*Alias
		Create_time AWXtime `json:"create_time"      orm:"column(create_time)"`
		Validity    AWXtime `json:"validity"         orm:"column(validity)"`
	}{
		Create_time: AWXtime{t.Create_time},
		Validity:    AWXtime{t.Validity},
		Alias:       (*Alias)(t),
	})
}

func (u *TypeAixinwuItem) TableName() string {
	return "lcn_item"
}

var AixinwuItemType = struct {
	AixinwuItemType_exchange string
	AixinwuItemType_rent     string
	AixinwuItemType_cash     string
}{
	AixinwuItemType_exchange: "置换",
	AixinwuItemType_rent:     "租赁",
	AixinwuItemType_cash:     "现金",
}

type TypeAixinwuItemReqResp struct {
	MataData TypeMataData `json:"mataData"`
	Token    string       `json:"token"`
	StartAt  int          `json:"startAt"`
	Length   int          `json:"length"`
	Category int          `json:"category"`
	/*
	 * the value of Type field must be one value in
	 * AixinwuItemType struct
	 */
	Type         string               `json:"type"`
	Status       TypeStatus           `json:"status"`
	AixinwuItems []TypeAixinwuProduct `json:"items"`
}

type TypeAixinwuJaccountInfo struct {
	Id          int    `json:"id"            orm:"pk;auto;column(id)"`
	Customer_id int    `json:"customer_id"   orm:"column(customer_id)"`
	Jaccount_id string `json:"jaccount_id"   orm:"column(jaccount_id)"`
	Citizenid   string `json:"citizenid"     orm:"column(citizenid)"`
	Realname    string `json:"realname"      orm:"column(realname)"`
	Dept        string `json:"dept"          orm:"column(dept)"`
	Tel         string `json:"tel"           orm:"column(tel)"`
	Snum        string `json:"snum"          orm:"column(snum)"`
	Is_student  string `json:"is_student"    orm:"column(is_student)"`
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
	ID           int    `json:"homePageItem"        orm:"pk;column(id)"`
	HomePageItem string `json:"homePageItem"        orm:"column(homePageItem)"`
}

type TypeParametersRwqResp struct {
	MataData   TypeMataData   `json:"mataData"`
	Parameters TypeParameters `json:"parameter"`
	Status     TypeStatus     `json:"status"`
}

type TypeHomePage struct {
	MataData   TypeMataData   `json:"mataData"`
	Parameters TypeParameters `json:"items"`
	Status     TypeStatus     `json:"status"`
}

type TypeOrderProduct struct {
	ProductID int  `json:"product_id"`
	IsBook    bool `json:"is_book"`
	Quantity  int  `json:"quantity"`
}

type TypeOrderProductReqResp struct {
	OrderID  int          `json:"order_id"` // this is output value
	MataData TypeMataData `json:"mataData"`
	Status   TypeStatus   `json:"status"`
	// the following are input values
	OrderInfo   []TypeOrderProduct `json:"order_info"`
	ConsigneeID int                `json:"consignee_id"`
	Token       string             `json:"token"`
}

type TestMessageVerification struct {
	MataData TypeMataData `json:"mataData"`
	Phone    string       `json:"phone"`
	Code     string       `json:"code"`
	Status   TypeStatus   `json:"status"`
}

type TypeAixinwuAddress struct {
	Id           int       `json:"id"             orm:"column(id);auto;pk"`
	Customer_id  string    `json:"customer_id"    orm:"column(customer_id)"`
	Address_name string    `json:"address_name"   orm:"column(address_name)"`
	Consignee    string    `json:"consignee"      orm:"column(consignee)"`
	Phone        string    `json:"phone"          orm:"column(phone)"`
	Mobile       string    `json:"mobile"         orm:"column(mobile)"`
	Fax          string    `json:"fax"            orm:"column(fax)"`
	Postcode     string    `json:"postcode"       orm:"column(postcode)"`
	Address      string    `json:"address"        orm:"column(address)"`
	Province_id  int       `json:"province_id"    orm:"column(province_id)"`
	City_id      int       `json:"city_id"        orm:"column(city_id)"`
	District_id  int       `json:"district_id"    orm:"column(district_id)"`
	Is_default   int       `json:"is_default"     orm:"column(is_default)"`
	Invoice_head string    `json:"invoice_head"   orm:"column(invoice_head)"`
	Remark       string    `json:"remark"         orm:"column(remark)"`
	Created_at   time.Time `json:"created_at"     orm:"column(created_at)"`
	Updated_at   time.Time `json:"updated_at"     orm:"column(updated_at)"`
	Snum         string    `json:"snum"           orm:"column(snum)"`
	Email        string    `json:"email"          orm:"column(email)"`
	Is_deleted   int       `json:"is_deleted"     orm:"column(is_deleted)"`
}

func (u *TypeAixinwuAddress) TableName() string {
	return "lcn_customer_address"
}

type TypeAixinwuAddressResp struct {
	MataData TypeMataData         `json:"mata_data"`
	Address  []TypeAixinwuAddress `json:"address"`
	Status   TypeStatus           `json:"status"`
}

type TypeAixinwuOrderGetReq struct {
	Token  string `json:"token"`
	Offset int    `json:"offset"`
	Length int    `json:"length"`
}

type TypeAixinwuOrderGetResp struct {
	Status TypeStatus         `json:"status"`
	Orders []TypeAixinwuOrder `json:"orders"`
}

type TypeAixinwuOrderItemGetReq struct {
	Token   string `json:"token"`
	OrderID int    `json:"order_id"`
}
type TypeAixinwuOrderItemResp struct {
	Items  []TypeAixinwuOrderItem `json:"items"`
	Status TypeStatus             `json:"status"`
}

type TypeAixinwuVolunteerAct struct {
	Id         int       `json:"id"         orm:"column(id)"`
	Name       string    `json:"name"       orm:"column(name)"`
	Poster     string    `json:"poster"     orm:"column(poster)"`
	About      string    `json:"about"      orm:"column(about)"`
	Content    string    `json:"content"    orm:"column(content)"`
	Work_date  time.Time `json:"work_date"  orm:"column(work_date)"`
	Workload   int       `json:"workload"   orm:"column(workload)"`
	Site       string    `json:"site"       orm:"column(site)"`
	Status     int       `json:"status"     orm:"column(status)"`
	Pay_cash   int       `json:"pay_cash"   orm:"column(pay_cash)"`
	Num_needed int       `json:"num_needed" orm:"column(num_needed)"`
	Num_signed int       `json:"num_signed" orm:"column(num_signed)"`
	Created_at time.Time `json:"created_at" orm:"column(created_at)"`
	Updated_at time.Time `json:"updated_at" orm:"column(updated_at)"`
	Image      string    `json:"image"      orm:"-"`
	Joined     int       `json:"joined"     orm:"-"`
}

func (u *TypeAixinwuVolunteerAct) TableName() string {
	return "lcn_volunteer_act"
}

type TypeServerParameters struct {
	Id    int    `orm:"column(id);pk;auto"`
	Key   string `orm:"column(key);unique"`
	Value string `orm:"column(value)"`
}

type TypeAixinwuVolunteer struct {
	Id         int       `json:"id"          orm:"column(id)"`
	Uid        int       `json:"uid"         orm:"column(uid)"`
	Project_id int       `json:"project_id"  orm:"column(project_id)"`
	Project    string    `json:"project"     orm:"column(project)"`
	Work_date  time.Time `json:"work_date"   orm:"column(work_date)"`
	Update_at  time.Time `json:"update_at"   orm:"column(update_at)"`
	Workload   float64   `json:"workload"    orm:"column(workload)"`
	Content    string    `json:"content"     orm:"column(content)"`
	Tel        string    `json:"tel"         orm:"column(tel)"`
	Pay_cash   int       `json:"pay_cash"    orm:"column(pay_cash)"`
}

func (u *TypeAixinwuVolunteer) TableName() string {
	return "lcn_volunteer"
}

type TypeAixinwuVolunteerJoinReq struct {
	Token      string    `json:"token"`
	Project_id int       `json:"project_id"`
	Work_date  time.Time `json:"work_date"`
	Tel        string    `json:"tel"`
}

type TypeSetAddressReq struct {
	Token     string `json:"token"`
	Consignee string `json:"consignee"      orm:"column(consignee)"`
	Mobile    string `json:"mobile"         orm:"column(mobile)"`
	Snum      string `json:"snum"           orm:"column(snum)"`
}

const (
	StatusCodeOK             = iota
	StatusCodeErrorLoginInfo = iota
	StatusCodeUndefinedError = iota
	StatusCodeNotImplemented = iota
	StatusCodeNotEnoughMoney = iota
	StatusCodeDatabaseErr    = iota
	StatusNoJaccountInfo     = iota
)

var ErrorDesp = map[int]string{
	StatusCodeOK:             "OK",
	StatusCodeErrorLoginInfo: "Wrong username or password",
	StatusCodeUndefinedError: "Not specified",
	StatusCodeNotImplemented: "Not implemented",
	StatusCodeNotEnoughMoney: "Not enourgh money",
	StatusCodeDatabaseErr:    "Database error",
	StatusNoJaccountInfo:     "No jaccount information found",
}
