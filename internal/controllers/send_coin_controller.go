package controllers

import (
	"avito_shop/internal/dto"
	"avito_shop/internal/services"
	"avito_shop/internal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// SendCoinController обрабатывает запрос на перевод монет между пользователями.
//
// Параметры:
//   - c: контекст запроса Gin.
//
// Функция выполняет следующие действия:
//  1. Извлекает идентификатор и имя пользователя из контекста.
//  2. Привязывает JSON-полезную нагрузку к структуре SendCoinRequest.
//  3. Проверяет, чтобы пользователь не переводил монеты самому себе.
//  4. Вызывает сервис SendCoinService для выполнения перевода.
//  5. Возвращает JSON-ответ с сообщением об успешном переводе.
func (ctr *Controller) SendCoinController(c *gin.Context) {
	userId, _ := c.Get("id")
	userName, _ := c.Get("username")

	userIdUint, ok := userId.(uint)
	if !ok {
		utils.HandleError(c, utils.ErrInvalidRequest)
		return
	}

	var req dto.SendCoinRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.HandleError(c, utils.ErrInvalidRequest)
		return
	}

	if userName == req.ToUser {
		utils.HandleError(c, utils.ErrSelfTransfer)
		return
	}

	if err := services.SendCoinService(ctr.DB, userIdUint, req.ToUser, req.Amount); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, dto.SendCoinResponse{Message: "Coin transfer successful"})
}
