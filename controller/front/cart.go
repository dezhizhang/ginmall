package front

import (
	"github.com/gin-gonic/gin"
)

func Cart(c *gin.Context) {
	c.HTML(200, "front/cart.html", nil)
}
