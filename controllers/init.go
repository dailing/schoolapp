package controllers

import (
	"fmt"
	"log"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/garyburd/redigo/redis"
	"time"
)

var redisPool = redis.Pool{
	MaxIdle:     100,
	IdleTimeout: 240 * time.Second,
	Dial: func() (redis.Conn, error) {
		c, err := redis.Dial("tcp", "localhost:6379")
		if err != nil {
			return nil, err
		}
		return c, err
	},
}

func SysInit() {
	err := orm.RegisterDriver("mysql", orm.DRMySQL)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	//err = orm.RegisterDataBase("default", "mysql", "aixinwu_test:$AIXINWU_test@tcp(localhost:3306)/aixinwu_test?charset=utf8&loc=Asia%2FShanghai")
	err = orm.RegisterDataBase("default", "mysql", "sjtu:dywb!3396@tcp(localhost:3306)/sjtu_aixin?charset=utf8&loc=Asia%2FShanghai")
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	orm.RegisterModel(new(TypeUserInfo))
	orm.RegisterModel(new(TypeItemInfo))
	orm.RegisterModel(new(TypeItemComments))
	orm.RegisterModel(new(TypeChatInfo))
	orm.RegisterModel(new(TypeLcnDonateBatch))
	orm.RegisterModel(new(TypeAixinwuProduct))
	orm.RegisterModel(new(TypeAixinwuProductImage))
	orm.RegisterModel(new(TypeAixinwuItem))
	orm.RegisterModel(new(TypeAixinwuJaccountInfo))
	orm.RegisterModel(new(TypeAixinwuCustomCash))
	orm.RegisterModel(new(TypeAixinwuBook))
	orm.RegisterModel(new(TypeAixinwuOrder))
	orm.RegisterModel(new(TypeAixinwuOrderItem))
	orm.RegisterModel(new(TypeParameters))
	orm.RegisterModel(new(TypeAixinwuAddress))
	orm.RegisterModel(new(TypeAixinwuVolunteerAct))
	orm.RegisterModel(new(TypeServerParameters))
	orm.RegisterModel(new(TypeAixinwuVolunteer))
	createTable()
	makeFakeUser()

	ServerParameterSet("halfPrice", "yes")

	//	val := GetAixintuItems(10, 10)
	//	content, _ := json.Marshal(val)
	//	beego.Info(string(content))
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
		//Phone:    "12345678",
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
	SetMainPageItem("8,7,3,4,5,6")
}
