package routers

import (
	"github.com/astaxie/beego"
	"blog/controllers/blog"
	"blog/controllers/admin"
)

func init() {

	//前台路由
	beego.Router("/", &blog.MainController{}, "*:Index")
	beego.Router("/404.html", &blog.MainController{}, "*:Go404")
	beego.Router("/index:page:int.html", &blog.MainController{}, "*:Index")


	//后台路由
	beego.Router("/login",&admin.AccountController{})
}
