package controllers

import (
	"fmt"
	"math/rand"
	"testing"

	"time"
)

func TestUserAdd(t *testing.T) {
	rand.Seed(int64(time.Now().UnixNano()))
	tdata := TypeUserInfo{
		Username: GenRandToken(),
		Password: "psw",
		NickName: GenRandToken(),
	}
	id, err := AddUser(tdata)
	fmt.Println(id, err)
}
