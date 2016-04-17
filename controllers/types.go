package controllers

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
	UserID      int    `json:"userid"        orm:"colume(userid)"`
	Price       int    `json:"price"         orm:"colume(price)"`
	Images      string `json:"images"        orm:"colume(price)"`
	Status      int    `json:"status"        orm:"colume(status)"`
}

type TypeItemReqResp struct {
	MataData TypeMataData `json:"mataData"`
	Token    string       `json:"token"`
	ItemInfo TypeItemInfo `json:"itemInfo"`
	Status   TypeStatus   `json:"status"`
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
