package routers

import (
	//"MALL/common"
	"MALL/controllers/frontend"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &frontend.IndexController{})
}
