package dbcontext

import (
	"keeper/data/dbmodel"

	"gorm.io/gorm"
)

type UserDBContext struct {
	DB *gorm.DB
}

func (context *UserDBContext) CreateUserByModel(model *dbmodel.User) {
	context.DB.Create(model)
}

func (context *UserDBContext) CreateUserByParams() {

}
