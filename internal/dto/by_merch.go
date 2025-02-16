// Package dto содержит определения структур данных (DTO),
// используемых для обмена информацией между клиентом и сервером.
package dto

// ByMerchResponse представляет структуру ответа на запрос о покупке мерча.
// В данном ответе содержится информация о выполненной покупке, включая
// сообщение о результате операции, название купленного товара, его стоимость
// и обновлённый баланс пользователя.
type ByMerchResponse struct {
	Message string `json:"message"`
	Item    string `json:"item"`
	Price   int    `json:"price"`
	Balance int    `json:"balance"`
}
