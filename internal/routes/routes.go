// Package routes настраивает маршруты HTTP-сервера с использованием Gin.
// В данном пакете определяются публичные и защищённые маршруты, а также применяется middleware.
package routes

import (
	"avito_shop/config"
	"avito_shop/internal/controllers"
	"avito_shop/internal/middleware"
	"database/sql"

	"github.com/gin-gonic/gin"
)

// SetupRouter инициализирует маршрутизатор Gin, регистрирует маршруты и подключает middleware.
//
// Параметры:
//   - database: указатель на sql.DB для работы с базой данных.
//   - cfg: указатель на конфигурацию приложения, содержащую настройки сервера, базы данных и JWT.
//
// Возвращает:
//   - *gin.Engine: настроенный маршрутизатор, готовый к запуску HTTP-сервера.
//
// Функция выполняет следующие действия:
//  1. Создаёт экземпляр маршрутизатора Gin с настройками по умолчанию.
//  2. Инициализирует контроллер, передавая ему подключение к базе данных и конфигурацию.
//  3. Применяет глобальное middleware для логирования запросов.
//  4. Определяет группу публичных маршрутов (без JWT аутентификации):
//     - POST /api/auth: маршрут для аутентификации пользователя.
//  5. Определяет группу защищённых маршрутов (с JWT аутентификацией):
//     - GET /api/buy/:item: маршрут для покупки мерча.
//     - GET /api/info: маршрут для получения информации о пользователе.
//     - POST /api/sendCoin: маршрут для перевода монет между пользователями.
func SetupRouter(database *sql.DB, cfg *config.Config) *gin.Engine {
	r := gin.Default()
	ctr := controllers.Controller{DB: database, CFG: cfg}
	r.Use(middleware.LoggerMiddleware())

	apiPublic := r.Group("/api")
	{
		apiPublic.POST("/auth", ctr.AuthHandler)
	}

	apiProtected := r.Group("/api")
	apiProtected.Use(middleware.JWTAuthMiddleware(cfg.Host.JWTSecret))
	{
		apiProtected.GET("/buy/:item", ctr.BuyMerchHandler)
		apiProtected.GET("/info", ctr.InfoHandler)
		apiProtected.POST("/sendCoin", ctr.SendCoinController)
	}

	return r
}
