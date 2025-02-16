// Package middleware содержит промежуточное ПО для обработки HTTP-запросов,
// в том числе для аутентификации с использованием JWT.
package middleware

import (
	"avito_shop/internal/utils"
	"avito_shop/pkg"
	"github.com/gin-gonic/gin"
	"strings"
)

// JWTAuthMiddleware создает middleware для проверки JWT токена в заголовке Authorization.
// Middleware извлекает токен, проверяет его корректность и при успешной валидации добавляет
// данные пользователя (идентификатор и имя пользователя) в контекст запроса.
//
// Параметры:
//   - secret: секретный ключ для валидации JWT.
//
// Возвращает:
//   - gin.HandlerFunc: функцию-обработчик, которую можно использовать в цепочке middleware Gin.
func JWTAuthMiddleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Получаем значение заголовка Authorization.
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			utils.HandleError(c, utils.ErrUnauthorized)
			c.Abort()
			return
		}

		// Разделяем заголовок на части, ожидая формат "Bearer <token>".
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			utils.HandleError(c, utils.ErrUnauthorized)
			c.Abort()
			return
		}

		// Извлекаем JWT токен.
		tokenString := tokenParts[1]

		// Валидируем токен с использованием заданного секретного ключа.
		id, username, err := pkg.ValidateJWT(tokenString, secret)
		if err != nil {
			utils.HandleError(c, utils.ErrUnauthorized)
			c.Abort()
			return
		}

		// Сохраняем данные пользователя в контекст запроса для последующего использования.
		c.Set("username", username)
		c.Set("id", id)

		c.Next()
	}
}
