package router

import (
	"exam/app/web/controller"
	"github.com/labstack/echo/v4"
)

func api(e *echo.Echo) {
	api := e.Group("/api")
	{
		api.GET("/test", controller.Test)
		api.GET("/test_mysql", controller.TestMysql)
		student := api.Group("/student")
		{
			student.GET("/list", controller.StudentList)
		}

		examRecord := api.Group("/exam_record")
		{
			examRecord.GET("/list", controller.ExamRecordList)
			examRecord.GET("/details", controller.ExamRecordDetails)
		}

	}
}
