# Todo REST API Service

REST API сервис на Go для управления задачами (Todo API) с автоматической интерактивной документацией Swagger (OpenAPI 2.0).

---

## Содержание

1. [Описание проекта](#описание-проекта)
2. [Требования к окружению](#требования-к-окружению)
3. [Используемые зависимости](#используемые-зависимости)
4. [Установка и запуск](#установка-и-запуск)
5. [Документация API (Swagger UI)](#документация-api-swagger-ui)
6. [Примеры использования API (cURL)](#примеры-использования-api-curl)
   - [1. Получить список всех задач](#1-получить-список-всех-задач)
   - [2. Создать новую задачу](#2-создать-новую-задачу)
   - [3. Получить задачу по ID](#3-получить-задачу-по-id)
   - [4. Обновить задачу по ID](#4-обновить-задачу-по-id)
   - [5. Удалить задачу по ID](#5-удалить-задачу-по-id)
   - [Примеры обработки ошибок](#примеры-обработки-ошибок)
7. [Структура проекта](#структура-проекта)
8. [Лицензия и контакты](#лицензия-и-контакты)

---

## Описание проекта

Сервис предоставляет HTTP REST API для создания, чтения, частичного обновления и удаления задач пользователей.

**Ключевые особенности:**
- Архитектура на основе стандартной библиотеки `net/http` и легковесного роутера `go-chi/chi`.
- Унифицированный формат ответов через обёртки `SuccessResponse` и `ErrorResponse`.
- Декларативная документация через аннотации `swaggo/swag` и интерактивный UI.

---

## Требования к окружению

- **Go**: версия `1.21` или выше
- **Git**
- Утилита кодогенерации **swag** (CLI для сборки спецификации Swagger)

Установка генератора Swagger CLI:
```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

---

## Используемые зависимости

- `github.com/go-chi/chi/v5` — легковесный, идиоматичный роутер HTTP-запросов.
- `github.com/swaggo/swag` — парсинг аннотаций Go в спецификацию Swagger 2.0.
- `github.com/swaggo/http-swagger` — HTTP-хендлер для раздачи Swagger UI.

---

## Установка и запуск

### 1. Клонирование репозитория

```bash
git clone https://github.com/WAKAJAKA/-ai-assist-it.git
cd ai-assist-it
```

### 2. Загрузка зависимостей

```bash
go mod download
```

### 3. Генерация документации Swagger

Перед первым запуском или после изменения аннотаций в хендлерах обновите файлы спецификации:

```bash
swag init
```

После выполнения команды в директории `docs/` будут созданы/обновлены файлы:
- `docs/docs.go`
- `docs/swagger.json`
- `docs/swagger.yaml`

### 4. Запуск сервиса

```bash
go run main.go
```

По умолчанию сервер стартует на порту **8080** (адрес: `http://localhost:8080`).

---

## Документация API (Swagger UI)

После запуска приложения интерактивная документация и песочница доступны по адресу:

👉 **http://localhost:8080/swagger/index.html**

Сырая спецификация в формате JSON:
**http://localhost:8080/swagger/doc.json**

---

## Примеры использования API (cURL)

Базовый URL: `http://localhost:8080`

### 1. Получить список всех задач

**Метод:** GET
**Маршрут:** `/todos`
**Код ответа:** 200 OK

```bash
curl -X GET http://localhost:8080/todos \
  -H "Accept: application/json"
```

**Пример ответа (200 OK):**

```json
{
  "success": true,
  "data": [
    {
      "id": 1,
      "user_id": 101,
      "title": "Написать документацию",
      "description": "Оформить README.md и сгенерировать Swagger",
      "done": false,
      "created_at": "2026-11-09T10:00:00Z",
      "updated_at": "2026-11-09T10:00:00Z"
    }
  ]
}
```

### 2. Создать новую задачу

**Метод:** POST
**Маршрут:** `/todos`
**Код ответа:** 201 Created

```bash
curl -X POST http://localhost:8080/todos \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": 101,
    "title": "Интегрировать Swagger",
    "description": "Добавить аннотации ко всем эндпоинтам",
    "done": false
  }'
```

**Пример ответа (201 Created):**

```json
{
  "success": true,
  "data": {
    "id": 1,
    "user_id": 101,
    "title": "Интегрировать Swagger",
    "description": "Добавить аннотации ко всем эндпоинтам",
    "done": false,
    "created_at": "2026-11-09T10:15:30.123456Z",
    "updated_at": "2026-11-09T10:15:30.123456Z"
  }
}
```

### 3. Получить задачу по ID

**Метод:** GET
**Маршрут:** `/todos/{id}`
**Код ответа:** 200 OK

```bash
curl -X GET http://localhost:8080/todos/1 \
  -H "Accept: application/json"
```

**Пример ответа (200 OK):**

```json
{
  "success": true,
  "data": {
    "id": 1,
    "user_id": 101,
    "title": "Интегрировать Swagger",
    "description": "Добавить аннотации ко всем эндпоинтам",
    "done": false,
    "created_at": "2026-11-09T10:15:30.123456Z",
    "updated_at": "2026-11-09T10:15:30.123456Z"
  }
}
```

### 4. Обновить задачу по ID

**Метод:** PUT
**Маршрут:** `/todos/{id}`
**Код ответа:** 200 OK

```bash
curl -X PUT http://localhost:8080/todos/1 \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Интегрировать Swagger (Готово)",
    "done": true
  }'
```

**Пример ответа (200 OK):**

```json
{
  "success": true,
  "data": {
    "id": 1,
    "user_id": 101,
    "title": "Интегрировать Swagger (Готово)",
    "description": "Добавить аннотации ко всем эндпоинтам",
    "done": true,
    "created_at": "2026-11-09T10:15:30.123456Z",
    "updated_at": "2026-11-09T10:20:00.654321Z"
  }
}
```

### 5. Удалить задачу по ID

**Метод:** DELETE
**Маршрут:** `/todos/{id}`
**Код ответа:** 204 No Content

```bash
curl -i -X DELETE http://localhost:8080/todos/1
```

**Пример ответа (204 No Content):**

```
HTTP/1.1 204 No Content
Date: Mon, 09 Nov 2026 10:25:00 GMT
```

*(Тело ответа пустое)*

### Примеры обработки ошибок

**Некорректный ID (400 Bad Request)**

```bash
curl -X GET http://localhost:8080/todos/invalid_id
```

Ответ:
```
Неверный ID
```

**Задача не найдена (404 Not Found)**

```bash
curl -X GET http://localhost:8080/todos/9999
```

Ответ:
```
Задача не найдена
```

---

## Структура проекта

```
.
├── cmd/
│   └── api/
│       └── main.go              # Точка входа в приложение, инициализация роутера и Swagger
├── docs/                        # Сгенерированные файлы Swagger (swag init)
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── internal/
│   ├── handlers/                # HTTP-обработчики запросов со Swagger-аннотациями
│   │   ├── todo_handler.go      # CRUD операции для задач
│   │   └── user_handler.go      # Операции для пользователей
│   └── models/                  # Доменные структуры данных и DTO
│       ├── todo.go              # Модели Todo, CreateTodoRequest, UpdateTodoRequest
│       └── user.go              # Модели User, SuccessResponse, ErrorResponse
├── go.mod                       # Определение модуля Go и зависимостей
├── go.sum                       # Контрольные суммы зависимостей
└── README.md                    # Документация проекта
```

---

## Лицензия и контакты

**Лицензия:** Проект распространяется под лицензией MIT.

**Автор:** Когут И.В

**Email:** igorek.kogut@gmail.com

**Репозиторий:** https://github.com/WAKAJAKA/-ai-assist-it.git
