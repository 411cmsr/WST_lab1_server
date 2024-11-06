package database

import (
	"WST_lab1_server/internal/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"

	"WST_lab1_server/config"
)

var db *gorm.DB

func InitDB(configFile string) error {
	var err error
	configuration, err := config.LoadConfig(configFile)

	if err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		configuration.Database.Host,
		configuration.Database.User,
		configuration.Database.Password,
		configuration.Database.Name,
		configuration.Database.Port,
		configuration.Database.SSLMode)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	return db.AutoMigrate(&models.Person{})
}
func UpdateDB(configFile string) error {
	var err error
	configuration, err := config.LoadConfig(configFile)
	if err != nil {
		log.Fatalf("error loading config: %v", err)
	}
	db.Exec("DELETE FROM people")
	result := db.Create(&configuration.Persons)
	if result.Error != nil {
		return result.Error
	}
	return err
}

func GetDB() *gorm.DB {
	return db
}
