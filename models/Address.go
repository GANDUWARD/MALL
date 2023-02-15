package models

import (
	_ "github.com/jinzhu/gorm"
)

type Address struct {
	Id             int    `json:"id"`
	Uid            int    `json:"uid"`
	Phone          string `json:"phone"`
	Name           string `json:"name"`
	Address        string `json:"address"`
	Zipcode        string `json:"zipcode"`
	DefaultAddress string `json:"default_address"`
	Addtime        string `json:"add_time"`
}

func (Address) TableName() string {
	return "address"
}
