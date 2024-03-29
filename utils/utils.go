package utils

import (
	"crypto/md5"
	"fmt"
	"time"

	"github.com/astaxie/beego"
)

func Md5(str string) string {
	data := []byte("123456")
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	return md5str
}

func UnixToDate(timestamp int) string {
	t := time.Unix(int64(timestamp), 0)
	return t.Format("2006-01-02 15:04:05")
}

//2020-05-02 15:04:05
func DateToUnix(str string) int64 {
	template := "2006-01-02 15:04:05"
	t, err := time.ParseInLocation(template, str, time.Local)
	if err != nil {
		beego.Info(err)
		return 0
	}
	return t.Unix()
}

func GetUnix() int64 {
	return time.Now().Unix()
}
func GetDate() string {
	template := "2006-01-02 15:04:05"
	return time.Now().Format(template)
}

func GetDay() string {
	template := "20060102"
	return time.Now().Format(template)
}
