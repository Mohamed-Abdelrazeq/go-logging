package services

import (
	"MohamedAbdelrazeq/go-logging/db"
	"MohamedAbdelrazeq/go-logging/models"
)

type Logger interface {
	CreateLogRecord(logRecord models.LogRecord) (models.LogRecord, error)
	GetLogRecords() ([]models.LogRecord, error)
	GetLogRecordsById(id int) (models.LogRecord, error)
	GetLogRecordsByLevel(level string) ([]models.LogRecord, error)
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

func (service loggerService) GetLogRecords() ([]models.LogRecord, error) {
	rows, err := service.db.Query("SELECT id, level, message, timestamp FROM log_records")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var logRecords []models.LogRecord
	for rows.Next() {
		var logRecord models.LogRecord
		err := rows.Scan(&logRecord.ID, &logRecord.Level, &logRecord.Message, &logRecord.Timestamp)
		if err != nil {
			return nil, err
		}
		logRecords = append(logRecords, logRecord)
	}
	return logRecords, nil
}

func (service loggerService) GetLogRecordsById(id int) (models.LogRecord, error) {
	row, err := service.db.Query("SELECT id, level, message, timestamp FROM log_records WHERE id = ?", id)
	if err != nil {
		return models.LogRecord{}, err
	}

	var logRecord models.LogRecord
	err = row.Scan(&logRecord.ID, &logRecord.Level, &logRecord.Message)
	if err != nil {
		return models.LogRecord{}, err
	}
	return logRecord, nil
}

func (service loggerService) GetLogRecordsByLevel(level string) ([]models.LogRecord, error) {
	rows, err := service.db.Query("SELECT id, level, message, timestamp FROM log_records WHERE level = ?", level)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var logRecords []models.LogRecord
	for rows.Next() {
		var logRecord models.LogRecord
		err := rows.Scan(&logRecord.ID, &logRecord.Level, &logRecord.Message, &logRecord.Timestamp)
		if err != nil {
			return nil, err
		}
		logRecords = append(logRecords, logRecord)
	}
	return logRecords, nil
}
