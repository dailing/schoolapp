package controllers

import (
	"io/ioutil"
	"os"

	"github.com/astaxie/beego"
)

type UploadController struct {
	beego.Controller
}

type TypeUploadResp struct {
	MataData TypeMataData `json:"mataData"`
	ImageID  string       `json:"imageID"`
	Status   TypeStatus   `json:"status"`
}

func (c *UploadController) Post() {
	beego.Trace("Recv Post")
	file, header, err := c.GetFile("binaryFile")
	ErrReport(err)
	beego.Trace(header.Filename)
	filecontant, err := ioutil.ReadAll(file)
	ErrReport(err)
	ioutil.WriteFile("/tmp/img/"+header.Filename, filecontant, os.FileMode(0644))
	file.Close()
	resp := TypeUploadResp{
		MataData: GenMataData(),
		ImageID:  GenRandToken(),
		Status:   GenStatus(StatusCodeOK),
	}
	c.Data["json"] = resp
	c.ServeJSON()
}
