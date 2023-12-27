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
	DB *gorm.DB
}

func Connect() {
	properties := server.GetServerProperties()
	connString := getConnectionString(properties)
	db, err := gorm.Open(postgres.Open(connString), &gorm.Config{})
	utils.HandleErrorWithPanic(err)
	dbmodel.DBMigrate(db)

}

func getConnectionString(props server.Properties) string {
	return fmt.Sprintf(
		"user=%s password=%s dbname=%s port=%s sslmode=disable",
		props.DBUser,
		props.DBPassword,
		props.DBName,
		props.DBPort)
}
