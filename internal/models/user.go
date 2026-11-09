// Package models содержит структуры данных предметной области (доменные модели),
// DTO для запросов/ответов API и вспомогательные структуры ответов.
package models

import "time"

// User представляет пользователя в системе.
type User struct {
	// ID — уникальный идентификатор пользователя.
	ID int `json:"id"`

	// Email — адрес электронной почты пользователя.
	Email string `json:"email"`

	// Name — имя или отображаемое имя пользователя.
	Name string `json:"name"`

	// CreatedAt — дата и время регистрации пользователя.
	CreatedAt time.Time `json:"created_at"`
}

// CreateUserRequest содержит входные данные для регистрации/создания нового пользователя.
type CreateUserRequest struct {
	// Email — адрес электронной почты (обязательное поле, должен быть валидным email).
	Email string `json:"email" validate:"required,email"`

	// Name — имя пользователя (обязательное поле, до 100 символов).
	Name string `json:"name" validate:"required,max=100"`
}

// UpdateUserRequest содержит данные для частичного обновления профиля пользователя.
// Поля определены как указатели, чтобы отличать отсутствие изменения (nil)
// от передачи пустого значения.
type UpdateUserRequest struct {
	// Email — новый адрес электронной почты (опционально, должен быть валидным email).
	Email *string `json:"email" validate:"omitempty,email"`

	// Name — новое имя пользователя (опционально, до 100 символов).
	Name *string `json:"name" validate:"omitempty,max=100"`
}

// ErrorResponse представляет стандартный формат ответа API при возникновении ошибки.
type ErrorResponse struct {
	// Error — машиночитаемый код или краткое обозначение ошибки.
	Error string `json:"error"`

	// Message — понятное человеку описание возникшей ошибки.
	Message string `json:"message"`
}

// SuccessResponse представляет стандартную обёртку для успешного ответа API.
type SuccessResponse struct {
	// Success указывает на успешность выполнения операции.
	Success bool `json:"success"`

	// Data содержит полезную нагрузку ответа (может быть опущено, если данных нет).
	Data interface{} `json:"data,omitempty"`
}
