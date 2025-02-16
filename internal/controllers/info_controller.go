package controllers

import (
	"avito_shop/internal/dto"
	"avito_shop/internal/services"
	"avito_shop/internal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// InfoHandler обрабатывает запрос на получение информации о пользователе.
// Функция извлекает идентификатор пользователя из контекста, вызывает сервис InfoService для получения
// данных о балансе, покупках и истории переводов, затем формирует и отправляет JSON-ответ.
//
// Параметры:
//   - c: контекст запроса Gin.
func (ctr *Controller) InfoHandler(c *gin.Context) {
	userId, _ := c.Get("id")
	userIdUint, ok := userId.(uint)
	if !ok {
		utils.HandleError(c, utils.ErrInvalidRequest)
		return
	}

	user, purchases, receivedCoins, sentCoins, err := services.InfoService(ctr.DB, userIdUint)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	var inventory []dto.UserPurchase
	for _, p := range purchases {
		inventory = append(inventory, dto.UserPurchase{
			Type:     p.Name,
			Quantity: p.Quantity,
		})
	}

	var received []dto.CoinTransferHistory
	for _, r := range receivedCoins {
		received = append(received, dto.CoinTransferHistory{
			Username: r.Username,
			Amount:   r.Amount,
		})
	}

	var sent []dto.CoinTransferHistory
	for _, s := range sentCoins {
		sent = append(sent, dto.CoinTransferHistory{
			Username: s.Username,
			Amount:   s.Amount,
		})
	}

	response := dto.InfoResponse{
		Balance:   user.Balance,
		Inventory: inventory,
		CoinHistory: dto.CoinHistory{
			Received: received,
			Sent:     sent,
		},
	}

	c.JSON(http.StatusOK, response)
}
