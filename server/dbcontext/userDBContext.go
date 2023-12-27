package dbcontext

import (
	"errors"
	"keeper/data/dbmodel"

	"gorm.io/gorm"
)

type UserDBContext struct {
	DB *gorm.DB
}

func (context *UserDBContext) CreateUserByModel(model *dbmodel.User) {
	context.DB.Create(model)
}

func (context *UserDBContext) GetUserByLogin(login string) (*dbmodel.User, error) {
	user := dbmodel.User{
		Username: login,
	}
	state := context.DB.First(&user)
	if state.RowsAffected == 0 {
		return nil, errors.New("No such user")
	}

	return &user, nil
}
