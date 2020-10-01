package controller

import (
	"github.com/gin-gonic/gin"
	"ginApi/src/model"
	
)

func GetTopDetail(c *gin.Context) {
	topic := model.Topic{
		Id:1,
		Title:"新闻",
	}
	c.JSON(200,gin.H{
		"code":200,
		"msg":"成功",
		"success":true,
		"data":topic,
	})
}

func GetTopList(c *gin.Context) {
	query := model.TopicList{}
	err := c.BindQuery(&query)
	if err != nil {
		c.JSON(404,gin.H{
			"code":"404",
			"msg":"请求错误",
			"success":false,
		});
		return
	}
	c.JSON(200,gin.H{
		"code":200,
		"msg":"成功",
		"success":true,
		"data":query,
	})
	
}