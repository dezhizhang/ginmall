package front

import (
	"github.com/gin-gonic/gin"
)

func Cart(c *gin.Context) {
	c.HTML(200, "cart.html", nil)
}
