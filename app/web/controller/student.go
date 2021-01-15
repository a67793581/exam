package controller

import (
	"exam/app/model"
	"exam/app/service/mysql"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Student(context *gin.Context) {
	db := mysql.NewConnection()
	s := model.Student{Name: "carlo", Key: "test"}
	db.Create(&s)
	db.Delete(&s)
	var result []model.Student
	db.Model(&model.Student{}).Find(&result)
	fmt.Println(result)

	context.JSON(200, result)
}
