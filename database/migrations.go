package database

import (
	"WST_lab1_server/models"
	"gorm.io/gorm"
	"log"
)

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(&models.Persons{})
	if err != nil {
		return err
	}
	log.Println("Database migration succeeded")
	//Удаляем ранее добавленные в базу данные
	db.Exec("DELETE FROM persons")
	//Добавляем данные в базу
	persons := []models.Persons{
		{Name: "Петр", Surname: "Петров", Age: 25},
		{Name: "Владимир", Surname: "Иванов", Age: 26},
		{Name: "Иван", Surname: "Иванов", Age: 27},
		{Name: "Иммануил", Surname: "Кант", Age: 28},
		{Name: "Джордж", Surname: "Клуни", Age: 29},
		{Name: "Билл", Surname: "Рубцов", Age: 30},
		{Name: "Марк", Surname: "Марков", Age: 31},
		{Name: "Галина", Surname: "Матвеева", Age: 32},
		{Name: "Святослав", Surname: "Павлов", Age: 33},
		{Name: "Ольга", Surname: "Бергольц", Age: 34},
		{Name: "Лев", Surname: "Рабинович", Age: 35},
	}

	result := db.Create(&persons)
	if result.Error != nil {
		return err
	}
	return nil

}
