// Package config обеспечивает загрузку конфигурационных параметров приложения
// из переменных окружения и файла .env.
package config

import (
	"os"

	"github.com/joho/godotenv"
)

// Host содержит настройки для HTTP-сервера и приложения.
type Host struct {
	// ServerHost определяет адрес сервера (например, "localhost").
	ServerHost string
	// ServerPort определяет порт, на котором работает сервер (например, "8080").
	ServerPort string
	// AppEnv указывает текущее окружение приложения (например, "development", "production").
	AppEnv string
	// JWTSecret используется для подписи JWT токенов.
	JWTSecret string
}

// Db содержит параметры для подключения к базе данных.
type Db struct {
	// Host указывает адрес хоста базы данных (например, "localhost").
	Host string
	// User определяет имя пользователя для подключения к базе данных.
	User string
	// Password задаёт пароль для подключения к базе данных.
	Password string
	// Db определяет имя базы данных.
	Db string
	// Port указывает порт, на котором доступна база данных (например, "5432").
	Port string
}

// Config объединяет настройки для сервера и базы данных.
type Config struct {
	Host Host
	Db   Db
}

// LoadConfig загружает конфигурационные параметры из переменных окружения,
// используя библиотеку godotenv для чтения из файла .env.
// Если переменная окружения не задана, используется значение по умолчанию.
func LoadConfig() *Config {
	// Загружаем переменные окружения из файла .env, если он существует.
	_ = godotenv.Load()

	return &Config{
		Host: Host{
			ServerHost: getEnv("SERVER_HOST", "localhost"),
			ServerPort: getEnv("SERVER_PORT", "8080"),
			AppEnv:     getEnv("APP_ENV", "development"),
			JWTSecret:  getEnv("JWT_SECRET", "secret"),
		},
		Db: Db{
			Host:     getEnv("POSTGRES_HOST", "localhost"),
			User:     getEnv("POSTGRES_USER", "postgres"),
			Password: getEnv("POSTGRES_PASSWORD", "password"),
			Db:       getEnv("POSTGRES_DB", "shop"),
			Port:     getEnv("POSTGRES_PORT", "5432"),
		},
	}
}

// getEnv возвращает значение переменной окружения для заданного ключа key.
// Если переменная не установлена, возвращается значение по умолчанию defaultValue.
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
