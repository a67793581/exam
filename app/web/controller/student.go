package controller

import (
	"exam/app/model"
	"exam/app/service/mysql"
	"github.com/gin-gonic/gin"
)

func StudentList(context *gin.Context) {
	db := mysql.NewConnection()
	var result []model.Student
	db.Model(&model.Student{}).Find(&result)
	context.JSON(200, result)
}
