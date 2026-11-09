package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"

	"api-doc-example/internal/models"
)

// @BasePath /api/v1

type TodoHandler struct {
	todos  map[int]models.Todo
	nextID int
}

// NewTodoHandler creates a new TodoHandler
// @Summary Create a new TodoHandler
// @Description Initialize and return a new TodoHandler instance
// @Tags todos
// @Success 200 {object} TodoHandler "Successfully created TodoHandler"
// @Router /todos/init [get]
func NewTodoHandler() *TodoHandler {
	return &TodoHandler{
		todos:  make(map[int]models.Todo),
		nextID: 1,
	}
}

// ListTodos returns a list of all existing tasks.
// @Summary      Получить список всех задач
// @Description  Возвращает массив объектов Todo, обернутый в стандартный SuccessResponse
// @Tags         todos
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.SuccessResponse{data=[]models.Todo}  "Успешное получение списка задач"
// @Failure      500  {object}  models.ErrorResponse                        "Внутренняя ошибка сервера"
// @Router       /todos [get]
func (h *TodoHandler) ListTodos(w http.ResponseWriter, r *http.Request) {
	var todos []models.Todo
	for _, todo := range h.todos {
		todos = append(todos, todo)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(models.SuccessResponse{
		Success: true,
		Data:    todos,
	})
}

// GetTodo gets a todo by ID
// @Summary Получить задачу по ID
// @Description Получить задачу по её ID
// @Tags todos
// @Accept json
// @Produce json
// @Param id path int true "Todo ID"
// @Success 200 {object} models.SuccessResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Router /todos/{id} [get]
func (h *TodoHandler) GetTodo(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Неверный ID", http.StatusBadRequest)
		return
	}

	todo, exists := h.todos[id]
	if !exists {
		http.Error(w, "Задача не найдена", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(models.SuccessResponse{
		Success: true,
		Data:    todo,
	})
}

// CreateTodo creates a new todo
// @Summary Создать новую задачу
// @Description Создать новую задачу
// @Tags todos
// @Accept json
// @Produce json
// @Param body body models.CreateTodoRequest true "Todo data"
// @Success 201 {object} models.SuccessResponse{data=models.Todo} "Задача успешно создана"
// @Failure 400 {object} models.ErrorResponse
// @Router /todos [post]
func (h *TodoHandler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	var req models.CreateTodoRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Невалидный запрос", http.StatusBadRequest)
		return
	}

	todo := models.Todo{
		ID:          h.nextID,
		UserID:      req.UserID,
		Title:       req.Title,
		Description: req.Description,
		Done:        req.Done,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	h.todos[h.nextID] = todo
	h.nextID++

	log.Printf("Создана новая задача: %+v", todo)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(models.SuccessResponse{
		Success: true,
		Data:    todo,
	})
}

// UpdateTodo updates an existing todo
// @Summary Обновить существующую задачу
// @Description Обновить задачу по ID
// @Tags todos
// @Accept json
// @Produce json
// @Param id path int true "Todo ID"
// @Param body body models.UpdateTodoRequest true "Updated todo data"
// @Success 200 {object} models.SuccessResponse{data=models.Todo} "Задача успешно обновлена"
// @Failure 400 {object} models.ErrorResponse "Неверный формат запроса"
// @Failure 404 {object} models.ErrorResponse "Задача не найдена"
// @Router /todos/{id} [put]
func (h *TodoHandler) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Неверный ID", http.StatusBadRequest)
		return
	}

	var req models.UpdateTodoRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Невалидный запрос", http.StatusBadRequest)
		return
	}

	todo, exists := h.todos[id]
	if !exists {
		http.Error(w, "Задача не найдена", http.StatusNotFound)
		return
	}

	if req.Title != nil {
		todo.Title = *req.Title
	}
	if req.Description != nil {
		todo.Description = *req.Description
	}
	if req.Done != nil {
		todo.Done = *req.Done
	}
	todo.UpdatedAt = time.Now()

	h.todos[id] = todo

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(models.SuccessResponse{
		Success: true,
		Data:    todo,
	})
}

// DeleteTodo deletes a todo by ID
// @Summary Удалить задачу
// @Description Удалить задачу по её ID
// @Tags todos
// @Accept json
// @Produce json
// @Param id path int true "Todo ID"
// @Success 204 "Задача успешно удалена"
// @Failure 400 {object} models.ErrorResponse "Неверный формат запроса"
// @Failure 404 {object} models.ErrorResponse "Задача не найдена"
// @Router /todos/{id} [delete]
func (h *TodoHandler) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Неверный ID", http.StatusBadRequest)
		return
	}

	if _, exists := h.todos[id]; !exists {
		http.Error(w, "Задача не найдена", http.StatusNotFound)
		return
	}

	delete(h.todos, id)

	w.WriteHeader(http.StatusNoContent)
}
