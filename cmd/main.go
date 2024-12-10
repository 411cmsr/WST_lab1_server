package main

import (
	"WST_lab1_server/config"
	"fmt"

	//"WST_lab1_server/internal/database"
	//"WST_lab1_server/internal/database/postgres"
	"WST_lab1_server/internal/transport"
	"encoding/xml"
	"log"
	"net/http"

	"go.uber.org/zap"
)

type Fault struct {
	XMLName     xml.Name `xml:"Fault"`
	FaultString string   `xml:"faultstring"`
}

var logger *zap.Logger

func init() {
	var err error
	logger, err = zap.NewProduction()
	if err != nil {
		log.Fatalf("Failed to create logger: %v", err)
	}
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Info("Incoming request",
			zap.String("method", r.Method),
			zap.String("url", r.URL.String()),
			zap.String("remote_addr", r.RemoteAddr),
		)

		rec := &responseRecorder{ResponseWriter: w}
		next.ServeHTTP(rec, r)

		if rec.isFault {
			logger.Warn("SOAP Fault occurred",
				zap.String("fault_string", rec.faultMessage),
				zap.String("method", r.Method),
				zap.String("url", r.URL.String()),
			)
		} else {
			logger.Info("Response sent",
				zap.String("method", r.Method),
				zap.String("url", r.URL.String()),
				zap.Int("status_code", rec.statusCode),
			)
		}
	})
}

type responseRecorder struct {
	http.ResponseWriter
	statusCode   int
	isFault      bool
	faultMessage string
}

func (r *responseRecorder) WriteHeader(code int) {
	r.statusCode = code
	r.ResponseWriter.WriteHeader(code)
}

func (r *responseRecorder) Write(b []byte) (int, error) {
	if isSOAPFault(b) {
		var envelope struct {
			Body struct {
				Fault Fault `xml:"Fault"`
			} `xml:"Body"`
		}
		if err := xml.Unmarshal(b, &envelope); err == nil {
			r.isFault = true
			r.faultMessage = envelope.Body.Fault.FaultString
		}
	}
	return r.ResponseWriter.Write(b)
}

func isSOAPFault(body []byte) bool {
	var envelope struct {
		Body struct {
			Fault Fault `xml:"Fault"`
		} `xml:"Body"`
	}
	err := xml.Unmarshal(body, &envelope)
	if err != nil || envelope.Body.Fault.FaultString == "" {
		return false
	}
	return true
}

func main() {
	fmt.Println("b\n\n\n")
	config.Init()
	fmt.Println("c\n\n\n")



	////
	//configFile := "config/config.yaml"

	//db := postgres.Init()
	// if err != nil {
	// 	logger.Fatal("Failed to connect to database", zap.Error(err))
	// }
	// logger.Info("Database connection established successfully.")

	// err = database.UpdateDB(configFile)
	// if err != nil {
	// 	logger.Fatal("Failed to update database", zap.Error(err))
	// }
	// logger.Info("Database updated successfully.")

	soapServer := transport.NewSOAPServer()
	var err error

	if err != nil {
		logger.Fatal("Failed to run SOAP Server", zap.Error(err))
	}

	http.Handle("/", loggingMiddleware(soapServer))

	logger.Info("Starting SOAP server on :8094")

	if err := http.ListenAndServe(":8094", nil); err != nil {
		logger.Fatal("Failed to start server", zap.Error(err))
	}

	if err := logger.Sync(); err != nil {
		log.Fatalf("Error syncing logger: %v", err)
	}
}
