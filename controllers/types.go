package controllers

type TypeMataData struct {
	TimeStamp int
	Device    string
}

type TypeUserInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
	NickName string `json:nickname`
	ID       string `json:"ID"`
	Phone    string `json:"Phone"`
	Email    string `json:"Email"`
}

type TypeStatus struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
}

const (
	StatusCodeOK             = iota
	StatusCodeErrorLoginInfo = iota
)

var ErrorDesp = map[int]string{
	StatusCodeOK:             "OK",
	StatusCodeErrorLoginInfo: "Wrong username or password",
}
