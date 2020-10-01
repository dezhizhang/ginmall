package main

import (
	"ginApi/src/utils"
	"github.com/gin-gonic/gin"
	"ginApi/src/controller"

)

func main() {
	
	
	router := gin.Default();
	

	//分组
	v1 := router.Group("/api/v1")
	{
		v1.GET("",controller.GetTopList)
		v1.POST("/create",controller.GetTopCreate)
		v1.POST("/add/user",controller.UserAdd)
		v1.GET("/user/info",controller.GetUser)
	}
	defer utils.DB.Close()
	router.Run();

	

}