package routers

import (
	"MALL/common"
	"MALL/controllers/backend"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/"+beego.AppConfig.String("adminPath"),
		beego.NSRouter("/", &backend.MainController{}),
		beego.NSRouter("/welcome", &backend.MainController{}, "get:Welcome"),
		beego.NSRouter("/main/changestatus", &backend.MainController{}, "get:ChangeStatus"),
		beego.NSRouter("/main/editnum", &backend.MainController{}, "get:EditNum"),
		beego.NSBefore(common.BaseAuth),
		beego.NSRouter("/login", &backend.LoginController{}),
		beego.NSRouter("/login/gologin", &backend.LoginController{}, "post:GoLogin"),
		beego.NSRouter("/login/loginout", &backend.LoginController{}, "get:LoginOut"),
		//商品管理
		beego.NSRouter("/product", &backend.ProductController{}),
		//beego.NSRouter("/product/add", &backend.ProductController{}, "get:Add"),
	)
	beego.AddNamespace(ns)
}
