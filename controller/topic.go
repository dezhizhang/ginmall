package controller

import (
	"ginmall/src/model"

	"github.com/gin-gonic/gin"
)

func GetTopDetail(c *gin.Context) {
	topic := model.Topic{
		Id:    1,
		Title: "新闻",
	}
	c.JSON(200, gin.H{
		"code":    200,
		"msg":     "成功",
		"success": true,
		"data":    topic,
	})
}

func GetTopList(c *gin.Context) {
	query := model.TopicList{}
	err := c.BindQuery(&query)
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"msg":     "请求错误",
			"success": false,
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    200,
		"msg":     "成功",
		"success": true,
		"data":    query,
	})
}

func GetTopCreate(c *gin.Context) {
	query := model.Topic{}
	err := c.BindJSON(&query)
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"msg":     "请求错误",
			"success": false,
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    200,
		"msg":     "成功",
		"success": true,
		"data":    query,
	})
}
