package dbmodel

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID                        uint `gorm:"primaryKey;autoIncrement:true"`
	Username                  string
	PasswordHash              string
	Email                     string
	CreatedAt                 time.Time
	LastLogin                 time.Time
	LastActionAt              time.Time
	LastAction                uint
	Disabled                  bool
	IsChangePasswordNextLogin bool
	IsSystem                  bool
}
