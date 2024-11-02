package main

import (
	"MohamedAbdelrazeq/go-logging/db"
	handlers "MohamedAbdelrazeq/go-logging/handers"
	"MohamedAbdelrazeq/go-logging/services"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	// init env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// init db
	connectionString := os.Getenv("DB_CONNECTION")
	driver := os.Getenv("DRIVER")
	log.Println("Connecting to database: ", connectionString)
	db, err := db.CreateAndConnectDB(driver, connectionString)
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
	http.HandleFunc("/get-records", loggerHandler.GetLogRecords)
	http.HandleFunc("/get-record", loggerHandler.GetLogRecordsById)
	http.HandleFunc("/get-records-by-level", loggerHandler.GetLogRecordsByLevel)
	http.HandleFunc("/get-records-by-date-range", loggerHandler.GetLogRecordsByDateRange)

	log.Println("Starting server on :8080...")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
