package back

import (
	"fmt"
	"ginmall/model"
	"ginmall/utils"

	"github.com/gin-gonic/gin"
)

func Manager(c *gin.Context) {
	managerList := []model.Manager{}
	err := utils.DB.Find(&managerList).Error
	if err != nil {
		fmt.Println("获取数据失败", err)
		return
	}
	c.HTML(200, "back/managerList.html", gin.H{
		"managerList": managerList,
	})
}

func ManagerAdd(c *gin.Context) {
	c.HTML(200, "back/managerAdd.html", nil)
}

func ManagerDoAdd(c *gin.Context) {
	manager := model.Manager{}
	err := c.Bind(&manager)
	if err != nil {
		fmt.Println("帮定失败", err)
		return
	}
	err = utils.DB.Create(&manager).Error
	if err != nil {
		fmt.Println("添加失败", err)
		return
	}

}
