package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default();
	r.GET("/hello",func(c *gin.Context) {
		c.JSON(200,gin.H{
			"code":200,
			"msg":"成功",
			"success":true,
		})
	});
	r.Run(":8082")
}