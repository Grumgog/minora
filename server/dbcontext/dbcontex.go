package dbcontext

import (
	"fmt"
	"keeper/data/dbmodel"
	"keeper/server"
	"keeper/utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBContext struct {
	DB        *gorm.DB
	User      *UserDBContext
	Parameter *ParameterDBContext
}

var dbContext *DBContext = nil

// Получает соединение с базой данных по параметрам сервера.
func Connect() {
	properties := server.GetServerProperties()
	connString := getConnectionString(properties)
	db, err := gorm.Open(postgres.Open(connString), &gorm.Config{})
	utils.HandleErrorWithPanic(err)
	dbmodel.DBMigrate(db)
	dbContext = createContext(db)
	InitDB()
}

func createContext(db *gorm.DB) *DBContext {
	return &DBContext{
		DB:        db,
		User:      &UserDBContext{DB: db},
		Parameter: &ParameterDBContext{DB: db},
	}
}

// Возвращает контекст базы данных
func GetDBContext() *DBContext {
	return dbContext
}

func getConnectionString(props server.Properties) string {
	return fmt.Sprintf(
		"user=%s password=%s dbname=%s port=%s sslmode=disable",
		props.DB.DBUser,
		props.DB.DBPassword,
		props.DB.DBName,
		props.DB.DBPort)
}
