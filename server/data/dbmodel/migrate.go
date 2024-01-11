package dbmodel

import "gorm.io/gorm"

func DBMigrate(db *gorm.DB) {
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Parameter{})
	db.AutoMigrate(&ParameterValue{})
}
