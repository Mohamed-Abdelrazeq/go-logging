package services

import (
	"MohamedAbdelrazeq/go-logging/db"
	"MohamedAbdelrazeq/go-logging/models"
)

type Logger interface {
	CreateLogRecord(logRecord models.LogRecord) (models.LogRecord, error)
}

type loggerService struct {
	db db.Database
}

func NewLoggerService(db db.Database) Logger {
	return &loggerService{db}
}

func (service loggerService) CreateLogRecord(logRecord models.LogRecord) (models.LogRecord, error) {
	_, err := service.db.Exec("INSERT INTO log_records (level, message, timestamp) VALUES (?, ?, ?)", logRecord.Level, logRecord.Message, logRecord.Timestamp)
	if err != nil {
		return models.LogRecord{}, err
	}
	return logRecord, nil
}
