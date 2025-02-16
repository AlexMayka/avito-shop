// Package repositories содержит функции для работы с базой данных,
// связанные с переводом монет между пользователями.
package repositories

import (
	"avito_shop/internal/models"
	"database/sql"
	"errors"
)

// CreateCoinTransfer создаёт запись о переводе монет между пользователями.
//
// Параметры:
//   - q: интерфейс для выполнения SQL-запросов.
//   - fromUserId: идентификатор пользователя, отправляющего монеты.
//   - ToUserId: идентификатор пользователя, получающего монеты.
//   - amount: количество монет для перевода.
//
// Возвращает:
//   - указатель на созданную запись перевода монет.
//   - ошибку, если перевод не был создан.
func CreateCoinTransfer(q Querier, fromUserId, ToUserId uint, amount int) (*models.CoinTransfer, error) {
	query := `
		INSERT INTO coin_transfers (from_user_id, to_user_id, amount)
		VALUES ($1, $2, $3)
		RETURNING id, from_user_id, to_user_id, amount;
	`

	var coinTransfer models.CoinTransfer

	err := q.QueryRow(query, fromUserId, ToUserId, amount).Scan(
		&coinTransfer.ID, &coinTransfer.FromUserId, &coinTransfer.ToUserId, &coinTransfer.Amount,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &coinTransfer, nil
}

// GetSentCoins возвращает историю переводов монет, отправленных пользователем.
//
// Параметры:
//   - q: интерфейс для выполнения SQL-запросов.
//   - userID: идентификатор пользователя, отправившего монеты.
//
// Возвращает:
//   - срез структур CoinTransferHistory с информацией о переведённых монетах.
//   - ошибку, если запрос не выполнен успешно.
func GetSentCoins(q Querier, userID uint) ([]models.CoinTransferHistory, error) {
	query := `
        SELECT u.username, ct.am 
        FROM (
            SELECT ct.to_user_id, SUM(ct.amount) am 
            FROM coin_transfers ct
            WHERE from_user_id = $1
            GROUP BY ct.to_user_id
        ) ct
        INNER JOIN users u 
        ON ct.to_user_id = u.id;
    `

	rows, err := q.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sentCoins []models.CoinTransferHistory
	for rows.Next() {
		var history models.CoinTransferHistory
		if err := rows.Scan(&history.Username, &history.Amount); err != nil {
			return nil, err
		}
		sentCoins = append(sentCoins, history)
	}

	return sentCoins, nil
}

// GetReceivedCoins возвращает историю переводов монет, полученных пользователем.
//
// Параметры:
//   - q: интерфейс для выполнения SQL-запросов.
//   - userID: идентификатор пользователя, получившего монеты.
//
// Возвращает:
//   - срез структур CoinTransferHistory с информацией о полученных монетах.
//   - ошибку, если запрос не выполнен успешно.
func GetReceivedCoins(q Querier, userID uint) ([]models.CoinTransferHistory, error) {
	query := `
        SELECT u.username, ct.am 
        FROM (
            SELECT ct.from_user_id, SUM(ct.amount) am 
            FROM coin_transfers ct
            WHERE to_user_id = $1
            GROUP BY ct.from_user_id
        ) ct
        INNER JOIN users u 
        ON ct.from_user_id = u.id;
    `

	rows, err := q.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var receivedCoins []models.CoinTransferHistory
	for rows.Next() {
		var history models.CoinTransferHistory
		if err := rows.Scan(&history.Username, &history.Amount); err != nil {
			return nil, err
		}
		receivedCoins = append(receivedCoins, history)
	}

	return receivedCoins, nil
}
