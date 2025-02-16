// Package models содержит определения моделей данных.
package models

// CoinTransferHistory представляет запись о переводе монет между пользователями.
type CoinTransferHistory struct {
	Username string `json:"username"`
	Amount   int    `json:"amount"`
}
