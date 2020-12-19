package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRouter() {
	engine := gin.Default()
	engine.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello World!")
	})
	Api(engine)
	err := engine.Run(":8088")
	if err != nil {
		panic(err)
	}
}
