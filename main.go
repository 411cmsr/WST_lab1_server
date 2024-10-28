package main

import (
	"WST_lab1_server/configs"
	"WST_lab1_server/migrations"
	"WST_lab1_server/models"
	"WST_lab1_server/services"
	"fmt"
	"log"
)

func main() {

	db, err := configs.InitializeDB()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	err = migrations.Migrate(db)
	if err != nil {
		log.Fatalf("failed to run migrations: %v", err)
	}
	log.Println("Database migration succeeded")

	var persons []models.Persons
	db.Find(&persons)
	for _, u := range persons {
		fmt.Println(u)
	}

	services.RunSoapServer(db)
}
