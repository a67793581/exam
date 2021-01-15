package controller

import (
	"exam/app/model"
	"exam/app/service/mysql"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Test(context *gin.Context) {
	name := context.DefaultQuery("name", "jack")
	context.JSON(200, gin.H{"name": name})
}

func TestMysql(context *gin.Context) {
	db := mysql.NewConnection()
	s := model.ExamRecord{Key: "test"}
	db.Create(&s)
	var result []model.Student
	db.Model(&model.Student{}).Find(&result)
	fmt.Println(result)

	context.JSON(200, result)
}
