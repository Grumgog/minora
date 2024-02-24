package dbcontext

import (
	"minora/data/dbmodel"
	"minora/server/server_errors"

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

func (context *CollectionDBContext) GetById(id uint) (*dbmodel.Collection, error) {
	collection := dbmodel.Collection{ID: id}
	state := context.DB.Find(&collection)

	if state.RowsAffected == 0 {
		return nil, server_errors.ErrEntityNotFound
	}

	return &collection, nil
}