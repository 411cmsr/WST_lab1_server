package transport

import (
	"WST_lab1_server/config"
	"WST_lab1_server/internal/handlers"

	"WST_lab1_server/internal/services"
	//"log"

	"github.com/globusdigital/soap"
)

func NewSOAPServer() *soap.Server {
	soapServer := soap.NewServer()
	soapServer.RegisterHandler(config.HTTPServerSetting.PatHTTP+config.HTTPServerSetting.PathSoap,
		"Request", "AddPerson", services.AddPersonRequestFactory, handlers.AddPersonHandler)
	soapServer.RegisterHandler(config.HTTPServerSetting.PatHTTP+config.HTTPServerSetting.PathSoap,
		"Request", "UpdatePerson", services.UpdatePersonRequestFactory, handlers.UpdatePersonHandler)
	soapServer.RegisterHandler(config.HTTPServerSetting.PatHTTP+config.HTTPServerSetting.PathSoap,
		"Request", "DeletePerson", services.DeletePersonRequestFactory, handlers.DeletePersonHandler)
	soapServer.RegisterHandler(config.HTTPServerSetting.PatHTTP+config.HTTPServerSetting.PathSoap,
		"Request", "GetPerson", services.GetPersonRequestFactory, handlers.GetPersonHandler)
	soapServer.RegisterHandler(config.HTTPServerSetting.PatHTTP+config.HTTPServerSetting.PathSoap,
		"Request", "GetAllPersons", services.GetAllPersonsRequestFactory, handlers.GetAllPersonsHandler)
	soapServer.RegisterHandler(config.HTTPServerSetting.PatHTTP+config.HTTPServerSetting.PathSoap,
		"Request", "SearchPerson", services.SearchPersonRequestFactory, handlers.SearchPersonHandler)
	return soapServer
}
