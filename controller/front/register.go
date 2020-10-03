package front

import (
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	c.HTML(200, "front/register.html", nil)
}
