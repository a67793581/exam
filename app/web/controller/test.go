package controller

import (
	"exam/app/model"
	"exam/app/service/mysql"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"net/http"
	"strconv"
)

func Test(context echo.Context) error {
	//获取http.Request
	//读取上传文件
	FileHeader, err := context.FormFile("uploadFile")
	if err != nil {
		return err
	}
	//文件大小检测
	if FileHeader.Size > 100 {
		return context.String(http.StatusBadRequest, "文件太大，当前大小："+strconv.FormatInt(FileHeader.Size, 10))
	}
	File, err := FileHeader.Open()
	if err != nil {
		return err
	}
	fileBytes, err := ioutil.ReadAll(File)
	if err != nil {
		return err
	}
	//文件类型检测
	detectedFileType := http.DetectContentType(fileBytes)
	switch detectedFileType {
	case "text/plain; charset=utf-8":
		break
	default:
		return context.String(http.StatusBadRequest, "文件类型不合法："+detectedFileType)
	}

	return context.JSON(http.StatusOK, fileBytes)
}

func TestMysql(context echo.Context) error {
	db := mysql.GetIns()
	s := model.ExamRecord{Key: "test"}
	db.Create(&s)
	var result []model.Student
	db.Model(&model.Student{}).Find(&result)
	return context.JSON(http.StatusOK, result)
}
