package dbmodel

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID                        uint   `gorm:"primaryKey;autoIncrement:true"`
	Username                  string `gorm:"unique"`
	PasswordHash              string
	Email                     string `gorm:"unique"`
	CreatedAt                 time.Time
	LastLogin                 time.Time
	LastActionAt              time.Time
	LastAction                uint
	Disabled                  bool
	IsChangePasswordNextLogin bool
	IsSystem                  bool
}
