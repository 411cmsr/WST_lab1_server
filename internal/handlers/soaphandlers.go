package handlers

import (
	"log"
	"net/http"

	"WST_lab1_server/internal/database"
	"WST_lab1_server/internal/models"
	"WST_lab1_server/internal/services"
)

func AddPersonHandler(request interface{}, w http.ResponseWriter, r *http.Request) (interface{}, error) {
	req := request.(*services.AddPersonRequest)
	log.Printf("Received AddPerson request: %+v\n", req)
	person := models.Person{Name: req.Name, Surname: req.Surname, Age: req.Age}

	db := database.GetDB()

	if err := db.Create(&person).Error; err != nil {
		log.Printf("Error adding person: %v\n", err)
		return nil, err
	}
	log.Printf("Person added successfully: %+v\n", person)
	return person, nil
}

func UpdatePersonHandler(request interface{}, w http.ResponseWriter, r *http.Request) (interface{}, error) {
	req := request.(*services.UpdatePersonRequest)
	log.Printf("Received UpdatePerson request: %+v\n", req)
	db := database.GetDB()

	var person models.Person
	if err := db.First(&person, req.ID).Error; err != nil {
		log.Printf("Error finding person with ID %d: %v\n", req.ID, err)
		return nil, err
	}

	person.Name = req.Name
	person.Surname = req.Surname
	person.Age = req.Age

	if err := db.Save(&person).Error; err != nil {
		log.Printf("Error updating person: %v\n", err)
		return nil, err
	}
	log.Printf("Person updated successfully: %+v\n", person)
	return person, nil
}

func DeletePersonHandler(request interface{}, w http.ResponseWriter, r *http.Request) (interface{}, error) {
	req := request.(*services.DeletePersonRequest)
	log.Printf("Received DeletePerson request: %+v\n", req)
	db := database.GetDB()

	if err := db.Delete(&models.Person{}, req.ID).Error; err != nil {
		log.Printf("Error deleting person with ID %d: %v\n", req.ID, err)
		return nil, err
	}
	log.Printf("Person with ID %d deleted successfully.\n", req.ID)
	return "Deleted successfully", nil
}

func GetPersonHandler(request interface{}, w http.ResponseWriter, r *http.Request) (interface{}, error) {
	req := request.(*services.GetPersonRequest)
	log.Printf("Received GetPerson request: %+v\n", req)
	var person models.Person

	db := database.GetDB()

	if err := db.First(&person, req.ID).Error; err != nil {
		log.Printf("Error finding person with ID %d: %v\n", req.ID, err)
		return nil, err
	}
	log.Printf("Retrieved person: %+v\n", person)
	return person, nil
}

func GetAllPersonsHandler(request interface{}, w http.ResponseWriter, r *http.Request) (interface{}, error) {
	log.Println("Received GetAllPersons request.")
	var persons []models.Person

	db := database.GetDB()

	if err := db.Find(&persons).Error; err != nil {
		log.Printf("Error retrieving all persons: %v\n", err)
		return nil, err
	}
	log.Printf("Retrieved all persons: %+v\n", persons)
	return services.GetAllPersonsResponse{Persons: persons}, nil
}

func SearchPersonHandler(request interface{}, w http.ResponseWriter, r *http.Request) (interface{}, error) {
	req := request.(*services.SearchPersonRequest)
	log.Printf("Received SearchPerson request with query: %s\n", req.Query)
	var persons []models.Person

	db := database.GetDB()

	if err := db.Where("name ILIKE ? OR surname ILIKE ?", "%"+req.Query+"%", "%"+req.Query+"%").Find(&persons).Error; err != nil {
		log.Printf("Error searching for persons with query '%s': %v\n", req.Query, err)
		return nil, err
	}
	log.Printf("Found persons matching query '%s': %+v\n", req.Query, persons)
	return services.GetAllPersonsResponse{Persons: persons}, nil
}
