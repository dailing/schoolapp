package controllers

import (
	"io/ioutil"
	"os"
	"path/filepath"

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
	imgPath := "./uploadimgs/"
	beego.Trace("Recv Post")
	file, header, err := c.GetFile("binaryFile")
	ErrReport(err)
	beego.Trace(header.Filename)
	fileContent, err := ioutil.ReadAll(file)
	ErrReport(err)
	fileToken := GenRandToken()
	ioutil.WriteFile(imgPath+fileToken+filepath.Ext(header.Filename), fileContent, os.FileMode(0644))
	file.Close()
	resp := TypeUploadResp{
		MataData: GenMataData(),
		ImageID:  fileToken,
		Status:   GenStatus(StatusCodeOK),
	}
	c.Data["json"] = resp
	c.ServeJSON()
}
