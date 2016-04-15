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
	"github.com/bmizerany/assert"
)

func TestUserAPIs(t *testing.T) {
	info := controllers.TypeUserInfo{
		Username: controllers.GenRandToken(),
		Password: controllers.GenRandToken(),
		NickName: controllers.GenRandToken(),
		Phone:    controllers.GenRandToken(),
		Email:    controllers.GenRandToken(),
		Coins:    0,
	}
	id, err := controllers.AddUser(info)
	assert.Equal(t, err, nil)
	assert.T(t, id > 0)
	succ, err := controllers.CheckUserNameExist(info.Username)
	assert.T(t, succ)
	assert.Equal(t, err, nil)
	rInfo, err := controllers.GetUserInfo(info.Username)
	assert.Equal(t, err, nil)
	assert.Equal(t, rInfo.Username, info.Username)
	assert.Equal(t, rInfo.ID, info.ID)
	assert.Equal(t, rInfo.Password, info.Password)
	assert.Equal(t, rInfo.Coins, info.Coins)
	assert.Equal(t, rInfo.NickName, info.NickName)
}

func _TestLogin(t *testing.T) {

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
	assert.Equal(t, err, nil)
	beego.Trace(string(body))
	r, err := http.NewRequest("POST", "/api/login", bytes.NewReader(body))
	assert.Equal(t, err, nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)
	assert.Equal(t, w.Code, 200)
	assert.T(t, w.Body.Len() > 0)
	response, err := ioutil.ReadAll(w.Body)
	respInfo := controllers.TypeLoginResp{}
	assert.Equal(t, err, nil)
	assert.Equal(t, json.Unmarshal(response, &respInfo), nil)
	assert.T(t, len(respInfo.Token) > 0)
}

func TestCheckUserNameFunc(t *testing.T) {
	succ, err := controllers.CheckUserNameLegal("asdfgADsad+das$#23")
	assert.T(t, succ)
	assert.Equal(t, err, nil)
	succ, err = controllers.CheckUserNameLegal("af fsd fs fds")
	assert.T(t, succ)
	assert.Equal(t, err, nil)
}
