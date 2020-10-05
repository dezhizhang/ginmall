package main

import (
	"ginmall/controller/back"
	"ginmall/controller/front"
	"ginmall/utils"
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
		b.POST("/manager/doAdd", back.ManagerDoAdd)
		// b.GET("//manager/delete", back.ManagerDelete)
		b.GET("/role", back.Role)
		b.GET("/role/add", back.RoleAdd)
		b.POST("/role/doAdd", back.RoleDoAdd)
		b.GET("/role/delete", back.RoleDelete)

	}

	//小程序api部分
	// api := router.Group("/api/v1")

	defer utils.DB.Close()
	router.Run()

}
