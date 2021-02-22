package router

import (
	"exam/app/service/token_jwt"
	"exam/app/web/controller"
	"exam/app/web/controller/graphql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func api(e *echo.Echo) {
	api := e.Group("/api")
	{
		//graphql
		g := api.Group("/graphql")
		{
			teacher := g.Group("/teacher")
			{
				teacher.Use(middleware.JWTWithConfig(middleware.JWTConfig{
					Claims:     &token_jwt.Claims{},
					SigningKey: []byte(token_jwt.Key),
				}))
				handler, _ := graphql.Teacher()
				teacher.Any("", echo.WrapHandler(handler))
			}

			Student := g.Group("/student")
			{
				handler, _ := graphql.Student()
				Student.Any("", echo.WrapHandler(handler))
			}
		}

		api.Any("/test", controller.Test)
		api.GET("/test_mysql", controller.TestMysql)
		student := api.Group("/student")
		{
			student.GET("/list", controller.StudentList)
		}

		teacher := api.Group("/teacher")
		{
			// Login route
			teacher.POST("/login", controller.Login)
			teacher.POST("/import", controller.Import)
		}

	}
}
