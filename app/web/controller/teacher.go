package controller

import (
	"bytes"
	"encoding/csv"
	"exam/app/model"
	"exam/app/service/mysql"
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
	decoder := mahonia.NewDecoder("utf-8")
	contents := decoder.ConvertString(strings.Replace(Buf.String(), "\uFEFF", "", 1))

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

	//执行业务逻辑
	mapStudent := make(map[string]model.Student)
	var whereStudent []interface{}
	var batchStudent []model.Student

	mapCourse := make(map[string]model.Course)
	var whereCourse []interface{}
	var batchCourse []model.Course

	for k, v := range records {
		if k == 0 {
			continue
		}
		for k2, v2 := range v {
			if v2 == "" {
				return context.JSON(http.StatusInternalServerError,
					map[string]interface{}{
						"message": "第" + strconv.Itoa(k) + "行数据,第" + strconv.Itoa(k2) + "格内容禁止为空",
					},
				)
			}
		}

		//'考试编号', '考试时间', '成绩', '考试批次','课程','学生姓名'

		if _, ok := mapStudent[v[6]]; !ok {
			batchStudent = append(batchStudent, model.Student{
				Name: v[5],
				Key:  v[6],
			})
			whereStudent = append(whereStudent, v[6])
			mapStudent[v[6]] = model.Student{}
		}

		if _, ok := mapCourse[v[4]]; !ok {
			batchCourse = append(batchCourse, model.Course{
				Name: v[4],
			})
			whereCourse = append(whereCourse, v[4])
			mapCourse[v[4]] = model.Course{}
		}

	}

	db := mysql.GetIns()

	var Students []model.Student
	db.Where("`key` IN ?", whereStudent).Find(&Students)
	for _, v := range Students {
		for k2, v2 := range batchStudent {
			if v.Key == v2.Key {
				batchStudent = append(batchStudent[:k2], batchStudent[(k2+1):]...)
			}
		}
	}
	db.Create(&batchStudent)
	batchStudent = append(batchStudent, Students...)
	for _, v := range batchStudent {
		if _, ok := mapStudent[v.Key]; ok {
			mapStudent[v.Key] = v
		}
	}

	var Courses []model.Course
	db.Where("`name` IN ?", whereCourse).Find(&Courses)
	for _, v := range Courses {
		for k2, v2 := range batchCourse {
			if v.Name == v2.Name {
				batchCourse = append(batchCourse[:k2], batchCourse[(k2+1):]...)
			}
		}
	}
	db.Create(&batchCourse)
	batchCourse = append(batchCourse, Courses...)
	for _, v := range batchCourse {
		if _, ok := mapCourse[v.Name]; ok {
			mapCourse[v.Name] = v
		}
	}

	mapExamRecord := make(map[string]model.ExamRecord)
	var whereExamRecord []interface{}
	var batchExamRecord []model.ExamRecord
	for k, v := range records {
		if k == 0 {
			continue
		}
		//'考试编号', '考试时间', '成绩', '考试批次','课程','学生姓名'
		Achievement, err := strconv.Atoi(v[2])
		if err != nil {
			return err
		}
		ExamTime, err := strconv.Atoi(v[1])
		if err != nil {
			return err
		}

		if _, ok := mapExamRecord[v[0]]; !ok {
			batchExamRecord = append(batchExamRecord, model.ExamRecord{
				Key:         v[0],
				ExamTime:    int32(ExamTime),
				Achievement: int32(Achievement),
				Code:        v[3],
				CourseID:    int32(mapCourse[v[4]].ID),
				StudentID:   int32(mapStudent[v[6]].ID),
			})
			whereExamRecord = append(whereExamRecord, v[0])
			mapExamRecord[v[0]] = model.ExamRecord{}
		}

	}

	var ExamRecords []model.ExamRecord
	db.Where("`key` IN ?", whereExamRecord).Find(&ExamRecords)
	for _, v := range ExamRecords {
		for k2, v2 := range batchExamRecord {
			if v.Key == v2.Key {
				batchExamRecord = append(batchExamRecord[:k2], batchExamRecord[(k2+1):]...)
			}
		}
	}
	db.Create(&batchExamRecord)
	batchExamRecord = append(batchExamRecord, ExamRecords...)
	var res = make(map[string]interface{})
	res["records"] = records
	res["mapStudent"] = mapStudent
	res["mapCourse"] = mapCourse
	res["batchExamRecord"] = batchExamRecord
	return context.JSON(http.StatusOK, res)
}
