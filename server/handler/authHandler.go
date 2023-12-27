package handler

import (
	"errors"
	"keeper/data/request"
	"keeper/data/response"
	"keeper/dbcontext"
	"keeper/server"
	"keeper/utils"
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

	secret, secret_err := server.GetServerSecret()
	if secret_err != nil {
		return nil, errors.New("Server Internal Error")
	}

	var duration time.Duration = time.Duration(0)
	if !login.IsRemember {
		duration = time.Duration(24) * time.Hour
	}

	jwt := utils.GenerateJWTToken(duration, "", secret)

	return &response.LoginInfo{Login: login.Login, Token: jwt}, nil

}
