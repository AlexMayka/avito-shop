package unit_test

import (
	"database/sql"

	"github.com/DATA-DOG/go-sqlmock"
)

// MockQuerier реализует интерфейс Querier для мокирования запросов к базе данных.
type MockQuerier struct {
	db   *sql.DB
	mock sqlmock.Sqlmock
}

func (m *MockQuerier) Exec(query string, args ...interface{}) (sql.Result, error) {
	return m.db.Exec(query, args...)
}

func (m *MockQuerier) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return m.db.Query(query, args...)
}

func (m *MockQuerier) QueryRow(query string, args ...interface{}) *sql.Row {
	return m.db.QueryRow(query, args...)
}
