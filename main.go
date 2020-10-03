package main

import (
	"ginApi/controller/front"
	"ginApi/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.StaticFS("static", http.Dir("./static"))
	router.LoadHTMLGlob("./views/*")

	//前端pc
	f := router.Group("")
	{
		f.GET("/", front.Home)
		f.GET("/index", front.Home)
		f.GET("/login", front.Login)
		f.GET("/register", front.Register)
		f.GET("/order", front.Order)
		f.GET("/cart", front.Cart)
		f.GET("/group", front.Group)
	}
	//管理后台部分
	// admin := router.Group("/admin")

	//小程序api部分
	// api := router.Group("/api/v1")

	defer utils.DB.Close()
	router.Run()

}
