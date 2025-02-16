// Package pkg содержит утилитарные функции для работы с JWT,
// включая генерацию и валидацию токенов для аутентификации пользователей.
package pkg

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Claims представляет набор утверждений (claims) для JWT.
// Включает пользовательские данные и стандартные утверждения, определённые в jwt.RegisteredClaims.
type Claims struct {
	// Id пользователя.
	Id uint `json:"id"`
	// Username пользователя.
	Username string `json:"username"`
	// RegisteredClaims содержит стандартные утверждения, такие как время истечения срока действия и дата выпуска токена.
	jwt.RegisteredClaims
}

// GenerateJWT создает новый JWT для заданного пользователя.
// Параметры:
//   - id: уникальный идентификатор пользователя.
//   - username: имя пользователя.
//   - jwtSecret: секретный ключ, используемый для подписи токена.
//
// Возвращает:
//   - tokenString: сгенерированный JWT в виде строки.
//   - error: ошибку, если процесс генерации токена завершился неудачно.
//
// Токен подписывается с использованием алгоритма HS256 и имеет срок действия 1 час.
func GenerateJWT(id uint, username, jwtSecret string) (string, error) {
	// Преобразуем секрет в срез байт
	secret := []byte(jwtSecret)

	// Формируем набор утверждений для токена
	claims := Claims{
		Id:       id,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)), // Токен истекает через 1 час
			IssuedAt:  jwt.NewNumericDate(time.Now()),                // Время выпуска токена
			Subject:   "user-auth",                                   // Предназначение токена
		},
	}

	// Создаём новый токен с использованием алгоритма HS256 и набора утверждений
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Подписываем токен с использованием секретного ключа
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ValidateJWT проверяет валидность JWT и извлекает из него идентификатор пользователя и имя пользователя.
// Параметры:
//   - tokenString: строковое представление JWT.
//   - jwtSecret: секретный ключ, используемый для проверки подписи токена.
//
// Возвращает:
//   - id: идентификатор пользователя, извлеченный из токена.
//   - username: имя пользователя, извлеченное из токена.
//   - error: ошибку, если токен недействителен или произошла ошибка при его разборе.
func ValidateJWT(tokenString, jwtSecret string) (uint, string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})
	if err != nil {
		return 0, "", err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return 0, "", err
	}

	return claims.Id, claims.Username, nil
}
