package handlers

import (
	"WST_lab1_server_new1/internal/database/postgres"
	"WST_lab1_server_new1/internal/logging"
	"WST_lab1_server_new1/internal/models"
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"regexp"

	//"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	//"gorm.io/gorm/logger"
)

/*
Структура обработчика для разделения логики обработки запросов от доступа к данным
*/
type StorageHandler struct {
	Storage *postgres.Storage
}

func createSOAPFault(code string, message string, errorCode string, errorMessage string) models.SOAPFault {
	fault := models.SOAPFault{}
	fault.Envelope.Body.Fault.Code = code
	fault.Envelope.Body.Fault.Message = message
	fault.Envelope.Body.Fault.Detail.ErrorCode = errorCode
	fault.Envelope.Body.Fault.Detail.ErrorMessage = errorMessage
	return fault
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

	fmt.Printf("Decoded Envelope: %+v\n", envelope)

	switch {
	case envelope.Body.AddPerson != nil:
		sh.addPersonHandler(c, envelope.Body.AddPerson)
	case envelope.Body.DeletePerson != nil:
		sh.deletePersonHandler(c, envelope.Body.DeletePerson) // Передаем контекст и запрос
	case envelope.Body.UpdatePerson != nil:
		sh.updatePersonHandler(c, envelope.Body.UpdatePerson)
	case envelope.Body.GetPerson != nil:
		sh.getPersonHandler(c, envelope.Body.GetPerson) // Передаем контекст и запрос
	case envelope.Body.GetAllPersons != nil:
		sh.getAllPersonsHandler(c)
	case envelope.Body.SearchPerson != nil:
		sh.searchPersonHandler(c, envelope.Body.SearchPerson) // Передаем контекст и запрос
	default:
		fmt.Println("Unsupported action")
		c.String(http.StatusBadRequest, "Unsupported action")
		return
	}
}

func (h *StorageHandler) addPersonHandler(c *gin.Context, request *models.AddPersonRequest) {
	// Создаем нового человека на основе запроса
	person := models.Person{
		Name:      request.Name,
		Surname:   request.Surname,
		Age:       request.Age,
		Email:     request.Email,
		Telephone: request.Telephone,
	}

	// Добавляем человека в базу данных
	id, err := h.Storage.PersonRepository.AddPerson(&person)
	if err != nil {
		fmt.Printf("Error adding person: %v\n", err)

		// Формируем SOAP Fault для ошибки добавления
		fault := createSOAPFault("soap:Server", "Internal Server Error", "500", "An unexpected error occurred.")
		c.XML(http.StatusInternalServerError, fault)
		return
	}

	fmt.Printf("Person added with ID: %d\n", id)

	// Формируем успешный ответ
	response := models.AddPersonResponse{
		ID: id,
	}

	// Возвращаем успешный ответ в формате XML
	c.XML(http.StatusOK, response)
}

func (h *StorageHandler) updatePersonHandler(c *gin.Context, request *models.UpdatePersonRequest) {
	// Проверяем, существует ли человек с данным ID
	checkByID, err := h.Storage.PersonRepository.CheckPersonByID(uint(request.ID))
	if err != nil {
		logging.Logger.Error("Error checking person with ID", zap.Uint("ID", uint(request.ID)), zap.Error(err))

		fault := createSOAPFault("soap:Server", "Internal Server Error", "500", "An unexpected error occurred.")
		c.XML(http.StatusInternalServerError, fault)
		return
	}

	// Если человек не найден, формируем SOAP Fault для клиента
	if !checkByID {
		fault := createSOAPFault("soap:Client", models.ErrorRecordNotFoundMessage, models.ErrorRecordNotFoundCode, models.ErrorRecordNotFoundDetail)
		c.XML(http.StatusNotFound, fault)
		return
	}

	// Создаем объект Person на основе запроса
	person := models.Person{
		ID:        uint(request.ID),
		Name:      request.Name,
		Surname:   request.Surname,
		Age:       request.Age,
		Email:     request.Email,
		Telephone: request.Telephone,
	}

	// Обновляем информацию о человеке в базе данных
	err = h.Storage.PersonRepository.UpdatePerson(&person)
	if err != nil {
		logging.Logger.Error("Error updating person with ID", zap.Uint("ID", uint(request.ID)), zap.Error(err))

		fault := createSOAPFault("soap:Server", "Internal Server Error", "500", "An unexpected error occurred.")
		c.XML(http.StatusInternalServerError, fault)
		return
	}

	logging.Logger.Info("Successfully updated person with ID", zap.Uint("ID", uint(request.ID)))

	// Формируем успешный ответ
	response := models.UpdatePersonResponse{
		Status: true,
	}

	// Возвращаем успешный ответ в формате XML
	c.XML(http.StatusOK, response)
}

func (h *StorageHandler) getPersonHandler(c *gin.Context, request *models.GetPersonRequest) {
	// Получаем информацию о человеке по ID
	fmt.Println("DDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDSWSSASADAAAAAAAAAAAAA", request.ID)
	person, err := h.Storage.PersonRepository.GetPerson(request.ID)
	if err != nil {
		logging.Logger.Error("Error getting person with ID", zap.Uint("ID", uint(request.ID)), zap.Error(err))

		// Формируем SOAP Fault для ошибки получения
		fault := createSOAPFault("soap:Server", "Internal Server Error", "500", "An unexpected error occurred.")
		c.XML(http.StatusInternalServerError, fault)
		return
	}

	// Если человек не найден, формируем SOAP Fault для клиента
	if person == nil {
		fmt.Printf("No person found with ID %d\n", request.ID)

		fault := createSOAPFault("soap:Client", models.ErrorRecordNotFoundMessage, models.ErrorRecordNotFoundCode, models.ErrorRecordNotFoundDetail)
		c.XML(http.StatusNotFound, fault)
		return
	}

	// Если человек найден, формируем ответ
	response := models.GetPersonResponse{
		Person: *person, // Разыменовываем указатель на структуру Person
	}

	// Возвращаем успешный ответ в формате XML
	c.XML(http.StatusOK, response)
}

func (h *StorageHandler) getAllPersonsHandler(c *gin.Context) {
	// Получаем всех людей из репозитория
	persons, err := h.Storage.PersonRepository.GetAllPersons()
	if err != nil {
		logging.Logger.Error("Error getting all persons", zap.Error(err))

		// Формируем SOAP Fault для ошибки получения
		fault := createSOAPFault("soap:Server", "Internal Server Error", "500", "An unexpected error occurred.")
		c.XML(http.StatusInternalServerError, fault)
		return
	}

	// Если людей не найдено, формируем SOAP Fault для клиента
	if len(persons) == 0 {
		fmt.Println("No persons found.")

		fault := createSOAPFault("soap:Client", models.ErrorRecordNotFoundMessage, models.ErrorRecordNotFoundCode, models.ErrorRecordNotFoundDetail)
		c.XML(http.StatusNotFound, fault)
		return
	}

	// Формируем ответ в формате SOAP
	response := models.GetAllPersonsResponse{
		Persons: persons,
	}

	// Возвращаем успешный ответ в формате XML
	c.XML(http.StatusOK, response)
}

func (h *StorageHandler) deletePersonHandler(c *gin.Context, request *models.DeletePersonRequest) {
	checkByID, err := h.Storage.PersonRepository.CheckPersonByID(uint(request.ID))
	if err != nil {
		logging.Logger.Error("Error getting person with ID", zap.Uint("ID", uint(request.ID)), zap.Error(err))

		fault := createSOAPFault("soap:Server", "Internal Server Error", "500", "An unexpected error occurred.")
		c.XML(http.StatusInternalServerError, fault)
		return
	} else if !checkByID {
		fault := createSOAPFault("soap:Client", models.ErrorRecordNotFoundMessage, models.ErrorRecordNotFoundCode, models.ErrorRecordNotFoundDetail)
		c.XML(http.StatusNotFound, fault)
		return
	}

	err = h.Storage.PersonRepository.DeletePerson(request)
	if err != nil {
		logging.Logger.Error("Error deleting person with ID", zap.Uint("ID", uint(request.ID)), zap.Error(err))

		fault := createSOAPFault("soap:Server", "Internal Server Error", "500", "An unexpected error occurred.")
		c.XML(http.StatusInternalServerError, fault)
		return
	}

	logging.Logger.Info("Successfully deleted person with ID", zap.Uint("ID", uint(request.ID)))
	c.XML(http.StatusOK, models.DeleteResponse{Status: true}) // Возвращаем успешный ответ
}

func (h *StorageHandler) searchPersonHandler(c *gin.Context, request *models.SearchPersonRequest) {
	persons, err := h.Storage.PersonRepository.SearchPerson(request.Query)
	if err != nil {
		logging.Logger.Error("Error searching for persons with query", zap.String("query", request.Query), zap.Error(err))

		fault := createSOAPFault("soap:Server", "Internal Server Error", "500", "An unexpected error occurred.")
		c.XML(http.StatusInternalServerError, fault)
		return
	}

	if len(persons) == 0 {
		fmt.Println("No persons found.")

		fault := createSOAPFault("soap:Client", models.ErrorRecordNotFoundMessage, models.ErrorRecordNotFoundCode, models.ErrorRecordNotFoundDetail)
		c.XML(http.StatusNotFound, fault)
		return
	} else {
		fmt.Printf("Found persons: %+v\n", persons)
	}

	// Формируем ответ в формате SOAP
	response := models.SearchPersonResponse{
		Persons: persons,
	}

	c.XML(http.StatusOK, response)
}
