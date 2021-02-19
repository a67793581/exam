package controller

import (
	"bytes"
	"encoding/csv"
	"exam/app/service/token_jwt"
	"github.com/axgle/mahonia"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func Login(c echo.Context) error {
	key := c.FormValue("key")

	if key != "carlo" {
		return echo.ErrUnauthorized
	}
	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, token_jwt.Claims{
		ID:       0,
		Identity: "admin",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	})

	// Generate encoded token and send it as response.
	//t, err := token.SigningString()
	t, err := token.SignedString([]byte(token_jwt.Key))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, echo.Map{"token": t})
}

func Import(context echo.Context) error {
	var Buf bytes.Buffer
	// 根据字段名获取表单文件
	formFile, FileHeader, err := context.Request().FormFile("uploadFile")
	if err != nil {
		if err.Error() == "http: no such file" {
			return context.JSON(http.StatusInternalServerError,
				map[string]interface{}{
					"message": "文件不存在",
				},
			)
		}
		return err
	}
	defer formFile.Close()
	//获取文件的内容

	_, _ = io.Copy(&Buf, formFile)
	//用gbk进行解码。
	decoder := mahonia.NewDecoder("gbk")
	contents := decoder.ConvertString(Buf.String())

	////文件大小检测
	if FileHeader.Size > 100000 {
		return context.JSON(http.StatusInternalServerError,
			map[string]interface{}{
				"message": "文件太大，当前大小：" + strconv.FormatInt(FileHeader.Size, 10),
			},
		)
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
		return context.JSON(http.StatusInternalServerError,
			map[string]interface{}{
				"message": "文件类型不合法：" + detectedFileType,
			},
		)
	}
	//读取文件内容到数组
	r := csv.NewReader(strings.NewReader(contents))
	records, err := r.ReadAll()
	if err != nil {
		return err
	}
	var res = make(map[string]interface{})
	res["FileHeader"] = FileHeader
	res["contents"] = contents
	res["detectedFileType"] = detectedFileType
	res["records"] = records
	//执行业务逻辑
	//for k, v := range records {
	//	for k2, v2 := range v {
	//
	//	}
	//}

	return context.JSON(http.StatusOK, res)
}
