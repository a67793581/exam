package main

import (
	"exam/app/model"
	"exam/app/service/mysql"
	"exam/app/web/router"
)

//初始化进程
func init() {

}

//web入口
func main() {
	db := mysql.GetIns()
	err := db.AutoMigrate(&model.Student{}, &model.ExamRecord{}, &model.Course{})
	if err != nil {
		print(err)
	}
	router.SetupRouter()
}
