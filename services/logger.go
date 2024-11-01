package services

import (
	"MohamedAbdelrazeq/go-logging/db"
	"MohamedAbdelrazeq/go-logging/models"
	"errors"
)

type Logger interface {
	CreateLogRecord(logRecord models.LogRecord) (models.LogRecord, error)
	GetLogRecords(page int) ([]models.LogRecord, error)
	GetLogRecordsById(id int) (models.LogRecord, error)
	GetLogRecordsByLevel(level string, page int) ([]models.LogRecord, error)
	GetLogRecordsByDateRange(startDate string, endDate string) ([]models.LogRecord, error)
}

type loggerService struct {
	db db.Database
}

func NewLoggerService(db db.Database) Logger {
	return &loggerService{db}
}

func (service loggerService) CreateLogRecord(logRecord models.LogRecord) (models.LogRecord, error) {
	results, err := service.db.Exec("INSERT INTO log_records (level, message, timestamp) VALUES (?, ?, ?)", logRecord.Level, logRecord.Message, logRecord.Timestamp)
	if err != nil {
		return models.LogRecord{}, err
	}
	tempId, err := results.LastInsertId()
	if err != nil {
		return models.LogRecord{}, err
	}
	logRecord.ID = int(tempId)
	return logRecord, nil
}

func (service loggerService) GetLogRecords(page int) ([]models.LogRecord, error) {
	rows, err := service.db.Query("SELECT id, level, message, timestamp FROM log_records LIMIT 10 OFFSET ?", page*10)
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
		return models.LogRecord{}, errors.New("record not found")
	}

	var logRecord models.LogRecord
	if row.Next() {
		err = row.Scan(&logRecord.ID, &logRecord.Level, &logRecord.Message, &logRecord.Timestamp)
		if err != nil {
			return models.LogRecord{}, errors.New("record not found")
		}
		return logRecord, nil
	}
	return models.LogRecord{}, errors.New("record not found")
}

func (service loggerService) GetLogRecordsByLevel(level string, page int) ([]models.LogRecord, error) {
	rows, err := service.db.Query("SELECT id, level, message, timestamp FROM log_records WHERE level = ? LIMIT 10 OFFSET ?", level, page*10)
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

func (service loggerService) GetLogRecordsByDateRange(startDate string, endDate string) ([]models.LogRecord, error) {
	rows, err := service.db.Query("SELECT id, level, message, timestamp FROM log_records WHERE timestamp BETWEEN ? AND ?", startDate, endDate)
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
