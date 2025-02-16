package e2e

import (
	"avito_shop/config"
	"github.com/joho/godotenv"
	"os"
)

func LoadConfig() *config.Config {
	_ = godotenv.Load("env.test")
	return &config.Config{
		Host: config.Host{
			ServerHost: getEnv("SERVER_HOST", "localhost"),
			ServerPort: getEnv("SERVER_PORT", "8080"),
			AppEnv:     getEnv("APP_ENV", "development"),
			JWTSecret:  getEnv("JWT_SECRET", "secret"),
		},
		Db: config.Db{
			Host:     getEnv("POSTGRES_HOST", "db"),
			User:     getEnv("POSTGRES_USER", "postgres"),
			Password: getEnv("POSTGRES_PASSWORD", "password"),
			Db:       getEnv("POSTGRES_DB", "shop_test"),
			Port:     getEnv("POSTGRES_PORT", "5432"),
		},
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
