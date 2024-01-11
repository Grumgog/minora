package dbmodel

import "gorm.io/gorm"

type ParameterValue struct {
	gorm.Model
	Parameter Parameter `gorm:"primaryKey;foreignKey:ID"`
	Path      string    `gorm:"primaryKey"`
	Value     string
}
