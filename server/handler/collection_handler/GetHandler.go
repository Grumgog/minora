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

func GetById(id uint) (*response.GetCollectionResponse, error) {
	db := dbcontext.GetDBContext()
	collection, err := db.Collection.GetById(id)
	
	if err != nil {
		return nil, err
	}	
	return &response.GetCollectionResponse{
		CMSItem: dbmodel.CMSItem{ApiName: collection.ApiName, DisplayName: collection.DisplayName},
		ID: collection.ID,
		AllowWrite: collection.AllowWrite,
		FieldSet: utils.Map(collection.FieldSet, func(field dbmodel.CollectionField) response.FieldSetResponse {
			return response.FieldSetResponse{
				CMSItem: field.CMSItem,
				ID: field.ID,
				Type: field.Type,
				DefaultValue: field.DefaultValue,
				CollectionId: field.CollectionID,
			}
		}),
	}, nil
}