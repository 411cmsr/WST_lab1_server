package transport

import (
	"WST_lab1_server/config"
	"WST_lab1_server/internal/handlers"
	"fmt"
	"github.com/globusdigital/soap"
	"log"
	"net/http"

	"WST_lab1_server/internal/services"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Incoming request: %s %s\n", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
		log.Println("Request processed.")
	})
}

func NewSOAPServer(configFile string) *soap.Server {
	config, err := config.LoadConfig(configFile)

	if err != nil {
		log.Fatalf("error loading config: %v", err)
	}
	fmt.Println(config.Soap.PathHttp, "\n")
	fmt.Println(config.Soap.PathSoap, "\n")
	soapServer := soap.NewServer()
	soapServer.RegisterHandler(config.Soap.PathHttp+config.Soap.PathSoap, "addPersonRequest", "AddPerson", services.AddPersonRequestFactory, handlers.AddPersonHandler)
	soapServer.RegisterHandler(config.Soap.PathHttp+config.Soap.PathSoap, "updatePersonRequest", "UpdatePerson", services.UpdatePersonRequestFactory, handlers.UpdatePersonHandler)
	soapServer.RegisterHandler(config.Soap.PathHttp+config.Soap.PathSoap, "deletePersonRequest", "DeletePerson", services.DeletePersonRequestFactory, handlers.DeletePersonHandler)
	soapServer.RegisterHandler(config.Soap.PathHttp+config.Soap.PathSoap, "getPersonRequest", "GetPerson", services.GetPersonRequestFactory, handlers.GetPersonHandler)
	soapServer.RegisterHandler(config.Soap.PathHttp+config.Soap.PathSoap, "getAllPersonsRequest", "GetAllPersons", services.GetAllPersonsRequestFactory, handlers.GetAllPersonsHandler)
	soapServer.RegisterHandler(config.Soap.PathHttp+config.Soap.PathSoap, "searchPersonRequest", "SearchPerson", services.SearchPersonRequestFactory, handlers.SearchPersonHandler)

	return soapServer
}
