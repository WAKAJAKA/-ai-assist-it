// Package models содержит структуры данных предметной области (доменные модели),
// DTO для запросов/ответов API и вспомогательные структуры ответов.
package models

import "time"

// User представляет пользователя в системе.
// @Description Модель пользователя с базовыми регистрационными данными
type User struct {
	// ID — уникальный идентификатор пользователя.
	ID int `json:"id" example:"1"`

	// Email — адрес электронной почты пользователя.
	Email string `json:"email" format:"email" example:"ivan@example.com"`

	// Name — имя или отображаемое имя пользователя.
	Name string `json:"name" example:"Иван Иванов"`

	// CreatedAt — дата и время регистрации пользователя.
	CreatedAt time.Time `json:"created_at" example:"2026-03-30T10:00:00Z"`
}

// CreateUserRequest содержит входные данные для регистрации/создания нового пользователя.
// @Description Данные для регистрации нового пользователя (DTO)
type CreateUserRequest struct {
	// Email — адрес электронной почты.
	Email string `json:"email" format:"email" example:"ivan@example.com"`

	// Name — имя пользователя.
	Name string `json:"name" example:"Иван Иванов"`
}

// UpdateUserRequest содержит данные для частичного обновления профиля пользователя.
// Поля определены как указатели, чтобы отличать отсутствие изменения (nil)
// от передачи пустого значения.
// @Description Данные для частичного обновления профиля пользователя (все поля опциональны)
type UpdateUserRequest struct {
	// Email — новый адрес электронной почты (опционально).
	Email *string `json:"email" format:"email" example:"new_ivan@example.com"`

	// Name — новое имя пользователя (опционально).
	Name *string `json:"name" example:"Иван Петров"`
}

// ErrorResponse представляет стандартный формат ответа API при возникновении ошибки.
// @Description Стандартный формат ответа API при возникновении ошибки
type ErrorResponse struct {
	// Error — машиночитаемый код или краткое обозначение ошибки.
	Error string `json:"error" example:"not_found"`

	// Message — понятное человеку описание возникшей ошибки.
	Message string `json:"message" example:"Запрашиваемый ресурс не найден"`
}

// SuccessResponse представляет стандартную обёртку для успешного ответа API.
// @Description Стандартная обёртка для успешного ответа API с полезной нагрузкой
type SuccessResponse struct {
	// Success указывает на успешность выполнения операции.
	Success bool `json:"success" example:"true"`

	// Data содержит полезную нагрузку ответа (может быть опущено, если данных нет).
	Data interface{} `json:"data,omitempty" swaggertype:"object"`
}
