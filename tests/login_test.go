package test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"git.oschina.net/dddailing/schoolapp/controllers"
	"github.com/astaxie/beego"
	. "github.com/smartystreets/goconvey/convey"
)

func TestLogin(t *testing.T) {

	req := controllers.TypeLoginInfo{
		MataData: controllers.TypeMataData{
			TimeStamp: controllers.GetTimeStamp(),
			Device:    "test",
		},
		UserInfo: controllers.TypeUserInfo{
			Username: "test",
			Password: "psw",
		},
	}

	body, err := json.Marshal(req)
	Convey("No err Marshal request", t, func() {
		So(err, ShouldBeNil)
	})
	beego.Trace(string(body))
	r, _ := http.NewRequest("POST", "/login", bytes.NewReader(body))
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	//	beego.Trace("testing", "TestLogin", "Code", w.Code, "\n", w.Body)

	Convey("Subject: Test Station Endpoint\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
		responce, err := ioutil.ReadAll(w.Body)
		Convey("read should be succ", func() {
			So(err, ShouldBeNil)
		})
		respinfo := controllers.TypeLoginResp{}
		Convey("unmarshal should be succ", func() {
			So(json.Unmarshal(responce, &respinfo), ShouldBeNil)
		})
		Convey("Should get token", func() {
			So(len(respinfo.Token), ShouldBeGreaterThan, 0)
		})
	})
}
