package handler

import (
	"errors"
	"minora/data/request"
	"minora/data/response"
	"minora/dbcontext"
	"minora/server"
	"minora/utils"
	"time"
)

func HandleAuth(login request.LoginRequest) (*response.LoginInfo, error) {
	db := dbcontext.GetDBContext()

	user, user_err := db.User.GetUserByLogin(login.Login)
	if user_err != nil {
		return nil, errors.New("Invalid password or login.")
	}

	if !utils.CompareHashPassword(user.PasswordHash, login.Password) {
		return nil, errors.New("Invalid password of login.")
	}

	secret := server.GetServerSecret()

	var duration time.Duration = time.Duration(0)
	if !login.IsRemember {
		duration = time.Duration(24) * time.Hour
	}

	jwt := utils.GenerateJWTToken(duration, "", secret)

	return &response.LoginInfo{Login: login.Login, Token: jwt}, nil

}
