// Package controllers содержит контроллеры для обработки HTTP-запросов.
package controllers

import (
	"avito_shop/config"
	"database/sql"
)

// Controller предоставляет базовые зависимости для контроллеров приложения.
// Он содержит подключение к базе данных и конфигурационные настройки.
type Controller struct {
	DB  *sql.DB        // Подключение к базе данных.
	CFG *config.Config // Конфигурация приложения.
}
