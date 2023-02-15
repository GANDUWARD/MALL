package models

import (
	_ "github.com/jinzhu/gorm"
)

type Product struct {
	Id              int
	Title           string
	SubTitle        string
	ProductSn       string
	CateId          int
	ClickCount      int
	ProductNumber   int
	Price           float64
	MarketPrice     float64
	RelationProduct string
	ProductAttr     string
	ProductVersion  string
	ProductImg      string
	ProductGift     string
	ProductFitting  string
	ProductColor    string
	ProductKeywords string
	ProductDesc     string
	ProductContent  string
	IsDelete        int
	isHot           int
	IsBest          int
	isNew           int
	ProductTypeId   int
	Sort            int
	Status          int
	AddTime         int
}

func (Product) TableName() string {
	return "product"
}
