package main

import (
	handlers "MohamedAbdelrazeq/go-logging/handers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/health", handlers.HandleHealth)
	http.HandleFunc("/create-record", handlers.CreateLogRecord)

	log.Println("Starting server on :8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
