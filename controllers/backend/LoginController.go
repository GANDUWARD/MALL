package backend

import (
	"MALL/common"
	"MALL/models"
	"strings"
)

type LoginController struct {
	BaseController
}

func (c *LoginController) Get() {
	c.TplName = "backend/login/login.html"
}

func (c *LoginController) GoLogin() {
	var flag = models.Cpt.VerifyReq(c.Ctx.Request)
	if flag {
		username := strings.Trim(c.GetString("username"), "")
		password := common.Md5(strings.Trim(c.GetString("password"), ""))
		administrator := []models.Administrator{}
		models.DB.Where("username=? AND password=? AND status=1", username, password).Find(&administrator)
		if len(administrator) == 1 {
			c.SetSession("userinfo", administrator[0])
			c.Success("登录成功", "/")
		} else {
			c.Error("无登录权限或用户名密码错误", "/login")
		}
	} else {
		c.Error("验证码错误", "/login")
	}
}

func (c *LoginController) LoginOut() {
	c.DelSession("userinfo")
	c.Success("退出登录成功,将返回登录页面！", "/login")
}
