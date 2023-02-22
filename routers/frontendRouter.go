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

}
