package handlers

import (
	"WST_lab1_server_new1/internal/database/postgres"
	"WST_lab1_server_new1/internal/models"
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"regexp"

	//"strconv"

	"github.com/gin-gonic/gin"
)

/*
Структура обработчика для разделения логики обработки запросов от доступа к данным
*/
type StorageHandler struct {
	Storage *postgres.Storage
}

/*
Функция проверки email на корректность
*/
func validateEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

/*
Функция проверки телефона на корректность
*/
func validatePhone(phone string) bool {
	re := regexp.MustCompile(`^\+7\d{10}$`)
	return re.MatchString(phone)
}

func (sh *StorageHandler) SOAPHandler(c *gin.Context) {
	var envelope models.Envelope

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error reading request body")
		return
	}
	c.Request.Body = io.NopCloser(bytes.NewBuffer(body))

	if err := xml.Unmarshal(body, &envelope); err != nil {
		fmt.Println("Error decoding XML:", err)
		c.String(http.StatusBadRequest, "Invalid request")
		return
	}

	// Печатаем содержание структуры в консоль
	fmt.Printf("Decoded Envelope: %+v\n", envelope)
	switch {
	case envelope.Body.AddPerson != nil:
		sh.addPersonHandler(envelope.Body.AddPerson)
		fmt.Println("envelope.Body.AddPerson")
	case envelope.Body.DeletePerson != nil:
		fmt.Println("envelope.Body.DeletePerson")
		sh.deletePersonHandler(envelope.Body.DeletePerson)
	case envelope.Body.UpdatePerson != nil:
		fmt.Println("envelope.Body.UpdatePerson")
		sh.updatePersonHandler(envelope.Body.UpdatePerson)
	case envelope.Body.GetPerson != nil:
		fmt.Println("envelope.Body.GetPerson")
		sh.getPersonHandler(envelope.Body.GetPerson)
	case envelope.Body.GetAllPersons != nil:
		fmt.Println("envelope.Body.GetAllPersons")
		sh.getAllPersonsHandler()
	case envelope.Body.SearchPerson != nil:
		fmt.Println("envelope.Body.SearchPerson")
		sh.searchPersonHandler(envelope.Body.SearchPerson)
	default:
		fmt.Println("Unsupported action")
		c.String(http.StatusBadRequest, "Unsupported action")
		return
	}

	c.String(http.StatusOK, "Request processed successfully")
}

func (h *StorageHandler) addPersonHandler(request *models.AddPersonRequest) {
	person := models.Person{
		Name:      request.Name,
		Surname:   request.Surname,
		Age:       request.Age,
		Email:     request.Email,
		Telephone: request.Telephone,
	}

	id, err := h.Storage.PersonRepository.AddPerson(&person)
	if err != nil {
		fmt.Printf("Error adding person: %v\n", err)
		return
	}

	fmt.Printf("Person added with ID: %d\n", id)
}

func (h *StorageHandler) updatePersonHandler(request *models.UpdatePersonRequest) {
	person := models.Person{
		ID:        uint(request.ID),
		Name:      request.Name,
		Surname:   request.Surname,
		Age:       request.Age,
		Email:     request.Email,
		Telephone: request.Telephone,
	}

	err := h.Storage.PersonRepository.UpdatePerson(&person)
	if err != nil {
		fmt.Printf("Error updating person with ID %d: %v\n", request.ID, err)
		return
	}

	fmt.Printf("Successfully updated person with ID: %d\n", request.ID)
}

func (h *StorageHandler) getPersonHandler(request *models.GetPersonRequest) {
	person, err := h.Storage.PersonRepository.GetPerson(request.ID)
	if err != nil {
		fmt.Printf("Error getting person with ID %d: %v\n", request.ID, err)
		return
	}

	if person == nil {
		fmt.Printf("No person found with ID %d\n", request.ID)
	} else {
		fmt.Printf("Retrieved person: %+v\n", person)
	}
}

func (h *StorageHandler) getAllPersonsHandler() {
	persons, err := h.Storage.PersonRepository.GetAllPersons()
	if err != nil {
		fmt.Printf("Error getting all persons: %v\n", err)
		return
	}

	fmt.Printf("Retrieved all persons: %+v\n", persons)
}

func (h *StorageHandler) deletePersonHandler(request *models.DeletePersonRequest) {
	// Создаем новый экземпляр DeletePersonRequest
	deleteRequest := &models.DeletePersonRequest{
		ID: request.ID,
	}
	checkByID, err := h.Storage.PersonRepository.CheckPersonByID(uint(deleteRequest.ID))
	if err != nil {
		fmt.Printf("Error getting person with ID %d: %v\n", request.ID, err)
		return
	} else if !checkByID {
		fmt.Println("Person not found")
		return

	}

	fmt.Println("deletePersonHandler: ", deleteRequest)
	// Вызываем метод DeletePerson из репозитория
	err = h.Storage.PersonRepository.DeletePerson(deleteRequest)
	if err != nil {

		fmt.Printf("Error deleting person with ID %d: %v\n", request.ID, err)
		return
	}

	fmt.Printf("Successfully deleted person with ID: %d\n", request.ID)
}

func (h *StorageHandler) searchPersonHandler(request *models.SearchPersonRequest) {
	persons, err := h.Storage.PersonRepository.SearchPerson(request.Query)
	if err != nil {
		fmt.Printf("Error searching for persons with query '%s': %v\n", request.Query, err)
		return
	}

	if len(persons) == 0 {
		fmt.Println("No persons found.")
	} else {
		fmt.Printf("Found persons: %+v\n", persons)
	}
}

