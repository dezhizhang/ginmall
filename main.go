package main

import (
	"github.com/gin-gonic/gin"
	"ginApi/src/dao"
)

func main() {
	router := gin.Default();

	//分组
	v1 := router.Group("/api/v1/topics")
	{
		v1.GET("",dao.GetTopList);
		v1.GET("/:id",dao.GetTopDetail)
	}
	

	router.Run();

}