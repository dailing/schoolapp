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
	assert.Equal(t, err, nil)
	beego.Trace(string(body))
	r, _ := http.NewRequest("POST", "/login", bytes.NewReader(body))
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
