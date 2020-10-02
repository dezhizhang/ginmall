package front

import (
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	c.HTML(200, "register.html", nil)
}
