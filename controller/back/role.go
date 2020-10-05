package back

import (
	"fmt"
	"ginApi/model"
	"ginApi/utils"

	"github.com/gin-gonic/gin"
)

func Role(c *gin.Context) {
	roleList := []model.Role{}
	err := utils.DB.Find(&roleList).Error
	if err != nil {
		fmt.Println("获取数据失败", err)
		return
	}
	c.HTML(200, "back/roleList.html", gin.H{
		"roleList": roleList,
	})
}

func RoleAdd(c *gin.Context) {
	c.HTML(200, "back/roleAdd.html", nil)
}

func RoleDoAdd(c *gin.Context) {
	role := model.Role{}
	err := c.Bind(&role)
	if err != nil {
		fmt.Println("绑定失败", err)
		return
	}
	role.AddTime = int(utils.GetUnix())
	role.Status = 1
	err = utils.DB.Create(&role).Error
	if err != nil {
		fmt.Println("插入数据失败", err)
		return
	}
}
