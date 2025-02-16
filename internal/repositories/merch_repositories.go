// Package repositories содержит функции для работы с базой данных.
package repositories

import (
	"avito_shop/internal/models"
	"database/sql"
	"errors"
)

// GetMerchByName ищет мерч по его названию.
//
// Параметры:
//   - q: интерфейс для выполнения SQL-запросов.
//   - itemName: название товара для поиска.
//
// Возвращает:
//   - указатель на найденный товар (models.Merch) или nil, если товар не найден,
//   - ошибку, если произошла ошибка при выполнении запроса.
func GetMerchByName(q Querier, itemName string) (*models.Merch, error) {
	query := `
        SELECT id, name, price
        FROM merch
        WHERE name = $1
        LIMIT 1
    `
	var merch models.Merch
	err := q.QueryRow(query, itemName).Scan(&merch.ID, &merch.Name, &merch.Price)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &merch, nil
}
