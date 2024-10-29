package services

import (
	"WST_lab1_server/handlers"
	"WST_lab1_server/models"
	"github.com/globusdigital/soap"
	"gorm.io/gorm"
	"log"
)

func RunSoapServer() *soap.Server {
	soapServer := soap.NewServer()
	log.Println("Server is successfully launched")
	soapServer.Log = log.Println
	soapServer.RegisterHandler("ProcessRequest", func(req models.Request)) (models.Response, error) {
		return handlers.ProcessRequest(req)
	})
	return soapServer
}
