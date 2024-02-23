package dbmodel

import "gorm.io/gorm"

type CollectionField struct {
	gorm.Model
	CMSItem
	Type string
	DefaultValue  string
	IsSystem      bool
	CollectionID uint
}