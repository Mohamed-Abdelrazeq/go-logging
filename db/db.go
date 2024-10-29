package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Database interface {
	Connect(connectionString string) error
	Close() error
	Query(query string, args ...interface{}) (Rows, error)
	Exec(query string, args ...interface{}) (Result, error)
}

type Rows interface {
	Next() bool
	Scan(dest ...interface{}) error
	Close() error
}

type Result interface {
	LastInsertId() (int64, error)
	RowsAffected() (int64, error)
}

type DB struct {
	db *sql.DB
}

func (s *DB) Connect(connectionString string) error {
	var err error
	s.db, err = sql.Open("sqlite3", connectionString)
	if err != nil {
		return err
	}
	return s.db.Ping()
}

func (s *DB) Close() error {
	return s.db.Close()
}

func (s *DB) Query(query string, args ...interface{}) (Rows, error) {
	rows, err := s.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	return &SQLiteRows{rows: rows}, nil
}

func (s *DB) Exec(query string, args ...interface{}) (Result, error) {
	result, err := s.db.Exec(query, args...)
	if err != nil {
		return nil, err
	}
	return &SQLiteResult{result: result}, nil
}

type SQLiteRows struct {
	rows *sql.Rows
}

func (r *SQLiteRows) Next() bool {
	return r.rows.Next()
}

func (r *SQLiteRows) Scan(dest ...interface{}) error {
	return r.rows.Scan(dest...)
}

func (r *SQLiteRows) Close() error {
	return r.rows.Close()
}

type SQLiteResult struct {
	result sql.Result
}

func (r *SQLiteResult) LastInsertId() (int64, error) {
	return r.result.LastInsertId()
}

func (r *SQLiteResult) RowsAffected() (int64, error) {
	return r.result.RowsAffected()
}

func CreateAndConnectDB(connectionString string) (Database, error) {
	db := &DB{}
	err := db.Connect(connectionString)
	if err != nil {
		return nil, err
	}
	return db, nil
}
