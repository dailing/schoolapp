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

type TypeUserInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
	NickName string `json:"nickname"`
	ID       int    `json:"ID"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Coins    int    `josn:"coins"`
}

type TypeStatus struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
}

type TypeTokenInfo struct {
	UserName string
}

const (
	StatusCodeOK             = iota
	StatusCodeErrorLoginInfo = iota
)

var ErrorDesp = map[int]string{
	StatusCodeOK:             "OK",
	StatusCodeErrorLoginInfo: "Wrong username or password",
}
