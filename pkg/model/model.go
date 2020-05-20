package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type BaseModel struct {
	ID        string     `gorm:"primary_key,type:char(36)"`
	CreatedAt time.Time  `gorm:"index"`
	UpdatedAt time.Time  `gorm:"index"`
	DeletedAt *time.Time `gorm:"index"`
}

func (base *BaseModel) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("ID", uuid.New().String())
}
