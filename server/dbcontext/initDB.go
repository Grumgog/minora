package dbcontext

import (
	"keeper/data/dbmodel"
	"keeper/utils"
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
		Username:                  "Admin",
		PasswordHash:              utils.HashPassword("admin"),
		IsSystem:                  true,
		IsChangePasswordNextLogin: true,
		CreatedAt:                 time.Now(),
	})

	db.Parameter.AddSystemParameter("title", "string", true)
	db.Parameter.AddSystemParameter("content", "string", true)
	db.Parameter.AddSystemParameter("previewImage", "path:image", true)
	db.Parameter.AddSystemParameter("author", "string", true)
	db.Parameter.AddSystemParameter("creationDate", "date", true)
	db.Parameter.AddSystemParameter("updateDate", "date", true)
}
