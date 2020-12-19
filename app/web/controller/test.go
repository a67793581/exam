package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Test(c *gin.Context) {
	name := c.DefaultQuery("name", "jack")
	c.String(200, fmt.Sprintf("hello %s\n", name))
}
