package front

import (
	"github.com/gin-gonic/gin"
)

// func Order(c *gin.Context) {
// 	c.HTML(200, "order.html", nil)
// }

func Group(c *gin.Context) {
	c.HTML(200, "front/order.html", nil)
}
