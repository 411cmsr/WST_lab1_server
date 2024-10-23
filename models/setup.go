package models

import (
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"gorm.io/gorm"
)

func SetupModels() *gorm.DB {
	db, err := gorm.Open("sqlite3", "test.db")

	if err != nil {
		panic("Failed to connect to database!")
	}

	db.AutoMigrate(&Persons{})

	return db
}
