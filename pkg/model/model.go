package model

import (
	"time"

	"github.com/qor/transition"
)

type BaseModel struct {
	ID                 string     `gorm:"primary_key,type:char(36)"`
	CreatedAt          time.Time  `sql:"index"`
	UpdatedAt          time.Time  `sql:"index"`
	DeletedAt          *time.Time `sql:"index"`
	EncryptionScopeKey string
}

type Dag struct {
	transition.Transition
}
