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

}