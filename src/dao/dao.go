package dao

import (
	"github.com/gin-gonic/gin"
	
)

func GetTopDetail(c *gin.Context) {
	c.JSON(200,gin.H{
		"code":200,
		"msg":"成功",
		"success":true,
		"data":c.Param("id"),
	})
}

func GetTopList(c *gin.Context) {
	
}