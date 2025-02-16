// Package repositories содержит функции для работы с базой данных.
package repositories

import (
	"avito_shop/internal/models"
	"database/sql"
	"errors"
)

// CreatePurchases создает запись о покупке мерча пользователем.
//
// Параметры:
//   - q: интерфейс для выполнения SQL-запросов.
//   - userId: идентификатор пользователя, совершившего покупку.
//   - merchId: идентификатор товара, купленного пользователем.
//   - priceBought: цена покупки товара.
//
// Возвращает:
//   - указатель на созданную запись Purchases.
//   - ошибку, если операция не удалась.
func CreatePurchases(q Querier, userId, merchId uint, priceBought int) (*models.Purchases, error) {
	query := `
		INSERT INTO purchases (user_id, merch_id, price_bought)
		VALUES ($1, $2, $3)
		RETURNING id, merch_id, price_bought;
	`

	var purchases models.Purchases

	err := q.QueryRow(query, userId, merchId, priceBought).Scan(&purchases.ID, &purchases.MerchId, &purchases.PriceBought)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &purchases, nil
}

// GetUserPurchases возвращает список покупок пользователя.
//
// Параметры:
//   - q: интерфейс для выполнения SQL-запросов.
//   - userID: идентификатор пользователя.
//
// Возвращает:
//   - срез структур UserPurchase с информацией о купленных товарах.
//   - ошибку, если запрос не выполнен успешно.
func GetUserPurchases(q Querier, userID uint) ([]models.UserPurchase, error) {
	query := `
        SELECT m.name, p.co 
        FROM (
            SELECT p.merch_id, count(*) co 
            FROM purchases p 
            WHERE p.user_id = $1
            GROUP BY merch_id
        ) p
        INNER JOIN merch m 
        ON m.id = p.merch_id;
    `

	rows, err := q.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var purchases []models.UserPurchase
	for rows.Next() {
		var purchase models.UserPurchase
		if err := rows.Scan(&purchase.Name, &purchase.Quantity); err != nil {
			return nil, err
		}
		purchases = append(purchases, purchase)
	}

	return purchases, nil
}
