package controllers

import (
	"avito_shop/internal/dto"
	"avito_shop/internal/services"
	"avito_shop/internal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// BuyMerchHandler обрабатывает запрос на покупку мерча.
//
// Параметры:
//   - c: контекст запроса Gin.
//
// Функция выполняет следующие действия:
//  1. Извлекает идентификатор пользователя из контекста.
//  2. Получает название товара из параметров URL.
//  3. Вызывает сервис BuyMerchService для обработки покупки мерча.
//  4. Формирует и возвращает JSON-ответ с информацией о покупке (сообщение, название товара, цена и обновлённый баланс).
func (ctr *Controller) BuyMerchHandler(c *gin.Context) {
	userId, _ := c.Get("id")
	userIdUint, ok := userId.(uint)
	if !ok {
		utils.HandleError(c, utils.ErrInvalidRequest)
		return
	}

	itemName := c.Param("item")
	if itemName == "" {
		utils.HandleError(c, utils.ErrInvalidRequest)
		return
	}

	userModel, merchModel, err := services.BuyMerchService(ctr.DB, userIdUint, itemName)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	response := dto.ByMerchResponse{
		Message: "Purchase successful",
		Item:    merchModel.Name,
		Price:   merchModel.Price,
		Balance: userModel.Balance,
	}

	c.JSON(http.StatusOK, response)
}
