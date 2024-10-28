package main

import (
	handlers "MohamedAbdelrazeq/go-payment/handers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/health", handlers.HandleHealth)
	http.HandleFunc("/create-payment-intent", handlers.HandleCreatePaymentIntent)

	log.Println("Starting server on :8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
