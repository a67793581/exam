package router

import (
	"exam/app/web/controller"
	"github.com/gin-gonic/gin"
)

func api(engine *gin.Engine) {
	api := engine.Group("api")
	{
		api.GET("test", controller.Test)
		api.GET("test_mysql", controller.TestMysql)
		student := api.Group("student")
		{
			student.GET("list", controller.StudentList)
		}

		examRecord := api.Group("exam_record")
		{
			examRecord.GET("list", controller.ExamRecordList)
		}

	}
}
