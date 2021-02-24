package model

type Base struct {
	ID        uint      `gorm:"primary_key;comment:自增id"`
	CreatedAt int       `gorm:"index;comment:创建时间"`
	UpdatedAt int       `gorm:"comment:更新时间"`
	DeletedAt DeletedAt `gorm:"comment:删除时间"`
}
