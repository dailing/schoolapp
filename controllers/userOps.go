package controllers

import (
	"github.com/astaxie/beego/orm"
)

func getUserInfo(name string) (info TypeUserInfo, err error) {
	o := orm.NewOrm()
	user := SQLuserinfo{
		Username: name,
	}
	err = o.Read(&user, "username")
	ErrReport(err)
	if err != nil {
		return
	}
	info = TypeUserInfo{
		Username: name,
		Password: user.Password,
		NickName: user.Nickname,
		Coins:    user.Coins,
	}
	return
}
