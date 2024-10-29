package main

import (
	"MohamedAbdelrazeq/go-logging/db"
	handlers "MohamedAbdelrazeq/go-logging/handers"
	"MohamedAbdelrazeq/go-logging/services"
	"log"
	"net/http"
)

func main() {

	// init db
	connectionString := "file:logger-db.sqlite?cache=shared&mode=rwc"
	db, err := db.CreateAndConnectDB(connectionString)
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}
	defer db.Close()

	// init services
	loggerService := services.NewLoggerService(db)

	// init handlers
	loggerHandler := handlers.NewLoggerHandler(loggerService)

	http.HandleFunc("/health", handlers.HandleHealth)
	http.HandleFunc("/create-record", loggerHandler.CreateLogRecord)

	log.Println("Starting server on :8080...")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
