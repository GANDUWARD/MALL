package models

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

// 定义私有缓存结构体
type cookie struct{}

// 写入数据的方法
func (c cookie) Set(ctx *context.Context, key string, value interface{}) {
	bytes, _ := json.Marshal(value)
	ctx.SetSecureCookie(beego.AppConfig.String("secureCookie"), key, string(bytes), 3600*24*30, "/", beego.AppConfig.String("domain"), nil, true)
}
