package services

import (
	"avito_shop/internal/models"
	"avito_shop/internal/repositories"
	"avito_shop/internal/utils"
	"database/sql"
)

// InfoService возвращает информацию о пользователе, его покупках,
// а также историю полученных и отправленных переводов монет.
//
// Параметры:
//   - db: подключение к базе данных.
//   - userId: идентификатор пользователя.
//
// Возвращает:
//   - *models.User: данные пользователя.
//   - []models.UserPurchase: список покупок пользователя.
//   - []models.CoinTransferHistory: историю полученных монет.
//   - []models.CoinTransferHistory: историю отправленных монет.
//   - error: ошибку, если произошла неудача.
func InfoService(db *sql.DB, userId uint) (*models.User, []models.UserPurchase, []models.CoinTransferHistory, []models.CoinTransferHistory, error) {

	user, err := repositories.GetUserById(db, userId)
	if err != nil {
		return nil, nil, nil, nil, utils.ErrInternalServerError
	}

	purchases, err := repositories.GetUserPurchases(db, userId)
	if err != nil {
		return nil, nil, nil, nil, utils.ErrUserNotFound
	}

	receivedCoins, err := repositories.GetReceivedCoins(db, userId)
	if err != nil {
		return nil, nil, nil, nil, utils.ErrInternalServerError
	}

	sentCoins, err := repositories.GetSentCoins(db, userId)
	if err != nil {
		return nil, nil, nil, nil, utils.ErrInternalServerError
	}

	return user, purchases, receivedCoins, sentCoins, nil
}
