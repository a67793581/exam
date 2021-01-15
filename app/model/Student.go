package model

import (
	"gorm.io/gorm"
)

type Student struct {
	Base
	Name string
	Key  string
}

func (o *Student) BeforeSave(db *gorm.DB) (err error) {
	println("保存前")
	return
}
func (o *Student) BeforeUpdate(db *gorm.DB) (err error) {
	println("更新前")
	return
}
func (o *Student) AfterUpdate(db *gorm.DB) (err error) {
	println("更新后")
	return
}
func (o *Student) AfterSave(db *gorm.DB) (err error) {
	println("保存后")
	return
}
func (o *Student) BeforeDelete(db *gorm.DB) (err error) {
	println("删除前")
	return
}
func (o *Student) AfterDelete(db *gorm.DB) (err error) {
	println("删除后")
	return
}
