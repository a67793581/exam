package model

import (
	"exam/app/service/mysql"
)

type Course struct {
	Base
	Name string `gorm:"index;comment:课程"`
}

func (receiver Course) CheckID(id int) bool {
	db := mysql.GetIns()
	result := db.First(&receiver, id)
	return result.RowsAffected > 0
}
