// Package models содержит определения моделей данных.
package models

// User представляет пользователя системы.
type User struct {
	ID       uint   // Идентификатор пользователя.
	Username string // Имя пользователя.
	Password string // Пароль пользователя (хэшированный).
	Balance  int    // Баланс пользователя.
}
