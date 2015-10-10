package routers

import (
	"htz_blog/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/admin", &controllers.AdminController{})
    beego.Router("/admin/login", &controllers.AdminController{},"get,post:Login")
}
