package controllers

import (
	"encoding/base64"
	"errors"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"regexp"
)

/*
 * 	User profile related operations
 */

func baseEncode(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

func baseDecode(str string) string {
	body, err := base64.StdEncoding.DecodeString(str)
	ErrReport(err)
	return string(body)
}

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
	//o.Using("default")
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
	//o.Using("default")
	// check user name
	//if succ, err := CheckUserNameLegal(usrInfo.Username); err != nil {
	//	return -1, err
	//} else if succ == false {
	//	return -1, errors.New("Unknown Errors.")
	//}
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
	//o.Using("default")
	itemInfo.Description = baseEncode(itemInfo.Description)
	id, err := o.Insert(&itemInfo)
	ErrReport(err)
	return int(id), err
}

func GetItemByID(id int) (TypeItemInfo, error) {
	o := orm.NewOrm()
	//o.Using("default")
	itemInfo := TypeItemInfo{
		ID: id,
	}
	err := o.Read(&itemInfo)
	ErrReport(err)
	itemInfo.Description = baseDecode(itemInfo.Description)
	return itemInfo, err
}

func GetItemsByUserID(id int) []TypeItemInfo {
	itemids := make([]TypeItemInfo, 0)
	o := orm.NewOrm()
	//o.Using("default")
	_, err := o.Raw("select * from type_item_info where owner_i_d = ?", id).QueryRows(&itemids)
	ErrReport(err)
	for i := 0; i < len(itemids); i++{
		itemids[i].Description = baseDecode(itemids[i].Description)
	}
	return itemids
}

func GetAllItem(startat int, length int) []TypeItemInfo {
	itemids := make([]TypeItemInfo, 0)
	o := orm.NewOrm()
	//o.Using("default")
	_, err := o.Raw("select * from type_item_info where i_d > ? and i_d <= ?", startat, startat+length).QueryRows(&itemids)
	ErrReport(err)
	for i := 0; i < len(itemids); i++{
		itemids[i].Description = baseDecode(itemids[i].Description)
	}
	return itemids
}

/*
 * 	Add and get comments
 */
func AddComments(comment TypeItemComments) (int, error) {
	o := orm.NewOrm()
	//o.Using("default")
	comment.Content = baseEncode(comment.Content)
	id, err := o.Insert(&comment)
	ErrReport(err)
	if err != nil {
		return -1, err
	}
	return int(id), nil
}

func GetComments(itemid int) []TypeItemComments {
	comments := make([]TypeItemComments, 0)
	o := orm.NewOrm()
	//o.Using("default")
	_, err := o.Raw("select * from type_item_comments where item_id = ?", itemid).QueryRows(&comments)
	ErrReport(err)
	for i := 0; i < len(comments); i++ {
		comments[i].Content = baseDecode(comments[i].Content)
	}
	return comments
}

func AddChat(chat TypeChatInfo) (int, error) {
	o := orm.NewOrm()
	//o.Using("default")
	chat.Content = baseEncode(chat.Content)
	id, err := o.Insert(&chat)
	ErrReport(err)
	if err != nil {
		return -1, err
	}
	return int(id), nil
}

/*
 *	Since a chat environment is determined by itemID-OwnerID and The Buyer.
 *	There are multiple ways to retrieve chat information.
 * 	Here Use OwnerID and Buyer ID if both are given.
 */
func GetChat(itemID, buyerID int) []TypeChatInfo {
	chats := make([]TypeChatInfo, 0)
	o := orm.NewOrm()
	//o.Using("default")
	_, err := o.Raw("select * from aixinwu_test.type_chat_info where item_id = ? and buyer_id = ?", itemID, buyerID).QueryRows(&chats)
	ErrReport(err)
	for i:= 0; i < len(chats); i++{
		chats[i].Content = baseDecode(chats[i].Content)
	}
	return chats
}
