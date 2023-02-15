package models

import (
	"reflect"

	_ "github.com/jinzhu/gorm"
)

type Setting struct {
	Id              int    `form:"id"`
	SiteTitle       string `form:"site_tile"`
	SiteLogo        string `form:"site_log"`
	SiteKeywords    string `form:"site_keywords`
	SiteDescription string `form:"site_description"`
	NoPicture       string `form:"no_picture"`
	SiteIcp         string `form:"site_icp"`
	SiteTel         string `form:"site_tel"`
	SearchKeywords  string `form:"search_keywords"`
	TongjiCode      string `form:"tongji_code"`
	Appid           string `form:"appid"`
	AppSecret       string `form:"app_secret"`
	EndPoint        string `form:"end_point"`
	BucketName      string `form:"bucket_name"`
	OssStatus       int    `form:"oss_status"`
}

func (Setting) TableName() string {
	return "setting"
}

func GetSettingByColumn(columnName string) string {
	//redis file
	setting := Setting{}
	DB.First(&setting)
	//反射获取

	v := reflect.ValueOf(setting)
	val := v.FieldByName(columnName).String()
	return val
}
