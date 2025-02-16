// Package dto содержит определения структур данных (DTO),
// используемых для обмена информацией между клиентом и сервером.
package dto

// AuthRequest представляет структуру запроса на аутентификацию пользователя.
// Поля Username и Password являются обязательными для заполнения.
type AuthRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// AuthResponse представляет структуру ответа на запрос аутентификации.
// Поле Token содержит JWT токен, который выдается пользователю после успешной аутентификации.
type AuthResponse struct {
	Token string `json:"token"`
}
