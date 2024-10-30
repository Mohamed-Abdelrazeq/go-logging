package handlers

import (
	"MohamedAbdelrazeq/go-logging/models"
	"MohamedAbdelrazeq/go-logging/services"
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
)

type Logger interface {
	CreateLogRecord(w http.ResponseWriter, r *http.Request)
	GetLogRecords(w http.ResponseWriter, r *http.Request)
	GetLogRecordsById(w http.ResponseWriter, r *http.Request)
	GetLogRecordsByLevel(w http.ResponseWriter, r *http.Request)
	GetLogRecordsByDateRange(w http.ResponseWriter, r *http.Request)
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

	// replace \n with \\n
	body = bytes.ReplaceAll(body, []byte("\n"), []byte(""))
	body = bytes.ReplaceAll(body, []byte("\r"), []byte(""))

	// validate the request body complies with the record struct
	var record models.LogRecord
	if err := json.Unmarshal(body, &record); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
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

func (handler loggerHandler) GetLogRecords(w http.ResponseWriter, r *http.Request) {
	// only allow GET requests
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// get the page from the URL
	pageStr := r.URL.Query().Get("page")
	if pageStr == "" {
		pageStr = "0"
	}

	// convert the page to an integer
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		http.Error(w, "Invalid page number", http.StatusBadRequest)
		return
	}

	// get the log records from db
	records, err := handler.service.GetLogRecords(page)
	if err != nil {
		log.Println("Unable to get log records: ", err)
		http.Error(w, "Unable to get log records", http.StatusInternalServerError)
		return
	}

	// return the log records
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(records)
}

func (handler loggerHandler) GetLogRecordsById(w http.ResponseWriter, r *http.Request) {
	// only allow GET requests
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// get the log record ID from the URL
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "Missing log record ID", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid log record ID", http.StatusBadRequest)
		return
	}

	// get the log record by ID from db
	record, err := handler.service.GetLogRecordsById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// return the log record
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(record)
}

func (handler loggerHandler) GetLogRecordsByLevel(w http.ResponseWriter, r *http.Request) {
	// only allow GET requests
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// get the log record level from the URL
	level := r.URL.Query().Get("level")
	if level == "" {
		http.Error(w, "Missing log record level", http.StatusBadRequest)
		return
	}

	// get the page from the URL
	pageStr := r.URL.Query().Get("page")
	if pageStr == "" {
		pageStr = "0"
	}

	// convert the page to an integer
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		http.Error(w, "Invalid page number", http.StatusBadRequest)
		return
	}

	// get the log records by level from db
	records, err := handler.service.GetLogRecordsByLevel(level, page)
	if err != nil {
		log.Println("Unable to get log records: ", err)
		http.Error(w, "Unable to get log records", http.StatusInternalServerError)
		return
	}

	// return the log records
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(records)
}

func (handler loggerHandler) GetLogRecordsByDateRange(w http.ResponseWriter, r *http.Request) {
	// only allow GET requests
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// get the log record date range from the URL
	startDate := r.URL.Query().Get("start_date")
	if startDate == "" {
		http.Error(w, "Missing log record start date", http.StatusBadRequest)
		return
	}
	endDate := r.URL.Query().Get("end_date")
	if endDate == "" {
		http.Error(w, "Missing log record end date", http.StatusBadRequest)
		return
	}

	// get the log records by date range from db
	records, err := handler.service.GetLogRecordsByDateRange(startDate, endDate)
	if err != nil {
		log.Println("Unable to get log records: ", err)
		http.Error(w, "Unable to get log records", http.StatusInternalServerError)
		return
	}

	// return the log records
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(records)
}
