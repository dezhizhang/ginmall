package back

import (
	"fmt"
	"ginApi/model"
	"ginApi/utils"

	"github.com/gin-gonic/gin"
)

func Role(c *gin.Context) {
	c.HTML(200, "back/roleList.html", nil)
}

func RoleAdd(c *gin.Context) {
	c.HTML(200, "back/roleAdd.html", nil)
}

func RoleDoAdd(c *gin.Context) {
	query := model.Role{}
	err := c.Bind(&query)
	if err != nil {
		fmt.Println("绑定失败", err)
		return
	}
	query.AddTime = int(utils.GetUnix())
	query.Status = 1
	err = utils.DB.Create(&query).Error
	if err != nil {
		fmt.Println("插入数据失败", err)
		return
	}
}
