package routers

import (
	"net/http"

	"git.oschina.net/dddailing/schoolapp/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/api/login", &controllers.LoginController{})
	beego.Router("/api/usr_add", &controllers.UserAddController{})
	beego.Router("/api/usr_get", &controllers.UserGetController{})
	beego.Router("/api/usr_update", &controllers.UserUpdateController{})

	beego.Router("/api/img_upload", &controllers.ImgUploadController{})
	beego.Router("/api/img_get", &controllers.ImgGetController{})

	beego.Router("/api/item_add", &controllers.ItemAddController{})
	beego.Router("/api/item_get", &controllers.ItemGetController{})
	beego.Router("/api/item_get_list", &controllers.ItemGetListController{})

	beego.ErrorHandler("404", serve404)
}

func serve404(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)
	w.Write([]byte("ERROR 404"))
}
