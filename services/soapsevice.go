package services

import (
	"WST_lab1_server/models"
	"fmt"
	"gorm.io/gorm"
	"log"
	"net/http"

	"github.com/globusdigital/soap"
)

// RunServer run a little demo server
func RunSoapServer(db *gorm.DB) {
	soapServer := soap.NewServer()
	soapServer.Log = log.Println
	soapServer.RegisterHandler(
		"/soap",
		"operationSearch", // SOAPAction
		"searchRequest",   // tagname of soap body content
		func() interface{} {
			return &SearchRequest{}
		},
		// OperationHandlerFunc - do something
		func(request interface{}, w http.ResponseWriter, httpRequest *http.Request) (response interface{}, err error) {
			searchRequest := request.(*SearchRequest)
			var results []models.Persons
			if err := db.Where("name LIKE ? OR surname LIKE ?", "%"+searchRequest.Query+"%", "%"+searchRequest.Query+"%").Find(&results).Error; err != nil {
				return nil, err // Возвращаем ошибку, если поиск не удался
			}

			searchResponse := &SearchResponse{
				Results: results,
			}
			response = searchResponse
			return
		},
	)
	err := http.ListenAndServe(":8082", soapServer)
	fmt.Println("exiting with error", err)
}
