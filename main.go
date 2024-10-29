package main

import (
	"WST_lab1_server/database"
	"WST_lab1_server/services"
	"log"
)

func main() {
	//Определяем конфигурационный файл
	configPath := "configs/configs.yaml"
	//Подключаемся к базе данных(нужно создавать экземпляр?)
	db, err := database.InitializeDB(configPath)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	//Добавляем схему если нет, удаляем старые данные и добавляем данные Persons
	err = database.Migrate(db)
	if err != nil {
		log.Fatalf("failed to run migrations: %v", err)
	}
	//Запускаем soap-server (добавить wsdl-файл)
	err = services.RunSoapServer(db)
	if err != nil {
		log.Fatalf("failed to run migrations: %v", err)
	}

}
