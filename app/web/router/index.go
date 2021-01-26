package router

import (
	"exam/app/web/controller/graphql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetupRouter() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{}))
	//web前端文件
	e.Static("/", "../public")
	//api接口
	api(e)
	//graphql
	h, _ := graphql.Teacher()
	e.POST("/graphql/teacher", echo.WrapHandler(h))
	Student, _ := graphql.Student()
	e.POST("/graphql/student", echo.WrapHandler(Student))
	//日志
	e.Logger.Fatal(e.Start(":8088"))
}
