package backend

import (
	"MALL/models"
	"math"
)

//var wg sync.WaitGroup

type ProductController struct {
	BaseController
}

func (c *ProductController) Get() {
	page, _ := c.GetInt("page")
	if page == 0 {
		page = 1
	}
	pageSize := 5
	keyword := c.GetString("keyword")
	where := "1=1" //全部数据，SQL注入
	if len(keyword) > 0 {
		where += "AND title like \"%" + keyword + "%\""
	}
	productList := []models.Product{}
	models.DB.Where(where).Offset((page - 1) * pageSize).Limit(pageSize).Find(&productList)
	var count int
	models.DB.Where(where).Table("product").Count(&count)
	c.Data["productList"] = productList
	c.Data["totalPages"] = math.Ceil(float64(count) / float64(pageSize))
	c.Data["page"] = page
	c.TplName = "backend/product/index.html"
}
