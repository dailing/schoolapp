package controllers

import (
	"io/ioutil"

	"encoding/json"
	"github.com/astaxie/beego"
)

type ImgGetController struct {
	beego.Controller
}

func (c *ImgGetController) Post() {
	beego.Trace("img get api")
	imgInfo := TypeImgResp{}
	body := c.Ctx.Input.CopyBody(1024 * 1024)
	err := json.Unmarshal(body, &imgInfo)
	ErrReport(err)
	if err != nil {
		c.Abort("400")
		return
	}
	img, err := ioutil.ReadFile(imgPath[0] + imgInfo.ImageID)
	ErrReport(err)
	if err != nil {
		c.Abort("500")
		return
	}
	//c.Ctx.ResponseWriter.Write(img)
	err = c.Ctx.Output.Body(img)
	ErrReport(err)

}
