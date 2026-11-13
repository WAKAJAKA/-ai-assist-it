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

// NewTodoHandler создает и инициализирует новый экземпляр TodoHandler.
func NewTodoHandler() *TodoHandler {
	return &TodoHandler{
		todos:  make(map[int]models.Todo),
		nextID: 1,
	}
}
// ListTodos возвращает список всех существующих задач.
// @Summary      Получить список всех задач
// @Description  Возвращает массив объектов Todo, обернутый в стандартный SuccessResponse. Если задач нет, поле data возвращает null.
// @Tags         todos
// @Produce      json
// @Success      200  {object}  models.SuccessResponse{data=[]models.Todo}  "Успешное получение списка задач (поле data равно null, если список пуст)"
// @Router       /todos/  [get]
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

// GetTodo возвращает задачу по её идентификатору.
// @Summary      Получить задачу по ID
// @Description  Возвращает одну задачу по переданному числовому идентификатору
// @Tags         todos
// @Produce      json,plain
// @Param        id   path      int                                       true  "Идентификатор задачи"
// @Success      200  {object}  models.SuccessResponse{data=models.Todo}  "Задача успешно найдена"
// @Failure      400  {string}  string                                    "Неверный ID (некорректный формат)"
// @Failure      404  {string}  string                                    "Задача не найдена"
// @Router       /todos/{id} [get]
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

// CreateTodo создает новую задачу.
// @Summary      Создать новую задачу
// @Description  Создает задачу на основе переданного JSON и возвращает созданный объект
// @Tags         todos
// @Accept       json
// @Produce      json,plain
// @Param        body  body      models.CreateTodoRequest                  true  "Данные для создания задачи"
// @Success      201   {object}  models.SuccessResponse{data=models.Todo}  "Задача успешно создана"
// @Failure      400   {string}  string                                    "Невалидный запрос (ошибка парсинга JSON)"
// @Router       /todos/ [post]
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

// UpdateTodo обновляет существующую задачу по ID.
// @Summary      Обновить существующую задачу
// @Description  Обновляет поля задачи по её числовому идентификатору
// @Tags         todos
// @Accept       json
// @Produce      json,plain
// @Param        id    path      int                       true  "Идентификатор задачи"
// @Param        body  body      models.UpdateTodoRequest  true  "Данные для обновления задачи"
// @Success      200   {object}  models.SuccessResponse{data=models.Todo}  "Задача успешно обновлена"
// @Failure      400   {string}  string                                    "Неверный ID или невалидный JSON"
// @Failure      404   {string}  string                                    "Задача не найдена"
// @Router       /todos/{id} [put]
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

// DeleteTodo удаляет задачу по ID.
// @Summary      Удалить задачу
// @Description  Удаляет задачу по её числовому идентификатору
// @Tags         todos
// @Produce      plain
// @Param        id   path      int     true  "Идентификатор задачи"
// @Success      204  "Задача успешно удалена (нет содержимого)"
// @Failure      400  {string}  string  "Неверный ID (некорректный формат)"
// @Failure      404  {string}  string  "Задача не найдена"
// @Router       /todos/{id} [delete]
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
