package model

import (
	"exam/app/service/mysql"
)

type Course struct {
	Base
	Name string
}

func (receiver Course) CheckID(id int) bool {
	db := mysql.GetIns()
	result := db.First(&receiver, id)
	return result.RowsAffected > 0
}
