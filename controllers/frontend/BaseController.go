package frontend

import (
	"MALL/models"
	"fmt"
	"net/url"
	"strings"

	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
)

type BaseController struct {
	beego.Controller
}

func (c *BaseController) BaseInit() {
	//获取顶部导航
	topMenu := []models.Menu{}
	if hasTopMenu := models.CacheDb.Get("topMenu", &topMenu); hasTopMenu == true { //先获取数据，如果有简单在map添加数据就行
		c.Data["topMenuList"] = topMenu
	} else {
		models.DB.Where("status=1 AND position=1").Order("sort desc").Find(&topMenu) //这里方法是gorm的查询方法Where根据条件筛选Order根据顺序Find查找结果并将结果复制到引用内
		c.Data["topMenuList"] = topMenu
		models.CacheDb.Set("topMenu", topMenu)
	}
	//左侧分类，也就是预加载
	productCate := []models.ProductCate{}
	if hasProductCate := models.CacheDb.Get("productCate", &productCate); hasProductCate == true {
		c.Data["productCateList"] = productCate
	} else {
		models.DB.Preload("ProductCateItem", func(db *gorm.DB) *gorm.DB {
			return db.Where("product_cate.status=1").Order("product_cate.sort DESC")
		}).Where("pid=0 AND status=1").Order("sort desc", true).Find(&productCate)
		c.Data["productCateList"] = productCate
		models.CacheDb.Set("productCate", productCate)
	}
	//获取中间导航的数据,利用middleware
	middleMenu := []models.Menu{}
	if hasMiddleMenu := models.CacheDb.Get("middleMenu", &middleMenu); hasMiddleMenu == true {
		c.Data["middleMenuList"] = middleMenu
	} else {
		models.DB.Where("status=1 AND position=2").Order("sort desc").Find(&middleMenu)
		for i := 0; i < len(middleMenu); i++ {
			//获取相关商品
			middleMenu[i].Relation = strings.ReplaceAll(middleMenu[i].Relation, "，", ",") //把菜单内的中文逗号换成英文
			relation := strings.Split(middleMenu[i].Relation, ",")
			product := []models.Product{}                                                                                         //存放六个产品信息的切片
			models.DB.Where("id in (?)", relation).Limit(6).Order("sort ASC").Select("id,title,product_img,price").Find(&product) //查找六条内含id，title，img，price记录赋值到产品对象
			middleMenu[i].ProductItem = product
		}
		c.Data["middleMenuList"] = middleMenu
		models.CacheDb.Set("middleMenu", middleMenu)
	}
	//判断用户是否登录
	user := models.User{}
	models.Cookie.Get(c.Ctx, "userinfo", &user)
	if len(user.Phone) == 11 { //如果有登记电话说明已经注册弹出个人界面
		str := fmt.Sprintf(
			`<ul>
			<li class="userinfo">
				<a href="#">%v</a>
			
				<i class="i"></i>
				<ol>
					<li><a href="/user">个人中心</a></li>
					<li><a href="#">我的收藏</a></li>
					<li><a href="/auth/loginOut">退出登录</a></li>
				<ol>
			</li>
		</ul>`, user.Phone)
		c.Data["userinfo"] = str
	} else { //没有注册弹出登录界面
		str := fmt.Sprintf(
			`<ul>
				<li><a href="/auth/login" target="_blank">登录</a></li>
				<li>|</li>
				<li><a href="/auth/registerStep1" target="_blank">注册</a></li>
			</ul>`)
		c.Data["userinfo"] = str
	}
	urlPath, _ := url.Parse(c.Ctx.Request.URL.String()) //保存url路径名
	c.Data["pathname"] = urlPath.Path
}
