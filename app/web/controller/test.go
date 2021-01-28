package controller

import (
	"exam/app/model"
	"exam/app/service/mysql"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Test(context echo.Context) error {
	return context.String(http.StatusOK, "你输入的name："+context.QueryParam("name"))
}

func TestMysql(context echo.Context) error {
	db := mysql.GetIns()
	s := model.ExamRecord{Key: "test"}
	db.Create(&s)
	var result []model.Student
	db.Model(&model.Student{}).Find(&result)
	return context.JSON(http.StatusOK, result)
}
