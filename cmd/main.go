package main

import (
	"WST_lab1_server/internal/database"
	"WST_lab1_server/internal/transport"
	"log"
	"net/http"
)

func main() {
	configFile := "config/config.yaml"

	err := database.InitDB(configFile)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	log.Println("Database connection established successfully.")

	err = database.UpdateDB(configFile)
	if err != nil {
		log.Fatalf("Failed to update database: %v", err)
	}
	log.Println("Database updated successfully.")

	soapServer := transport.NewSOAPServer(configFile)
	if err != nil {
		log.Fatalf("Failed to run SOAP Server: %v", err)
	}

	http.Handle("/", soapServer)
	log.Println("Starting SOAP server on :8094")
	if err := http.ListenAndServe(":8094", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
