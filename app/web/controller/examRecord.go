package controller

import (
	"exam/app/model"
	"exam/app/service/mysql"
	"github.com/gin-gonic/gin"
)

func ExamRecordList(context *gin.Context) {
	db := mysql.NewConnection()
	var result []model.ExamRecord
	db.Model(&model.ExamRecord{}).Find(&result)
	context.JSON(200, result)
}
