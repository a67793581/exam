package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetupRouter() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{}))
	//web前端文件
	e.Static("/", "../public")
	//api接口
	api(e)
	//日志
	e.Logger.Fatal(e.Start(":8088"))
}
