// Package models содержит определения моделей данных.
package models

// UserPurchase описывает покупку пользователя.
type UserPurchase struct {
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}
