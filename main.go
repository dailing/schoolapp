package main

import (
	"io/ioutil"
	"os"
	"os/exec"
	"time"

	"git.oschina.net/dddailing/schoolapp/controllers"
	"git.oschina.net/dddailing/schoolapp/natsChatServer"
	_ "git.oschina.net/dddailing/schoolapp/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type TypeItemInfo struct {
	ID                      int       `json:"ID"                       orm:"pk;auto;(id)"`
	Caption                 string    `json:"caption"                  orm:"(caption);type(text);null"`
	BoughtAt                time.Time `json:"boughtAt"                 orm:"(boughtAt);type(datetime)null"`
	ItemCondition           int       `json:"itemCondition"            orm:"(itemCondition);null"`
	EstimatedPriceByUser    int       `json:"estimatedPriceByUser"     orm:"(estimatedPriceByUser);null"`
	EstimatedPriceByAiXinWu int       `json:"estimatedPriceByAiXinWu"  orm:"(estimatedPriceByAiXinWu);null"`
	Location                string    `json:"location"                 orm:"(location);type(text);null"`
	Category                int       `json:"category"                 orm:"(category);null"`
	PublishedAt             time.Time `json:"publishedAt"              orm:"(publishedAt);null"`
	UnShelveReason          string    `json:"unShelveReason"           orm:"(unShelveReason);type(text);null"`
	Description             string    `json:"description"              orm:"(description);type(text);null"`
	OwnerID                 int       `json:"ownerID"                  orm:"(ownerID);null"`
	Price                   int       `json:"price"                    orm:"(price);null"`
	UserSuggestedPrice      int       `json:"user_suggested_price"    orm:"(user_suggested_price);null"`
	Images                  string    `json:"images"                   orm:"(images);type(text);null"`
	Status                  int       `json:"status"                   orm:"(status);null"`
}

func main() {

	orm.DebugLog = orm.NewLog(os.Stdout)
	beego.SetLogger("file", `{"filename":"test.log"}`)
	beego.Info("This is an info log")
	controllers.SysInit()
	// start gnatsd
	if _, err := os.Stat("gnatsd"); os.IsNotExist(err) {
		beego.Warning("No gnatsd found")
	} else {
		cmd := exec.Command("./gnatsd", "-l", "gnatsd.log", "-V")
		if cmd == nil {
			beego.Error()
		}
		cmd.Stderr = ioutil.Discard
		cmd.Stdout = ioutil.Discard
		controllers.ErrReport(cmd.Start())
		defer cmd.Process.Kill()
	}
	// start chat server
	time.Sleep(time.Second)
	server := natsChatServer.NewServer()
	server.Start()
	beego.Run()
}
