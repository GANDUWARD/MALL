package models

import (
	_ "github.com/jinzhu/gorm"
)

type OrderItem struct {
	Id             int
	OrderId        int
	Uid            int
	ProductTitle   string
	ProductId      int
	ProductNum     int
	ProductImg     string
	ProductPrice   float64
	ProductVersion string
	ProductColor   string
	AddTime        int
}

func (OrderItem) TableName() string {
	return "order_item"
}
