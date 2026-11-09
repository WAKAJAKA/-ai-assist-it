// Package models содержит структуры данных предметной области (доменные модели)
// и модели запросов (DTO) для работы с задачами (Todo).
package models

import "time"

// Todo представляет задачу в системе.
type Todo struct {
	// ID — уникальный идентификатор задачи.
	ID int `json:"id"`

	// UserID — идентификатор пользователя, которому принадлежит задача.
	UserID int `json:"user_id"`

	// Title — краткий заголовок задачи.
	Title string `json:"title"`

	// Description — подробное описание задачи.
	Description string `json:"description"`

	// Done указывает, завершена ли задача.
	Done bool `json:"done"`

	// CreatedAt — метка времени создания записи.
	CreatedAt time.Time `json:"created_at"`

	// UpdatedAt — метка времени последнего обновления записи.
	UpdatedAt time.Time `json:"updated_at"`
}

// CreateTodoRequest содержит входные данные для создания новой задачи.
type CreateTodoRequest struct {
	// UserID — идентификатор пользователя, создающего задачу (обязательное поле).
	UserID int `json:"user_id" validate:"required"`

	// Title — заголовок задачи (обязательное поле, до 200 символов).
	Title string `json:"title" validate:"required,max=200"`

	// Description — подробное описание задачи (опционально, до 1000 символов).
	Description string `json:"description" validate:"max=1000"`

	// Done — начальный статус выполнения задачи.
	Done bool `json:"done"`
}

// UpdateTodoRequest содержит данные для частичного обновления существующей задачи.
// Использование указателей позволяет различать непереданные поля (nil)
// от переданных значений по умолчанию (например, пустая строка или false).
type UpdateTodoRequest struct {
	// Title — новый заголовок задачи (опционально, до 200 символов).
	Title *string `json:"title" validate:"omitempty,max=200"`

	// Description — новое описание задачи (опционально, до 1000 символов).
	Description *string `json:"description" validate:"omitempty,max=1000"`

	// Done — новый статус выполнения задачи (опционально).
	Done *bool `json:"done"`
}
