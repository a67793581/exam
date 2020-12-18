package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Base struct {
	CreatedAt int
	UpdatedAt int
}

func (b *Base) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("CreatedAt", time.Now().Unix())
}

func (b *Base) BeforeUpdate(scope *gorm.Scope) error {
	return scope.SetColumn("UpdatedAt", time.Now().Unix())
}
