package response

import "minora/data/dbmodel"

type CollectionValueResponse struct {
	dbmodel.CMSItem
	ID uint `json:"id"`
	AllowWrite bool `json:"allowWrite"`
}

type CollectionListResponse struct {
	Items []CollectionValueResponse `json:"items"`
}