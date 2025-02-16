// Package db предоставляет функциональность для инициализации подключения к базе данных PostgreSQL.
package db

import (
	"database/sql"
	"fmt"

	// Импорт драйвера PostgreSQL.
	_ "github.com/lib/pq"
)

// InitDB устанавливает соединение с базой данных PostgreSQL.
// Параметры:
//   - host: адрес хоста базы данных.
//   - port: порт, на котором работает база данных.
//   - user: имя пользователя для подключения к базе данных.
//   - password: пароль для подключения.
//   - dbName: имя базы данных.
//
// Функция возвращает:
//   - *sql.DB: объект подключения к базе данных, если соединение успешно установлено.
//   - error: ошибку, если произошла неудача при открытии или пинге базы данных.
func InitDB(host, port, user, password, dbName string) (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbName,
	)

	// Открытие подключения к базе данных.
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("sql.Open error: %w", err)
	}

	// Проверка соединения с базой данных посредством отправки запроса Ping.
	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("db.Ping error: %w", err)
	}

	// Вывод сообщения об успешном подключении.
	fmt.Println("Connected to PostgreSQL")
	return db, nil
}
