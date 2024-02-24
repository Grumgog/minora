package dbmodel

import "gorm.io/gorm"

type CollectionField struct {
	gorm.Model
	CMSItem
	Type string
	DefaultValue  string
	CollectionID uint
}