package frontend

import (
	"MALL/models"
)

type IndexController struct {
	BaseController
}

func (c *IndexController) Get() {
	//初始化
	c.BaseInit()
	//开始时间
	//startTime := time.Now().UnixNano()
	//获取轮播图
	banner := []models.Banner{}
	if hasBanner := models.CacheDb.Get(("banner"), &banner); hasBanner == true {
		c.Data["bannerList"] = banner
	} else {
		models.DB.Where("status=1 AND banner_type=1").Order("sort desc").Find(&banner)
		c.Data["bannerList"] = banner
		models.CacheDb.Set("banner", banner)
	}
	//获取手机类商品列表
	redisPhone := []models.Product{}
	if hasPhone := models.CacheDb.Get("phone", &redisPhone); hasPhone == true {
		c.Data["phoneList"] = redisPhone
	} else {
		phone := models.GetProductByCategory(1, "hot", 8)
		c.Data["phoneList"] = redisPhone
		models.CacheDb.Set("phone", phone)
	}
	//获取电视类商品列表
	redisTv := []models.Product{}
	if hasTv := models.CacheDb.Get("phone", &redisTv); hasTv == true {
		c.Data["tvList"] = redisTv
	} else {
		tv := models.GetProductByCategory(4, "best", 8)
		c.Data["tvList"] = redisTv
		models.CacheDb.Set("tv", tv)
	}
	//endTime := time.Now().UnixNano()
	c.TplName = "frontend/index/index.html"
}
