package handlers

import (
	"MohamedAbdelrazeq/go-logging/models"
	"MohamedAbdelrazeq/go-logging/services"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Logger interface {
	CreateLogRecord(w http.ResponseWriter, r *http.Request)
}

type loggerHandler struct {
	service services.Logger
}

func NewLoggerHandler(service services.Logger) Logger {
	return loggerHandler{service}
}

func (handler loggerHandler) CreateLogRecord(w http.ResponseWriter, r *http.Request) {
	// only allow POST requests
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// read the request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// validate the request body complies with the record struct
	var record models.LogRecord
	if err := json.Unmarshal(body, &record); err != nil {
		log.Println("Unable to open log fil: ", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// create a log record into db
	record, err = handler.service.CreateLogRecord(record)
	if err != nil {
		log.Println("Unable to create log record: ", err)
		http.Error(w, "Unable to create log record", http.StatusInternalServerError)
		return
	}

	// return the log record
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(record)
}
