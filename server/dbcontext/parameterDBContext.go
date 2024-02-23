package dbcontext

import (
	"minora/data/dbmodel"
	"minora/server/server_errors"

	"gorm.io/gorm"
)

type ParameterDBContext struct {
	DB *gorm.DB
}

// func (context *ParameterDBContext) AddParameter(name string, parameterType string, isDefault bool) {
// 	context.DB.Create(&dbmodel.Parameter{
// 		Name:          name,
// 		ParameterType: parameterType,
// 		IsDefault:     isDefault,
// 		IsSystem:      false,
// 	})
// }

// func (context *ParameterDBContext) AddSystemParameter(name string, parameterType string, isDefault bool) {
// 	context.DB.Create(&dbmodel.Parameter{
// 		Name:          name,
// 		ParameterType: parameterType,
// 		IsDefault:     isDefault,
// 		IsSystem:      true,
// 	})
// }

func (context *ParameterDBContext) GetParameters() ([]dbmodel.CollectionItem, error) {
	var parameters []dbmodel.CollectionItem
	result := context.DB.Find(&parameters)
	return parameters, result.Error
}

func (context *ParameterDBContext) DeleteParameter(id uint) error {
	parameter := dbmodel.CollectionItem{ID: id}
	result := context.DB.First(&parameter)
	if result.RowsAffected == 0 {
		return server_errors.EntityNotFound
	}

	context.DB.Delete(parameter)
	return nil
}
