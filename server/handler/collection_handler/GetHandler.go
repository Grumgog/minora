package collection_handler

import (
	"minora/data/dbmodel"
	"minora/data/response"
	"minora/dbcontext"
	"minora/utils"
)

func GetList() response.CollectionListResponse {
	db := dbcontext.GetDBContext()

	collectionList := db.Collection.GetList()

	
	result := utils.Map(collectionList, func (item dbmodel.Collection) response.CollectionValueResponse {
		return response.CollectionValueResponse{
			ID: item.ID,
			CMSItem: dbmodel.CMSItem{ApiName: item.ApiName, DisplayName: item.DisplayName},
			AllowWrite: item.AllowWrite,
		}
	})

	return response.CollectionListResponse{
		Items: result,
	}
}