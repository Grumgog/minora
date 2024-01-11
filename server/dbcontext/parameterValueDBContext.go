package dbcontext

import (
	"keeper/data/dbmodel"
	"keeper/server"

	"gorm.io/gorm"
)

type ParameterValueDBContext struct {
	DB *gorm.DB
}

func (context *ParameterValueDBContext) GetSystemParameterValues() map[string]string {
	return map[string]string{"AppName": server.GetAppName()}
}

func (context *ParameterValueDBContext) GetParameterValuesList() []dbmodel.ParameterValue {
	var parameters []dbmodel.ParameterValue
	context.DB.Find(&parameters)
	return parameters
}

func (context *ParameterValueDBContext) GetParameterValuesListByPath(path string) []dbmodel.ParameterValue {
	var parameters []dbmodel.ParameterValue
	context.DB.Where(&dbmodel.ParameterValue{Path: path}).Find(&parameters)
	return parameters
}

func (context *ParameterValueDBContext) CreateParameterValue(path string, parameterId uint, value string) {
	context.DB.Create(&dbmodel.ParameterValue{Parameter: dbmodel.Parameter{ID: parameterId}, Path: path, Value: value})
}
