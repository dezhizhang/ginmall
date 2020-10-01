package main

import (
	"github.com/gin-gonic/gin"
	"ginApi/src/controller"
)

func main() {
	router := gin.Default();

	//分组
	v1 := router.Group("/api/v1/topics")
	{
		v1.GET("",controller.GetTopList);
		v1.GET("/:id",controller.GetTopDetail)
	}
	

	router.Run();

}