// Package models содержит структуры данных предметной области (доменные модели)
// и модели запросов (DTO) для работы с задачами (Todo).
package models

import "time"

// Todo представляет задачу в системе.
// @Description Модель задачи с полной информацией о статусе и времени создания
type Todo struct {
	// ID — уникальный идентификатор задачи.
	ID int `json:"id" example:"1"`

	// UserID — идентификатор пользователя, которому принадлежит задача.
	UserID int `json:"user_id" example:"1"`

	// Title — краткий заголовок задачи.
	Title string `json:"title" example:"Купить продукты"`

	// Description — подробное описание задачи.
	Description string `json:"description" example:"Купить молоко, хлеб и сыр"`

	// Done указывает, завершена ли задача.
	Done bool `json:"done" example:"false"`

	// CreatedAt — метка времени создания записи.
	CreatedAt time.Time `json:"created_at" example:"2026-03-30T10:00:00Z"`

	// UpdatedAt — метка времени последнего обновления записи.
	UpdatedAt time.Time `json:"updated_at" example:"2026-03-30T10:00:00Z"`
}

// CreateTodoRequest содержит входные данные для создания новой задачи.
// @Description Данные для создания новой задачи (DTO)
type CreateTodoRequest struct {
	// UserID — идентификатор пользователя, создающего задачу (обязательное поле).
	UserID int `json:"user_id" validate:"required" binding:"required" example:"1"`

	// Title — заголовок задачи (обязательное поле, до 200 символов).
	Title string `json:"title" validate:"required,max=200" binding:"required" maxLength:"200" example:"Купить продукты"`

	// Description — подробное описание задачи (опционально, до 1000 символов).
	Description string `json:"description" validate:"max=1000" maxLength:"1000" example:"Купить молоко, хлеб и яблоки"`

	// Done — начальный статус выполнения задачи.
	Done bool `json:"done" example:"false"`
}
// UpdateTodoRequest содержит данные для частичного обновления существующей задачи.
// Использование указателей позволяет различать непереданные поля (nil)
// от переданных значений по умолчанию (например, пустая строка или false).
// @Description Данные для частичного обновления задачи (все поля опциональны)
type UpdateTodoRequest struct {
	// Title — новый заголовок задачи (опционально, до 200 символов).
	Title *string `json:"title" validate:"omitempty,max=200" maxLength:"200" example:"Купить продукты (обновлено)"`

	// Description — новое описание задачи (опционально, до 1000 символов).
	Description *string `json:"description" validate:"omitempty,max=1000" maxLength:"1000" example:"Купить безлактозное молоко и цельнозерновой хлеб"`

	// Done — новый статус выполнения задачи (опционально).
	Done *bool `json:"done" example:"true"`
}