package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email     string         `gorm:"uniqueIndex;not null"`
	Username  string         `gorm:"uniqueIndex;not null"`
	Password  string         `gorm:"not null"` // should be hashed
	Role      string         `gorm:"default:'user'"` // or enum in DB
	IsActive  bool           `gorm:"default:true"`
	IsVerified  bool         `gorm:"default:false"`
	LastLogin *time.Time
}
