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
		return utils.Map(dbParameters, func(v dbmodel.CollectionItem) response.Parameter {
			return response.Parameter{
				ID: v.ID,
				// Name:          v.Name,
				// ParameterType: v.ParameterType,
				// IsDefault:     v.IsDefault,
				// DefaultValue:  v.DefaultValue,
				// IsSystem:      v.IsSystem,
			}
		}), nil
	}

	return nil, err
}

func ParameterCreateHandler(name string, paramType string, isDefault bool) {
	//db := dbcontext.GetDBContext()

	//db.Parameter.AddParameter(name, paramType, isDefault)
}
