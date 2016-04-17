package controllers

import (
	"errors"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"regexp"
)

func GetUserInfo(name string) (info TypeUserInfo, err error) {
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
		ID:       user.Uid,
		Username: name,
		Password: user.Password,
		NickName: user.Nickname,
		Coins:    user.Coins,
	}
	return
}

func UpdateUserInfo(usrinfo TypeUserInfo) error {
	o := orm.NewOrm()
	o.Using("default")
	// check user name
	if succ, err := CheckUserNameLegal(usrinfo.Username); err != nil {
		return err
	} else if succ == false {
		return errors.New("Unknown Errors.")
	}
	s := SQLuserinfo{
		Username: usrinfo.Username,
	}
	err := o.Read(&s, "username")
	ErrReport(err)
	if err != nil {
		return err
	}
	beego.Trace("updating information for user id", s.Uid)
	s.Nickname = usrinfo.NickName
	s.Password = usrinfo.Password
	s.Coins = usrinfo.Coins
	num, err := o.Update(&s)
	ErrReport(err)
	beego.Info("affected rows when update:", num)
	return err
}

func AddUser(usrinfo TypeUserInfo) (int, error) {
	o := orm.NewOrm()
	o.Using("default")
	// check user name
	if succ, err := CheckUserNameLegal(usrinfo.Username); err != nil {
		return -1, err
	} else if succ == false {
		return -1, errors.New("Unknown Errors.")
	}
	// check if user already exist
	if succ, err := CheckUserNameExist(usrinfo.Username); succ {
		beego.Error(err)
		return -1, errors.New("User name already exists")
	}
	s := SQLuserinfo{
		Username: usrinfo.Username,
		Password: usrinfo.Password,
		Nickname: usrinfo.NickName,
		Coins:    0,
	}
	id, err := o.Insert(&s)
	if err != nil {
		return -1, err
	}
	beego.Info("Adding user name")
	return int(id), nil
}

func DelUser(name string) (bool, error) {
	uinfo, err := GetUserInfo(name)
	o := orm.NewOrm()
	info := SQLuserinfo{
		Uid:      uinfo.ID,
		Username: uinfo.Username,
	}
	beego.Trace("Del user id ", info.Uid, " for ", name)
	num, err := o.Delete(&info)
	ErrReport(err)
	beego.Trace("affected rows:", num)
	return true, err
}

func CheckUserNameLegal(name string) (bool, error) {
	beego.Info("Checking if User name legalty for", name)
	if name != "" {
		if succ, err := regexp.MatchString(`\w+`, name); succ == true {
			return true, nil
		} else {
			beego.Info("Check failed,", err)
			return false, err
		}
	}
	return false, errors.New("Name empty")
}

func CheckUserNameExist(name string) (bool, error) {
	beego.Info("Checking User name Existing")
	_, err := GetUserInfo(name)
	beego.Trace(err)
	if err != nil {
		return false, err
	}
	return true, nil
}
