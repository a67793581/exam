package controller

import (
	"exam/app/model"
	"exam/app/service/mysql"
	"github.com/labstack/echo/v4"
)

func StudentList(context echo.Context) error {
	db := mysql.NewConnection()
	var result []model.Student
	db.Model(&model.Student{}).Find(&result)
	return context.JSON(200, result)
}
