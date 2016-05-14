package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "aixinwu_test:$AIXINWU_test@tcp(localhost:3306)/aixinwu_test?charset=UTF8")
	orm.RegisterModel(new(TypeUserInfo))
	orm.RegisterModel(new(TypeItemInfo))
	orm.RegisterModel(new(TypeItemComments))
	orm.RegisterModel(new(TypeChatInfo))
	orm.RegisterModel(new(TypeLcnJacountInfo))
	orm.RegisterModel(new(TypeLcnDonateBatch))
	orm.RegisterModel(new(TypeAixinwuProduct))
	orm.RegisterModel(new(TypeAixinwuItem))
	createTable()
}

func createTable() {
	fmt.Println("here")
	name := "default"
	force := false
	verbose := true
	err := orm.RunSyncdb(name, force, verbose)
	fmt.Println("here2")
	if err != nil {
		beego.Error(err)
	}
}
