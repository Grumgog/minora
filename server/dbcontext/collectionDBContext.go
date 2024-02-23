package dbcontext

import (
	"minora/data/dbmodel"

	"gorm.io/gorm"
)

type CollectionDBContext struct {
	DB *gorm.DB
}

func (context *CollectionDBContext) GetList() []dbmodel.Collection {
	var collections []dbmodel.Collection
	context.DB.Find(&collections)
	return collections
}