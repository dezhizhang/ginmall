package main

import (
	"ginApi/controller/back"
	"ginApi/controller/front"
	"ginApi/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.StaticFS("static", http.Dir("./static"))
	router.LoadHTMLGlob("./views/**/*")

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
	b := router.Group("/admin")
	{
		b.GET("/home", back.Home)
		b.GET("/login", back.Login)
		b.GET("/manager", back.Manager)
		b.GET("/manager/add", back.ManagerAdd)
		b.GET("/role", back.Role)
		b.GET("/role/add", back.RoleAdd)

	}

	//小程序api部分
	// api := router.Group("/api/v1")

	defer utils.DB.Close()
	router.Run()

}
