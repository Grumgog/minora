package response

import "minora/data/dbmodel"

type FieldSetResponse struct {
	dbmodel.CMSItem
	ID uint `json:"id"`
	Type string `json:"type"`
	DefaultValue string `json:"defaultValue"`
	CollectionId uint `json:"collectionId"`
}

type GetCollectionResponse struct {
	dbmodel.CMSItem
	ID uint `json:"id"`
	AllowWrite bool `json:"allowWrite"`
	FieldSet []FieldSetResponse
}