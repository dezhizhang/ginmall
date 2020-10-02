package main

import (
	"fmt"
	"ginApi/src/controller"
	"ginApi/src/utils"

	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
)

func main() {

	router := gin.Default()

	conn := utils.RedisDefaultPool.Get()
	ret, err := redis.String(conn.Do("get", "username"))
	if err != nil {
		fmt.Println("出错了", err)
		return
	}
	fmt.Println("+++")
	fmt.Println(ret)
	fmt.Println("+++++")

	//分组
	v1 := router.Group("/api/v1")
	{
		v1.GET("", controller.GetTopList)
		v1.POST("/create", controller.GetTopCreate)
		v1.POST("/user/add", controller.UserAdd)
		v1.GET("/user/info", controller.GetUser)
		v1.GET("/user/info/:id", controller.GetUserOne)
	}
	defer utils.DB.Close()
	router.Run()

}
