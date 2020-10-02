package controller

import (
	"fmt"
	"time"
	"strconv"
	"github.com/gin-gonic/gin"
	"ginApi/src/model"
	"ginApi/src/utils"
)

func UserAdd(c *gin.Context) {
	user := model.User{}
	err := c.BindJSON(&user)
	if err != nil {
		c.String(200,"传入的参数有误")
		return
	}
	fmt.Println(time.Now())
	utils.DB.Create(&user)
	c.JSON(200,gin.H{
		"code":200,
		"msg":"添加用户成功",
		"success":true,
	})
}

func GetUser(c *gin.Context) {
	user := []model.User{}
	utils.DB.Find(&user)
	c.JSON(200,gin.H{
		"code":200,
		"msg":"成功",
		"success":true,
		"data":user,
	})
}

func GetUserOne(c *gin.Context) {
	id := c.Param("id")
	newId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("类型转换错误")
		return
	}
	user := model.User{ID:newId}
	utils.DB.Find(&user)
	c.JSON(200,gin.H{
		"code":200,
		"msg":"成功",
		"success":true,
		"data":user,
	})
}
