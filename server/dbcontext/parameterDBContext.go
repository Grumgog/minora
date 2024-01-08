package dbcontext

import (
	"keeper/data/dbmodel"
	"keeper/server/server_errors"

	"gorm.io/gorm"
)

type ParameterDBContext struct {
	DB *gorm.DB
}

func (context *ParameterDBContext) AddParameter(name string, parameterType string, isDefault bool) {
	context.DB.Create(&dbmodel.Parameter{
		Name:          name,
		ParameterType: parameterType,
		IsDefault:     isDefault,
		IsSystem:      false,
	})
}

func (context *ParameterDBContext) AddSystemParameter(name string, parameterType string, isDefault bool) {
	context.DB.Create(&dbmodel.Parameter{
		Name:          name,
		ParameterType: parameterType,
		IsDefault:     isDefault,
		IsSystem:      true,
	})
}

func (context *ParameterDBContext) GetParameters() ([]dbmodel.Parameter, error) {
	var parameters []dbmodel.Parameter
	result := context.DB.Find(&parameters)
	return parameters, result.Error
}

func (context *ParameterDBContext) DeleteParameter(id uint) error {
	parameter := dbmodel.Parameter{ID: id}
	result := context.DB.First(&parameter)
	if result.RowsAffected == 0 {
		return server_errors.EntityNotFound
	}

	if parameter.IsSystem {
		return server_errors.NotProcessiableEntity
	}

	context.DB.Delete(parameter)
	return nil
}
