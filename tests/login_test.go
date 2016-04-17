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
	// add a test user
	info := controllers.TypeUserInfo{
		Username: controllers.GenRandToken(),
		Password: controllers.GenRandToken(),
		NickName: controllers.GenRandToken(),
		Phone:    controllers.GenRandToken(),
		Email:    controllers.GenRandToken(),
		Coins:    0,
	}
	id, err := controllers.AddUser(info)
	info.ID = id
	assert.Equal(t, err, nil)
	assert.T(t, id > 0)
	// retrieve user info again to check confidential
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
	// delete user from database
	succ, err = controllers.DelUser(info.Username)
	assert.T(t, succ)
	assert.Equal(t, err, nil)
	// retrieve user info again
	succ, _ = controllers.CheckUserNameExist(info.Username)
	assert.T(t, !succ)
}

func TestUserAPIs_http(t *testing.T) {
	regAddUser := controllers.TypeUserInfo{
		Username: controllers.GenRandToken(),
		Password: controllers.GenRandToken(),
		NickName: controllers.GenRandToken(),
		Phone:    controllers.GenRandToken(),
		Email:    controllers.GenRandToken(),
		Coins:    0,
	}
	req := controllers.TypeLoginInfo{
		MataData: controllers.GenMataData(),
		UserInfo: controllers.TypeUserInfo{
			Username: regAddUser.Username,
			Password: regAddUser.Password,
		},
	}
	// add a user
	beego.Trace("Adding user using web api")
	body, err := json.Marshal(regAddUser)
	assert.Equal(t, err, nil)
	r, err := http.NewRequest("POST", "/api/usr_add", bytes.NewReader(body))
	assert.Equal(t, err, nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)
	assert.Equal(t, w.Code, 200)
	assert.T(t, w.Body.Len() > 0)
	beego.Trace(string(w.Body.Bytes()))
	resp := controllers.TypeRegularResp{}
	err = json.Unmarshal(w.Body.Bytes(), &resp)
	assert.Equal(t, err, nil)
	assert.Equal(t, resp.Status.Code, controllers.StatusCodeOK)

	// login using web api
	beego.Trace("logging in using web api")
	body, err = json.Marshal(req)
	assert.Equal(t, err, nil)
	beego.Trace(string(body))
	r, err = http.NewRequest("POST", "/api/login", bytes.NewReader(body))
	assert.Equal(t, err, nil)
	w = httptest.NewRecorder()
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
