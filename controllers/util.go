package controllers

import (
	"fmt"
	"math/rand"
	"time"

	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/dvsekhvalnov/jose2go"
	"strings"
)

func init() {
	rand.Seed(int64(time.Now().UnixNano()))
}

func Hash(str string) string {
	return str
}

func GenToken(info TypeTokenInfo) string {
	payload, err := json.Marshal(info)
	ErrReport(err)
	token, err := jose.Sign(string(payload), jose.NONE, nil)
	ErrReport(err)
	return token
}

func ParseToken(token string) TypeTokenInfo {
	beego.Info("token:", token)
	payload, _, err := jose.Decode(token, nil)
	ErrReport(err)
	info := TypeTokenInfo{}
	err = json.Unmarshal([]byte(payload), &info)
	ErrReport(err)
	if err != nil {
		info.UserName = ""
	}
	return info
}

func GenRandToken() string {
	return fmt.Sprintf("%016X", rand.Int63())
}

func GetTimeStamp() int {
	return int(time.Now().UnixNano())
}

func ErrReport(v interface{}) {
	if v != nil {
		beego.BeeLogger.Error("%v", v)
	}
}

func GenMataData() TypeMataData {
	return TypeMataData{
		TimeStamp: GetTimeStamp(),
		Device:    "server",
	}
}

func GenStatus(code int) TypeStatus {
	return TypeStatus{
		Code:        code,
		Description: ErrorDesp[code],
	}
}

func CheckToken(token string) bool {
	_, _, err := jose.Decode(token, nil)
	ErrReport(err)
	if err == nil {
		return true
	}
	return false
}

func getSqlSearchAixinwu(keywords string) string {
	strs := strings.Split(keywords, " ")
	retval := "select * from type_item_info where (name "
	for index, s := range strs {
		if index != len(strs)-1 {
			retval += fmt.Sprintf(" like '%s' OR name ", s)
		} else {
			retval += fmt.Sprintf(" like %s)", s)
		}
	}
	return retval
}

type InterfaceRandomSNData interface {
	GetSN() string
	SetSN(string)
}

func GenerateRandSN(randSNData *TypeAixinwuOrder) string {
	o := orm.NewOrm()
	stop := false
	for !stop {
		randSNData.SetSN(getDonationSN())
		if err := o.Read(randSNData, "order_sn"); err == orm.ErrNoRows {
			stop = true
		} else if err != nil {
			ErrReport(err)
			return ""
		}
	}
	return randSNData.GetSN()
}
