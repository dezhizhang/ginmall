package back

import (
	"github.com/gin-gonic/gin"
)

func Manager(c *gin.Context) {
	c.HTML(200, "back/managerList.html", nil)
}
