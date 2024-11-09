package transport

import (
	"WST_lab1_server/config"
	"WST_lab1_server/internal/handlers"

	"WST_lab1_server/internal/services"
	"github.com/globusdigital/soap"
	"log"
)

func NewSOAPServer(configFile string) *soap.Server {
	configuration, err := config.LoadConfig(configFile)

	if err != nil {
		log.Fatalf("error loading config: %v", err)
	}
	soapServer := soap.NewServer()
	soapServer.RegisterHandler(configuration.Soap.PathHttp+configuration.Soap.PathSoap, "Request", "AddPerson", services.AddPersonRequestFactory, handlers.AddPersonHandler)
	soapServer.RegisterHandler(configuration.Soap.PathHttp+configuration.Soap.PathSoap, "Request", "UpdatePerson", services.UpdatePersonRequestFactory, handlers.UpdatePersonHandler)
	soapServer.RegisterHandler(configuration.Soap.PathHttp+configuration.Soap.PathSoap, "Request", "DeletePerson", services.DeletePersonRequestFactory, handlers.DeletePersonHandler)
	soapServer.RegisterHandler(configuration.Soap.PathHttp+configuration.Soap.PathSoap, "Request", "GetPerson", services.GetPersonRequestFactory, handlers.GetPersonHandler)
	soapServer.RegisterHandler(configuration.Soap.PathHttp+configuration.Soap.PathSoap, "Request", "GetAllPersons", services.GetAllPersonsRequestFactory, handlers.GetAllPersonsHandler)
	soapServer.RegisterHandler(configuration.Soap.PathHttp+configuration.Soap.PathSoap, "Request", "SearchPerson", services.SearchPersonRequestFactory, handlers.SearchPersonHandler)

	return soapServer
}
