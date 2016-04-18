package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "aixinwu_test:$AIXINWU_test@tcp(localhost:3306)/aixinwu_test?charset=utf8")
	orm.RegisterModel(new(TypeUserInfo))
	orm.RegisterModel(new(TypeItemInfo))
	orm.RegisterModel(new(TypeItemComments))
	orm.RegisterModel(new(TypeChatInfo))
	createTable()
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
