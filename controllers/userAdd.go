package controllers

import (
	"database/sql"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type UserAddController struct {
	beego.Controller
}

type SQLuserinfo struct {
	Uid      int    `orm:"pk;auto"`
	Username string `orm:"unique;column(username)"`
	Password string `orm:"column(password)"`
	Nickname string `orm:"column(nickname)"`
	Coins    int    `orm:"column(coins)"`
}

/*
	This function add a user to mysql and
	return the ID generated by mysql.
*/
func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "aixinwu:aixinwu@tcp(localhost:3306)/appdev?charset=utf8")
	orm.RegisterModel(new(SQLuserinfo))
}

func createTable() {
	name := "default"
	force := false
	verbose := true
	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		beego.Error(err)
	}
}

func addUser(usrinfo TypeUserInfo) (string, error) {
	o := orm.NewOrm()
	o.Using("default")
	createTable()
	fmt.Println(o)
	s := SQLuserinfo{
		Username: usrinfo.Username,
		Password: usrinfo.Password,
		Nickname: usrinfo.NickName,
		Coins:    0,
	}
	retval, err := o.Insert(&s)
	if err != nil {
		return "", err
	}
	return fmt.Sprint(retval), nil
}

func _addUser(usrinfo TypeUserInfo) (string, error) {
	db, err := sql.Open("mysql", "aixinwu:aixinwu@tcp(localhost:3306)/appdev")
	if err != nil {
		fmt.Print(err)
		return "", err
	}
	stmt, err := db.Prepare("INSERT userinfo SET username=?,nickname=?,password=?,coins=?")
	if err != nil {
		fmt.Print(err)
		return "", err
	}

	res, err := stmt.Exec(usrinfo.Username, usrinfo.NickName, usrinfo.Password, 0)
	if err != nil {
		fmt.Print(err)
		return "", err
	}

	id, err := res.LastInsertId()
	if err != nil {
		fmt.Print(err)
		return "", err
	}
	stmt.Close()
	db.Close()
	return fmt.Sprint(id), nil
}

func (c *UserAddController) Post() {
	beego.Debug("add user")

}
