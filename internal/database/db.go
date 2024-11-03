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
	config, err := config.LoadConfig(configFile)

	if err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		config.Database.Host,
		config.Database.User,
		config.Database.Password,
		config.Database.Name,
		config.Database.Port,
		config.Database.SSLMode)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	return db.AutoMigrate(&models.Person{})
}
func UpdateDB(configFile string) error {
	var err error
	config, err := config.LoadConfig(configFile)
	if err != nil {
		log.Fatalf("error loading config: %v", err)
	}
	db.Exec("DELETE FROM people")
	result := db.Create(&config.Persons)
	if result.Error != nil {
		return result.Error
	}
	return err
}

func GetDB() *gorm.DB {
	return db
}
