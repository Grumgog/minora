package dbmodel

import "gorm.io/gorm"

type Collection struct {
	gorm.Model
	CMSItem
	ID           uint `gorm:"primaryKey;autoIncrement:true"`
	AllowWrite   bool
	FieldSet []CollectionField
}
