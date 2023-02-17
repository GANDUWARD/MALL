package common

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"math/rand"
	"path"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/gomarkdown/markdown"
	"github.com/hunterhug/go_image"
	_ "github.com/jinzhu/gorm"
)

// 将时间戳转换为日期格式
func TimestampToDate(timestamp int) string {
	t := time.Unix(int64(timestamp), 0)
	return t.Format("2006-01-02 15:04:05")
}

// 获取当前时间戳
func GetUnix() int64 {
	t := time.Now().Unix()
	fmt.Println(t)
	return t
}

// 获取时间戳的Nano时间
func GetUnixNano() int64 {
	return time.Now().Unix()
}

func GetDate() string {
	t := "2006-01-02 15:04:05"
	return time.Now().Format(t)
}

// Md5加密
func Md5(str string) string {
	m := md5.New()
	m.Write([]byte(str))
	return string(hex.EncodeToString(m.Sum(nil)))
}

// 验证邮箱
func VerifyEmail(email string) bool {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

// 获取日期
func FormatDay() string {
	t := "20060102"
	return time.Now().Format(t)
}

// 生成订单号
func GenerateOrderId() string {
	t := "200601021504"
	return time.Now().Format(t) + GetRandomNum()
}

// 发送验证码
func SendMsg(str string) {
	//短信验证码需要到相关网站申请，为了方便先固定一个值
	ioutil.WriteFile("text_send.TXT", []byte(str), 06666)
}

// 重新裁剪图片
func ResizeImage(filename string) {
	extName := path.Ext(filename) //获取文件拓展名包括.
	resizeImage := strings.Split(beego.AppConfig.String("resizeImageSize"), ",")
	for i := 0; i < len(resizeImage); i++ {
		w := resizeImage[i]
		width, _ := strconv.Atoi(w)
		savepath := filename + "_" + w + "x" + w + extName
		err := go_image.ThumbnailF2F(filename, savepath, width, width)
		if err != nil {
			beego.Error(err)
		}
	}
}

// 格式化图片
func FormatImage(picName string) string {
	ossStatus, err := beego.AppConfig.Bool("ossStatus")
	if err != nil {
		//判断目录前面是否有"/"
		f := strings.Contains(picName, "/static")
		if f {
			return picName
		}
		return "/" + picName
	} else if ossStatus {
		return beego.AppConfig.String("ossDomain" + "/" + picName)
	} else {
		//再次判断前面是否有"/"
		f := strings.Contains(picName, "/static")
		if f {
			return picName
		}
		return "/" + picName
	}
}

// 格式化级标题
func FormatAttribute(str string) string {
	md := []byte(str)
	htmlByte := markdown.ToHTML(md, nil, nil)
	return string(htmlByte)
}

// 计算乘法
func Mul(price float64, num int) float64 {
	return price * float64(num)
}

// 封装一个生产随机数的方法
func GetRandomNum() string {
	var str string
	for i := 0; i < 4; i++ {
		current := rand.Intn(10) //0-9
		str += strconv.Itoa(current)
	}
	return str
}
