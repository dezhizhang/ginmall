package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default();

	//分组
	v1 := router.Group("/api/v1/topics")
	v1.GET("",func(c *gin.Context) {
		username := c.Query("username")
		if  username != "" {
			c.JSON(200,gin.H{
				"code":200,
				"msg":"成功",
				"success":true,
			});
			return
		}
		c.JSON(200,gin.H{
			"code":404,
			"msg":"失败",
			"success":false,
		})
	});
	v1.GET("/:id",func(c *gin.Context) {
		c.JSON(200,gin.H{
			"code":200,
			"msg":"成功",
			"success":true,
			"data":c.Param("id"),
		})
	})

	router.Run();

}