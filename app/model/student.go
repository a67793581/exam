package model

import (
	"exam/app/service/mysql"
	"gorm.io/gorm"
)

type Student struct {
	Base
	Name string `gorm:"index;comment:姓名"`
	Key  string `gorm:"index;comment:学号"`
}

func (receiver *Student) BeforeSave(db *gorm.DB) (err error) {
	println("保存前")
	return
}

func (receiver *Student) BeforeUpdate(db *gorm.DB) (err error) {
	println("更新前")
	return
}

func (receiver *Student) AfterUpdate(db *gorm.DB) (err error) {
	println("更新后")
	return
}

func (receiver *Student) AfterSave(db *gorm.DB) (err error) {
	println("保存后")
	return
}

func (receiver *Student) BeforeDelete(db *gorm.DB) (err error) {
	println("删除前")
	return
}

func (receiver *Student) AfterDelete(db *gorm.DB) (err error) {
	println("删除后")
	return
}

func (receiver Student) CheckID(id int) bool {
	db := mysql.GetIns()
	result := db.First(&receiver, id)
	return result.RowsAffected > 0
}
