package dbmodel

import "gorm.io/gorm"

func DBMigrate(db *gorm.DB) {
	db.AutoMigrate(&CollectionField{})
	db.AutoMigrate(&Collection{})
	db.AutoMigrate(&CollectionItem{});

	db.AutoMigrate((&User{}))
}
