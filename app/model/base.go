package model

type Base struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt int
	UpdatedAt int
	DeletedAt DeletedAt
}
