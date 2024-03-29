package dbcontext

import (
	"minora/data/dbmodel"
	"minora/utils"
	"time"
)

// Этот файл содержит код который иницилизирует базу данных и заполняет её значениями по умолчанию.
// К примеру создаёт базу данных и создаёт пользователя-админа.
func InitDB() {
	db := GetDBContext()
	// Если в Базе данных нет системного пользователя Admin, то база явно неиницилизирована.
	var user dbmodel.User = dbmodel.User{ID: 1}
	result := db.DB.First(&user)
	if result.RowsAffected != 0 {
		return
	}

	db.User.CreateUserByModel(&dbmodel.User{
		Username:                  "admin",
		PasswordHash:              utils.HashPassword("admin"),
		IsSystem:                  true,
		IsChangePasswordNextLogin: true,
		CreatedAt:                 time.Now(),
	})
}
