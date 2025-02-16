// Package utils предоставляет утилитарные функции для обработки ошибок и формирования корректных HTTP-ответов.
package utils

import (
	"avito_shop/pkg"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ErrorResponse представляет структуру ответа об ошибке,
// отправляемую клиенту в формате JSON.
type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

// Определение стандартных ошибок, используемых в приложении.
var (
	// ErrUnauthorized означает, что пользователь не авторизован.
	ErrUnauthorized = errors.New("unauthorized")
	// ErrNotEnoughCoins означает, что у пользователя недостаточно средств для выполнения операции.
	ErrNotEnoughCoins = errors.New("not enough coins")
	// ErrInvalidRequest означает, что запрос некорректен или отсутствуют необходимые поля.
	ErrInvalidRequest = errors.New("invalid request")
	// ErrUserNotFound означает, что запрашиваемый пользователь не найден.
	ErrUserNotFound = errors.New("user not found")
	// ErrItemNotFound означает, что запрашиваемый товар не найден.
	ErrItemNotFound = errors.New("item not found")
	// ErrSelfTransfer означает, что пользователь пытается выполнить перевод средств самому себе.
	ErrSelfTransfer = errors.New("cannot send coins to yourself")
	// ErrInternalServerError означает, что произошла внутренняя ошибка сервера.
	ErrInternalServerError = errors.New("internal server error")
)

// HandleError обрабатывает возникшую ошибку и отправляет соответствующий HTTP-ответ.
// Функция логирует ошибку и формирует JSON-ответ с нужным статусом.
//
// Параметры:
//   - c: контекст Gin, через который отправляется ответ клиенту.
//   - err: ошибка, которая должна быть обработана.
func HandleError(c *gin.Context, err error) {
	// Логирование ошибки с использованием глобального логгера
	pkg.Logger.Error("Error: ", err)

	// Определение типа ошибки и отправка соответствующего ответа
	switch {
	case errors.Is(err, ErrUnauthorized):
		c.JSON(http.StatusUnauthorized, ErrorResponse{
			Error:   "unauthorized",
			Message: "Invalid credentials",
		})
	case errors.Is(err, ErrNotEnoughCoins):
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "not_enough_coins",
			Message: "Not enough coins to complete the transaction",
		})
	case errors.Is(err, ErrInvalidRequest):
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "invalid_request",
			Message: "Request is invalid or missing required fields",
		})
	case errors.Is(err, ErrUserNotFound):
		c.JSON(http.StatusNotFound, ErrorResponse{
			Error:   "user_not_found",
			Message: "User not found",
		})
	case errors.Is(err, ErrItemNotFound):
		c.JSON(http.StatusNotFound, ErrorResponse{
			Error:   "item_not_found",
			Message: "Requested item not found",
		})
	case errors.Is(err, ErrSelfTransfer):
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "self_transfer",
			Message: "You cannot send coins to yourself",
		})
	default:
		// Для неизвестных ошибок отправляем ответ с кодом 500
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error:   "internal_error",
			Message: "An unexpected error occurred",
		})
	}
}
