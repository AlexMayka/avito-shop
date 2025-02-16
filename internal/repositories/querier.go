// Package repositories содержит функции для работы с базой данных и определения интерфейса для выполнения SQL-запросов.
package repositories

import "database/sql"

// Querier абстрагирует выполнение SQL-запросов, позволяя использовать как *sql.DB, так и *sql.Tx.
type Querier interface {
	// Exec выполняет SQL-запрос, не возвращающий строки (например, INSERT, UPDATE, DELETE).
	Exec(query string, args ...interface{}) (sql.Result, error)
	// Query выполняет SQL-запрос, возвращающий набор строк.
	Query(query string, args ...interface{}) (*sql.Rows, error)
	// QueryRow выполняет SQL-запрос, возвращающий одну строку.
	QueryRow(query string, args ...interface{}) *sql.Row
}
