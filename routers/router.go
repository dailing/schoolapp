package routers

import (
	"net/http"

	"git.oschina.net/dddailing/schoolapp/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})

	beego.ErrorHandler("404", serve404)
}

func serve404(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)
	w.Write([]byte("ERROR 404"))
}
