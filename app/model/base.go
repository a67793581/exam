package model

type Base struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt int
	UpdatedAt int
	DeletedAt DeletedAt
}
