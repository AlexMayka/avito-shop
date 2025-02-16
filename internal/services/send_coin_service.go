package services

import (
	"avito_shop/internal/repositories"
	"avito_shop/internal/utils"
	"database/sql"
)

// SendCoinService осуществляет перевод монет от одного пользователя другому.
//
// Параметры:
//   - db: подключение к базе данных.
//   - idUserFrom: идентификатор пользователя, отправляющего монеты.
//   - nameUserTo: имя пользователя, получающего монеты.
//   - amount: количество монет для перевода.
//
// Возвращает:
//   - error: ошибку, если операция не удалась, или nil при успешном выполнении.
func SendCoinService(db *sql.DB, idUserFrom uint, nameUserTo string, amount int) error {
	tx, err := db.Begin()
	if err != nil {
		return utils.ErrInternalServerError
	}

	userTo, err := repositories.DeductionBalanceByUsername(tx, nameUserTo, amount)
	if err != nil {
		_ = tx.Rollback()
		return utils.ErrUserNotFound
	}

	userFrom, err := repositories.DeductionBalanceByIndex(tx, idUserFrom, -amount)
	if err != nil {
		_ = tx.Rollback()
		return utils.ErrNotEnoughCoins
	}

	_, err = repositories.CreateCoinTransfer(tx, userFrom.ID, userTo.ID, amount)
	if err != nil {
		_ = tx.Rollback()
		return utils.ErrInternalServerError
	}

	err = tx.Commit()
	if err != nil {
		return utils.ErrInternalServerError
	}

	return nil
}
