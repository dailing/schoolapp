package controllers

import (
	"errors"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"regexp"
)

/*
 * 	User profile related operations
 */

func GetUserInfo(name string) (info TypeUserInfo, err error) {
	o := orm.NewOrm()
	info.Username = name
	err = o.Read(&info, "username")
	ErrReport(err)
	if err != nil {
		info.Username = ""
		return
	}
	return
}

func GetUserInfoByToken(token string) TypeUserInfo {
	tokenInfo := ParseToken(token)
	if tokenInfo.UserName == "" {
		return TypeUserInfo{}
	}
	usrInfo, err := GetUserInfo(tokenInfo.UserName)
	ErrReport(err)
	return usrInfo
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
	s := TypeUserInfo{
		Username: usrinfo.Username,
	}
	err := o.Read(&s, "username")
	ErrReport(err)
	if err != nil {
		return err
	}
	beego.Trace("updating information for user id", s.ID)
	s.NickName = usrinfo.NickName
	s.Password = usrinfo.Password
	s.Coins = usrinfo.Coins
	num, err := o.Update(&s)
	ErrReport(err)
	beego.Info("affected rows when update:", num)
	return err
}

func AddUser(usrInfo TypeUserInfo) (int, error) {
	o := orm.NewOrm()
	o.Using("default")
	// check user name
	if succ, err := CheckUserNameLegal(usrInfo.Username); err != nil {
		return -1, err
	} else if succ == false {
		return -1, errors.New("Unknown Errors.")
	}
	// check if user already exist
	if succ, err := CheckUserNameExist(usrInfo.Username); succ {
		beego.Error(err)
		return -1, errors.New("User name already exists")
	}
	s := TypeUserInfo{
		Username: usrInfo.Username,
		Password: usrInfo.Password,
		NickName: usrInfo.NickName,
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
	info := TypeUserInfo{
		ID:       uinfo.ID,
		Username: uinfo.Username,
	}
	beego.Trace("Del user id ", info.ID, " for ", name)
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

/*
 * 	Item Related Operations
 */
func AddItem(itemInfo TypeItemInfo) (int, error) {
	o := orm.NewOrm()
	o.Using("default")
	id, err := o.Insert(&itemInfo)
	ErrReport(err)
	return int(id), err
}

func GetItemByID(id int) (TypeItemInfo, error) {
	o := orm.NewOrm()
	o.Using("default")
	itemInfo := TypeItemInfo{
		ID: id,
	}
	err := o.Read(&itemInfo)
	ErrReport(err)
	return itemInfo, err
}

func GetItemIDsByUserID(id int) []int {
	return nil
}
