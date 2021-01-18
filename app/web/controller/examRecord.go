package controller

import (
	"exam/app/model"
	"exam/app/service/mysql"
	"github.com/labstack/echo/v4"
)

func ExamRecordList(context echo.Context) error {
	db := mysql.NewConnection()
	var result []model.ExamRecord
	db.Model(&model.ExamRecord{}).Find(&result)
	return context.JSON(200, result)
}