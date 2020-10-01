

package utils

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)
var DB *gorm.DB
var err error


func init() {
	DB,err = gorm.Open("mysql", "root:701XTAY1993@/gin?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("数据库连接失败",err);
		return
	}
}

