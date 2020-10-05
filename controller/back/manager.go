package back

import (
	"fmt"
	"ginApi/model"
	"ginApi/utils"

	"github.com/gin-gonic/gin"
)

func Manager(c *gin.Context) {
	c.HTML(200, "back/managerList.html", nil)
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
