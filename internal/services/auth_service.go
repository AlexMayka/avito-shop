package services

import (
	"avito_shop/internal/repositories"
	"avito_shop/internal/utils"
	"avito_shop/pkg"
	"database/sql"
)

// AuthService осуществляет аутентификацию пользователя.
// Если пользователь не существует, функция создает его с начальным балансом 1000.
// Если пользователь существует, проверяется корректность пароля.
// При успешной аутентификации генерируется JWT токен с использованием secretJWT.
//
// Параметры:
//   - db: подключение к базе данных.
//   - secretJWT: секрет для подписи JWT.
//   - username: имя пользователя.
//   - password: пароль пользователя.
//
// Возвращает:
//   - сгенерированный JWT токен.
//   - ошибку, если что-то пошло не так.
func AuthService(db *sql.DB, secretJWT, username, password string) (string, error) {
	user, err := repositories.GetUserByUsername(db, username)
	if err != nil {
		return "", utils.ErrInternalServerError
	}

	if user == nil {
		hashed, err := pkg.HashPassword(password)
		if err != nil {
			return "", utils.ErrInternalServerError
		}

		user, err = repositories.CreateUser(db, username, hashed, 1000)
		if err != nil {
			return "", utils.ErrInternalServerError
		}
	} else {
		if !pkg.CheckPasswordHash(password, user.Password) {
			return "", utils.ErrUnauthorized
		}
	}

	tokenJWT, err := pkg.GenerateJWT(user.ID, user.Username, secretJWT)
	if err != nil {
		return "", utils.ErrInternalServerError
	}

	return tokenJWT, nil
}
