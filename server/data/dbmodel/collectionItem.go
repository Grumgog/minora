package dbmodel

import "gorm.io/gorm"

type CollectionItem struct {
	gorm.Model
	CMSItem
	ID            uint `gorm:"primaryKey;autoIncrement:true"`
	CollectionFieldID uint
	CollectionField CollectionField
	Value string
}
