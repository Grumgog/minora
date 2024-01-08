package dbcontext

import (
	"keeper/data/dbmodel"

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
