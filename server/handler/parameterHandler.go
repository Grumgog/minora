package handler

import (
	"keeper/data/dbmodel"
	"keeper/data/response"
	"keeper/dbcontext"
	"keeper/utils"
)

func ParameterGetListHandler() ([]response.Parameter, error) {
	db := dbcontext.GetDBContext()

	dbParameters, err := db.Parameter.GetParameters()

	if err == nil {
		return utils.Map(dbParameters, func(v dbmodel.Parameter) response.Parameter {
			return response.Parameter{
				ID:            v.ID,
				Name:          v.Name,
				ParameterType: v.ParameterType,
				IsDefault:     v.IsDefault,
				DefaultValue:  v.DefaultValue,
				IsSystem:      v.IsSystem,
			}
		}), nil
	}

	return nil, err
}

func ParameterCreateHandler(name string, paramType string, isDefault bool) {
	db := dbcontext.GetDBContext()

	db.Parameter.AddParameter(name, paramType, isDefault)
}

func ParameterGetValuesHandler(path string) response.ParameterValuesResponce {
	db := dbcontext.GetDBContext()

	var parameterValues []dbmodel.ParameterValue
	if path == "" {
		parameterValues = db.ParameterValue.GetParameterValuesList()
	} else {
		parameterValues = db.ParameterValue.GetParameterValuesListByPath(path)
	}

	responseData := utils.Map(parameterValues, func(values dbmodel.ParameterValue) response.ParameterValue {
		return response.ParameterValue{Value: values.Value, Path: values.Path, Name: values.Parameter.Name}
	})

	return response.ParameterValuesResponce{
		Parameters: responseData,
	}
}
