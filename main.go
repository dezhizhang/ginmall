package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"ginApi/src/controller"
	"github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {


	router := gin.Default();
	db, err := gorm.Open("mysql", "root:701XTAY1993@/beego?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("数据库连接失败")
		return
	}

	//分组
	v1 := router.Group("/api/v1/topics")
	{
		v1.GET("",controller.GetTopList)
		v1.GET("/:id",controller.GetTopDetail)
		v1.POST("/create",controller.GetTopCreate)
	}
	

	router.Run()
	defer db.Close()

}