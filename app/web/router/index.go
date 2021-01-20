package router

import (
	"exam/app/service/mysql"
	"exam/app/web/controller/graphql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetupRouter() {
	db := mysql.GetIns()
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{}))
	//web前端文件
	e.Static("/", "../public")
	//api接口
	api(e)
	//graphql
	h, _ := graphql.NewHandler(db)
	e.POST("/graphql", echo.WrapHandler(h))
	//日志
	e.Logger.Fatal(e.Start(":8088"))
}
