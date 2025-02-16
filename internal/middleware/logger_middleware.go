// Package middleware содержит промежуточное ПО для обработки HTTP-запросов,
// включая логирование входящих запросов с использованием logrus.
package middleware

import (
	"avito_shop/pkg"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// LoggerMiddleware создает middleware для логирования информации о каждом HTTP-запросе.
// Middleware фиксирует время начала обработки запроса, а после выполнения запроса
// вычисляет длительность обработки и логирует следующие данные:
//   - HTTP-метод запроса
//   - URL запроса
//   - HTTP-статус ответа
//   - IP-адрес клиента
//   - Время, затраченное на обработку запроса
//
// Возвращает:
//   - gin.HandlerFunc: функцию-обработчик, которая добавляется в цепочку middleware Gin.
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Фиксируем время начала обработки запроса.
		start := time.Now()

		c.Next()
		duration := time.Since(start)

		pkg.Logger.WithFields(logrus.Fields{
			"method": c.Request.Method,   // HTTP-метод запроса
			"path":   c.Request.URL.Path, // URL запроса
			"status": c.Writer.Status(),  // HTTP-статус ответа
			"ip":     c.ClientIP(),       // IP-адрес клиента
			"time":   duration,           // Длительность обработки запроса
		}).Info("Request processed")
	}
}
