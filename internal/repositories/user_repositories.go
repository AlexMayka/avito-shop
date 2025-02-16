// Package repositories содержит функции для работы с данными пользователей.
package repositories

import (
	"avito_shop/internal/models"
	"database/sql"
	"errors"
)

// GetUserByUsername возвращает пользователя по имени.
// Если пользователь не найден, возвращается (nil, nil).
func GetUserByUsername(q Querier, username string) (*models.User, error) {
	query := `
        SELECT id, username, password, balance
        FROM users
        WHERE username = $1
        LIMIT 1
    `
	var user models.User
	err := q.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.Password, &user.Balance)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// GetUserById возвращает пользователя по его идентификатору.
// Если пользователь не найден, возвращается (nil, nil).
func GetUserById(q Querier, id uint) (*models.User, error) {
	query := `
		SELECT id, username, password, balance
		FROM users
		WHERE id = $1
	`
	var user models.User
	err := q.QueryRow(query, id).Scan(&user.ID, &user.Username, &user.Password, &user.Balance)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// CreateUser создаёт нового пользователя с заданными данными.
// Возвращает созданного пользователя или ошибку.
func CreateUser(q Querier, username, hashedPassword string, initialBalance int64) (*models.User, error) {
	query := `
        INSERT INTO users (username, password, balance)
        VALUES ($1, $2, $3)
        RETURNING id, username, password, balance
    `
	var user models.User
	err := q.QueryRow(query, username, hashedPassword, initialBalance).
		Scan(&user.ID, &user.Username, &user.Password, &user.Balance)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// DeductionBalanceByIndex обновляет баланс пользователя по его идентификатору.
// При обновлении баланс увеличивается на указанную сумму (price).
// Если пользователь не найден, возвращается ошибка.
func DeductionBalanceByIndex(q Querier, id uint, price int) (*models.User, error) {
	query := `
		UPDATE users 
		SET balance = balance + $1 
		WHERE id = $2
		RETURNING id, username, password, balance;
	`
	var user models.User
	err := q.QueryRow(query, price, id).Scan(&user.ID, &user.Username, &user.Password, &user.Balance)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

// DeductionBalanceByUsername обновляет баланс пользователя по имени.
// При обновлении баланс увеличивается на указанную сумму (price).
// Если пользователь не найден, возвращается ошибка.
func DeductionBalanceByUsername(q Querier, username string, price int) (*models.User, error) {
	query := `
		UPDATE users 
		SET balance = balance + $1 
		WHERE username = $2
		RETURNING id, username, password, balance;
	`
	var user models.User
	err := q.QueryRow(query, price, username).Scan(&user.ID, &user.Username, &user.Password, &user.Balance)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}
