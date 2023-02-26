package routers

import (
	//"MALL/common"
	"MALL/controllers/frontend"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &frontend.IndexController{})
	beego.Router("/auth/sendCode", &frontend.AuthController{}, "get:SendCode")
	beego.Router("/auth/doRegister", &frontend.AuthController{}, "post:GoRegister")
	beego.Router("/auth/validateSmsCode", &frontend.AuthController{}, "get:ValidateSmsCode")
	beego.Router("/auth/login", &frontend.AuthController{}, "get:Login")
	beego.Router("/auth/registerStep1", &frontend.AuthController{}, "get:RegisterStep1")
	beego.Router("/auth/registerStep2", &frontend.AuthController{}, "get:RegisterStep2")
	beego.Router("/auth/registerStep3", &frontend.AuthController{}, "get:RegisterStep3")
	beego.Router("/auth/goLogin", &frontend.AuthController{}, "post:GoLogin")
	beego.Router("/auth/loginOut", &frontend.AuthController{}, "get:LoginOut")
	beego.Router("/user", &frontend.UserController{})
	beego.Router("/user/order", &frontend.UserController{}, "get:OrderList")
	beego.Router("/user/orderinfo", &frontend.UserController{}, "get:OrderInfo")
	beego.Router("/category_:id([0-9]+).html", &frontend.ProductController{}, "get:CategoryList")
	beego.Router("/item_:id([0-9]+).html", &frontend.ProductController{}, "get:ProductItem")
	beego.Router("/product/getImgList", &frontend.ProductController{}, "get:GetImgList")
	beego.Router("/product/collect", &frontend.ProductController{}, "get:Collect")
	beego.Router("/cart", &frontend.CartController{})
	beego.Router("/cart/addCart", &frontend.CartController{}, "get:AddCart")
	beego.Router("/cart/incCart", &frontend.CartController{}, "get:IncCart")
	beego.Router("/cart/decCart", &frontend.CartController{}, "get:DecCart")
	beego.Router("/cart/delCart", &frontend.CartController{}, "get:DelCart")
	beego.Router("/cart/changeOneCart", &frontend.CartController{}, "get:ChangeOneCart")
	beego.Router("/cart/changeAllCart", &frontend.CartController{}, "get:ChangeAllCart")
}
