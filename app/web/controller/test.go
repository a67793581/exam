package controller

import (
	"bytes"
	"encoding/csv"
	"exam/app/model"
	"exam/app/service/mysql"
	"github.com/axgle/mahonia"
	"github.com/labstack/echo/v4"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func Test(context echo.Context) error {
	var Buf bytes.Buffer
	// 根据字段名获取表单文件
	formFile, FileHeader, err := context.Request().FormFile("uploadFile")
	if err != nil {
		if err.Error() == "http: no such file" {
			return context.String(http.StatusBadRequest, "文件不存在")
		}
		return err
	}
	defer formFile.Close()
	//获取文件的内容

	_, _ = io.Copy(&Buf, formFile)
	//用gbk进行解码。
	decoder := mahonia.NewDecoder("gbk")
	contents := decoder.ConvertString(Buf.String())

	r := csv.NewReader(strings.NewReader(contents))
	records, err := r.ReadAll()
	if err != nil {
		return err
	}

	////文件大小检测
	if FileHeader.Size > 100000 {
		return context.String(http.StatusBadRequest, "文件太大，当前大小："+strconv.FormatInt(FileHeader.Size, 10))
	}

	//文件类型检测
	fileBytes, err := ioutil.ReadAll(formFile)
	if err != nil {
		return err
	}
	detectedFileType := http.DetectContentType(fileBytes)
	switch detectedFileType {
	case "text/plain; charset=utf-8":
		break
	default:
		return context.String(http.StatusBadRequest, "文件类型不合法："+detectedFileType)
	}

	var res = make(map[string]interface{})
	res["FileHeader"] = FileHeader
	res["contents"] = contents
	res["detectedFileType"] = detectedFileType
	res["records"] = records
	//for k, v := range records {
	//	for k2, v2 := range v {
	//
	//	}
	//}

	return context.JSON(http.StatusOK, res)
}

func TestMysql(context echo.Context) error {
	db := mysql.GetIns()
	s := model.ExamRecord{Key: "test"}
	db.Create(&s)
	var result []model.Student
	db.Model(&model.Student{}).Find(&result)
	return context.JSON(http.StatusOK, result)
}
