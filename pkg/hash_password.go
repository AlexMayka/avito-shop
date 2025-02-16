// Package pkg содержит утилитарные функции для работы с паролями,
// включая хэширование паролей и проверку пароля по хэшу.
// Для работы используется алгоритм bcrypt из пакета "golang.org/x/crypto/bcrypt".
package pkg

import "golang.org/x/crypto/bcrypt"

// HashPassword принимает строку пароля и возвращает его bcrypt-хэш.
// Параметры:
//   - password: исходный пароль в виде строки.
//
// Возвращает:
//   - string: сгенерированный bcrypt-хэш пароля.
//   - error: ошибку, если процесс хэширования завершился неудачно.
//
// Используемая стоимость (cost) равна 7, что обеспечивает баланс между безопасностью и производительностью.
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 6)
	return string(bytes), err
}

// CheckPasswordHash сравнивает исходный пароль с его bcrypt-хэшем.
// Параметры:
//   - password: исходный пароль в виде строки.
//   - hash: bcrypt-хэш, с которым производится сравнение.
//
// Возвращает:
//   - bool: true, если пароль соответствует хэшу, и false в противном случае.
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
