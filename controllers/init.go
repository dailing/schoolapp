package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"log"
)

func SysInit() {
	err := orm.RegisterDriver("mysql", orm.DRMySQL)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	err = orm.RegisterDataBase("default", "mysql", "aixinwu_test:$AIXINWU_test@tcp(localhost:3306)/aixinwu_test?charset=utf8")
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	orm.RegisterModel(new(TypeUserInfo))
	orm.RegisterModel(new(TypeItemInfo))
	orm.RegisterModel(new(TypeItemComments))
	orm.RegisterModel(new(TypeChatInfo))
	//orm.RegisterModel(new(TypeLcnJacountInfo))
	orm.RegisterModel(new(TypeLcnDonateBatch))
	orm.RegisterModel(new(TypeAixinwuProduct))
	orm.RegisterModel(new(TypeAixinwuItem))
	orm.RegisterModel(new(TypeAixinwuJaccountInfo))
	orm.RegisterModel(new(TypeAixinwuCustomCash))
	orm.RegisterModel(new(TypeParameters))
	createTable()
	makeFakeUser()
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

func makeFakeUser() {
	beego.Trace("Adding fake users ......")
	user := TypeUserInfo{
		Username: "a@a.com",
		Password: "1234",
		NickName: "nick",
		Phone:    "12345678",
		Email:    "a@a.com",
		JAccount: "liangyuding",
	}
	AddUser(user)
	user.Username = "b@b.com"
	user.JAccount = "SJTULR"
	user.NickName = "bbb"
	AddUser(user)
	user.Username = "c@c.com"
	user.JAccount = "----lihengfu"
	user.NickName = "ccc"
	AddUser(user)
	beego.Trace("Finished fake users ......")
	SetMainPageItem("1,2,3,4,5,6")
}
