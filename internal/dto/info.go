// Package dto содержит определения структур данных (DTO),
// используемых для обмена информацией между клиентом и сервером.
package dto

// InfoResponse представляет ответ с информацией о пользователе,
// включая баланс монет, инвентарь и историю переводов монет.
type InfoResponse struct {
	Balance     int            `json:"coins"`
	Inventory   []UserPurchase `json:"inventory"`
	CoinHistory CoinHistory    `json:"coinHistory"`
}

// UserPurchase представляет информацию о покупке пользователя.
//
// Поле Type указывает тип или название покупки,
// а поле Quantity — количество единиц приобретённого товара.
type UserPurchase struct {
	Type     string `json:"type"`
	Quantity int    `json:"quantity"`
}

// CoinHistory содержит историю переводов монет, разделённую на полученные и отправленные.
type CoinHistory struct {
	Received []CoinTransferHistory `json:"received"`
	Sent     []CoinTransferHistory `json:"sent"`
}

// CoinTransferHistory представляет запись о переводе монет между пользователями.
//
// Поле Username указывает имя пользователя, участвовавшего в транзакции,
// а поле Amount — количество переведённых монет.
type CoinTransferHistory struct {
	Username string `json:"username"`
	Amount   int    `json:"amount"`
}
