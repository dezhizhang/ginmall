package main

import (
	"ginApi/src/controller/front"
	"ginApi/src/utils"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.Static("/static", "./static")
	router.LoadHTMLGlob("./src/views/*")

	//前端pc
	f := router.Group("")
	{
		f.GET("/", front.Home)
	}
	//管理后台部分
	// admin := router.Group("/admin")

	//小程序api部分
	// api := router.Group("/api/v1")

	defer utils.DB.Close()
	router.Run()

}
