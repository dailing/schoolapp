package controllers

type TypeMataData struct {
	TimeStamp int
	Device    string
}

type TypeRegularResp struct {
	MataData TypeMataData `json:"mataData"`
	Status   TypeStatus   `json:"status"`
}

type TypeUserInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
	NickName string `json:"nickname"`
	ID       string `json:"ID"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Coins    string `josn:"coins"`
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
