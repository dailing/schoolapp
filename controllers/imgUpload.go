package controllers

import (
	"io/ioutil"
	"os"

	"github.com/astaxie/beego"
)

type ImgUploadController struct {
	beego.Controller
}

var imgPath = "./uploadimgs/"

func (c *ImgUploadController) Post() {
	beego.Trace("Recv Post")
	file, header, err := c.GetFile("binaryFile")
	ErrReport(err)
	beego.Trace(header.Filename)
	fileContent, err := ioutil.ReadAll(file)
	ErrReport(err)
	fileToken := GenRandToken()
	ioutil.WriteFile(imgPath+fileToken, fileContent, os.FileMode(0644))
	file.Close()
	resp := TypeImgResp{
		MataData: GenMataData(),
		ImageID:  fileToken,
		Status:   GenStatus(StatusCodeOK),
	}
	c.Data["json"] = resp
	beego.Trace("last step")
	c.ServeJSON()
}
