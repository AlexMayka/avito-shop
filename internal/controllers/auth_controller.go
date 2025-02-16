package controllers

import (
	"avito_shop/internal/dto"
	"avito_shop/internal/services"
	"avito_shop/internal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// AuthHandler обрабатывает запрос на аутентификацию пользователя.
//
// Параметры:
//   - c: контекст Gin, через который передается HTTP-запрос.
//
// Функция выполняет следующие действия:
//  1. Привязывает JSON-полезную нагрузку запроса к структуре AuthRequest.
//  2. При ошибке привязки возвращает ответ с ошибкой "invalid request".
//  3. Вызывает AuthService для аутентификации пользователя с использованием
//     username, password и секрета для подписи JWT.
//  4. При успешной аутентификации возвращает ответ с кодом 200 и JSON, содержащий JWT токен.
func (ctr *Controller) AuthHandler(c *gin.Context) {
	var req dto.AuthRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.HandleError(c, utils.ErrInvalidRequest)
		return
	}

	token, err := services.AuthService(ctr.DB, ctr.CFG.Host.JWTSecret, req.Username, req.Password)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, dto.AuthResponse{Token: token})
}
