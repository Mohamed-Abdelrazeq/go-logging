package handlers

import (
	"MohamedAbdelrazeq/go-logging/models"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

func CreateLogRecord(w http.ResponseWriter, r *http.Request) {
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
	file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		http.Error(w, "Unable to open log file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	if _, err := file.Write([]byte(record.Message + " " + record.Timestamp.GoString())); err != nil {
		http.Error(w, "Unable to write log record to file", http.StatusInternalServerError)
		return
	}

	if _, err := file.WriteString("\n"); err != nil {
		http.Error(w, "Unable to write newline to file", http.StatusInternalServerError)
		return
	}
	// return the log record
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(record)
}
