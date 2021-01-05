package router

import (
	"exam/app/web/controller"
	"github.com/gin-gonic/gin"
)

func api(engine *gin.Engine) {
	api := engine.Group("/api")
	{
		api.GET("test", controller.Test)
	}
}
