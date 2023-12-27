package dbmodel

import "gorm.io/gorm"

type Parameter struct {
	gorm.Model
	ID            uint `gorm:"primaryKey;autoIncrement:true"`
	Name          string
	ParameterType string
	IsDefault     bool
	DefaultValue  string
	IsSystem      bool
}
