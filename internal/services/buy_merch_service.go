package services

import (
	"avito_shop/internal/models"
	"avito_shop/internal/repositories"
	"avito_shop/internal/utils"
	"database/sql"
)

// BuyMerchService осуществляет покупку мерча пользователем.
//
// Параметры:
//   - db: подключение к базе данных.
//   - userId: идентификатор пользователя, совершающего покупку.
//   - itemName: название товара для покупки.
//
// Функция выполняет следующие действия:
//  1. Начинает транзакцию.
//  2. Ищет товар по названию. Если товар не найден, откатывает транзакцию и возвращает ошибку.
//  3. Пытается списать с баланса пользователя сумму, равную стоимости товара.
//     Если средств недостаточно, откатывает транзакцию и возвращает ошибку.
//  4. Создает запись о покупке товара.
//  5. Фиксирует транзакцию и возвращает обновленного пользователя и товар.
func BuyMerchService(db *sql.DB, userId uint, itemName string) (*models.User, *models.Merch, error) {
	tx, err := db.Begin()
	if err != nil {
		return nil, nil, utils.ErrInternalServerError
	}

	merch, err := repositories.GetMerchByName(tx, itemName)
	if err != nil {
		_ = tx.Rollback()
		return nil, nil, utils.ErrInternalServerError
	}
	if merch == nil {
		_ = tx.Rollback()
		return nil, nil, utils.ErrItemNotFound
	}

	user, err := repositories.DeductionBalanceByIndex(tx, userId, -merch.Price)
	if err != nil {
		_ = tx.Rollback()
		return nil, nil, utils.ErrNotEnoughCoins
	}

	_, err = repositories.CreatePurchases(tx, user.ID, merch.ID, merch.Price)
	if err != nil {
		_ = tx.Rollback()
		return nil, nil, utils.ErrInternalServerError
	}

	err = tx.Commit()
	if err != nil {
		return nil, nil, utils.ErrInternalServerError
	}

	return user, merch, nil
}
