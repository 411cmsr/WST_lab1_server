package handlers

import (
	"go.uber.org/zap"
	"net/http"

	"WST_lab1_server/internal/database"
	"WST_lab1_server/internal/models"
	"WST_lab1_server/internal/services"
)

var logger *zap.Logger

func init() {
	var err error
	logger, err = zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {

		}
	}(logger)
}

func AddPersonHandler(request interface{}, w http.ResponseWriter, r *http.Request) (interface{}, error) {
	req := request.(*services.AddPersonRequest)
	logger.Info("Received AddPerson request", zap.Any("request", req))
	person := models.Person{Name: req.Name, Surname: req.Surname, Age: req.Age}

	db := database.GetDB()

	if err := db.Create(&person).Error; err != nil {
		logger.Error("Error adding person", zap.Error(err))
		return nil, err
	}
	logger.Info("Person added successfully", zap.Any("person", person))
	return person, nil
}

func UpdatePersonHandler(request interface{}, w http.ResponseWriter, r *http.Request) (interface{}, error) {
	req := request.(*services.UpdatePersonRequest)
	logger.Info("Received UpdatePerson request", zap.Any("request", req))
	db := database.GetDB()

	var person models.Person
	if err := db.First(&person, req.ID).Error; err != nil {
		logger.Error("Error finding person with ID", zap.Uint("ID", req.ID), zap.Error(err))
		return nil, err
	}

	person.Name = req.Name
	person.Surname = req.Surname
	person.Age = req.Age

	if err := db.Save(&person).Error; err != nil {
		logger.Error("Error updating person", zap.Error(err))
		return nil, err
	}
	logger.Info("Person updated successfully", zap.Any("person", person))
	return person, nil
}

func DeletePersonHandler(request interface{}, w http.ResponseWriter, r *http.Request) (interface{}, error) {
	req := request.(*services.DeletePersonRequest)
	logger.Info("Received DeletePerson request", zap.Any("request", req))
	db := database.GetDB()

	if err := db.Delete(&models.Person{}, req.ID).Error; err != nil {
		logger.Error("Error deleting person with ID", zap.Uint("ID", req.ID), zap.Error(err))
		return nil, err
	}
	logger.Info("Person deleted successfully", zap.Uint("ID", req.ID))
	return "Deleted successfully", nil
}

func GetPersonHandler(request interface{}, w http.ResponseWriter, r *http.Request) (interface{}, error) {
	req := request.(*services.GetPersonRequest)
	logger.Info("Received GetPerson request", zap.Any("request", req))
	var person models.Person

	db := database.GetDB()

	if err := db.First(&person, req.ID).Error; err != nil {
		logger.Error("Error finding person with ID", zap.Uint("ID", req.ID), zap.Error(err))
		return nil, err
	}
	logger.Info("Retrieved person successfully", zap.Any("person", person))
	return person, nil
}

func GetAllPersonsHandler(request interface{}, w http.ResponseWriter, r *http.Request) (interface{}, error) {
	logger.Info("Received GetAllPersons request.")
	var persons []models.Person

	db := database.GetDB()

	if err := db.Find(&persons).Error; err != nil {
		logger.Error("Error retrieving all persons", zap.Error(err))
		return nil, err
	}
	logger.Info("Retrieved all persons successfully", zap.Any("persons", persons))
	return services.GetAllPersonsResponse{Persons: persons}, nil
}

func SearchPersonHandler(request interface{}, w http.ResponseWriter, r *http.Request) (interface{}, error) {
	req := request.(*services.SearchPersonRequest)
	logger.Info("Received SearchPerson request with query", zap.String("query", req.Query))
	var persons []models.Person

	db := database.GetDB()

	if err := db.Where("name ILIKE ? OR surname ILIKE ? OR age::text ILIKE ?", "%"+req.Query+"%", "%"+req.Query+"%", "%"+req.Query+"%").Find(&persons).Error; err != nil {
		logger.Error("Error searching for persons with query", zap.String("query", req.Query), zap.Error(err))
		return nil, err
	}
	if len(persons) == 0 {
		logger.Info("Search completed with no results", zap.String("query", req.Query))
	} else {
		logger.Info("Search completed successfully", zap.String("query", req.Query), zap.Int("count", len(persons)), zap.Any("results", persons))
	}
	return services.GetAllPersonsResponse{Persons: persons}, nil
}
