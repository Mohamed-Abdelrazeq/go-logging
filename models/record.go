package models

import "time"

type Level int

// The record have 3 levels
// level 1 is for info
// level 2 is for warning
// level 3 is for errors

const (
	Info Level = iota + 1
	Warning
	Error
)

type LogRecord struct {
	ID        int       `json:"id"`
	Level     Level     `json:"level"`
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
}
