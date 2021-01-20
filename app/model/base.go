package model

type Base struct {
	ID        uint `gorm:"primary_key" json:"id"`
	CreatedAt int
	UpdatedAt int
	DeletedAt DeletedAt
}
