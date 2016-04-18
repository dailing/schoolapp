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
