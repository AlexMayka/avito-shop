// Package models содержит определения моделей данных.
package models

// Merch представляет товар, доступный для покупки.
// Содержит идентификатор товара, его название и цену.
type Merch struct {
	ID    uint   // Идентификатор товара.
	Name  string // Название товара.
	Price int    // Цена товара.
}
