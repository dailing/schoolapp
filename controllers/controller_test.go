package controllers

import (
	"fmt"
	"github.com/bmizerany/assert"
	"testing"
)

func TestStaticGetMoney(t *testing.T) {
	info := getStaticInfo()
	fmt.Println("return value:", info)
	assert.T(t, info.Item != "")
	assert.T(t, info.User != "")
	assert.T(t, info.Visit != "")
	assert.T(t, info.Money != "")
}
