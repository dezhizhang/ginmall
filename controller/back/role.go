package back

import (
	"github.com/gin-gonic/gin"
)

func Role(c *gin.Context) {
	c.HTML(200, "back/roleList.html", nil)
}

func RoleAdd(c *gin.Context) {
	c.HTML(200, "back/roleAdd.html", nil)
}
