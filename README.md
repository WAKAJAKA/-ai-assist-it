# Todo REST API Service

Сервис для управления задачами (Todo API) и пользователями, написанный на языке Go. Предоставляет программный интерфейс (API) для работы по протоколу HTTP и оснащен интерактивной веб-документацией Swagger (OpenAPI 2.0).

---

## Возможности сервиса

- **Пользователи**: полный CRUD — просмотр списка, получение по ID, создание, обновление, удаление
- **Задачи**: полный CRUD — просмотр списка, получение по ID, создание, обновление, удаление
- **Веб-документация**: просмотр описания всех маршрутов и отправка тестовых запросов через Swagger UI
- **Единый формат ответов**: все успешные данные возвращаются в объекте с отметкой успеха (`"success": true`) и полезной нагрузкой (`"data"`)

---

## Требования к окружению

Для сборки и запуска проекта потребуется:

1. **Go** — версии `1.22` или новее ([скачать с официального сайта](https://go.dev/dl/))
2. **Git** — для клонирования репозитория

Проверка установки:

```bash
go version
```

---

## Установка и запуск

### 1. Клонирование и установка зависимостей

```bash
git clone https://github.com/WAKAJAKA/ai-assist-it.git
cd ai-assist-it
go mod download
```

### 2. Установка swag и генерация Swagger-документации

Для генерации файлов документации установите утилиту swag (если еще не установлена):

```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

Сгенерируйте спецификацию:

```bash
swag init -g main.go
```

В каталоге `docs/` будут обновлены файлы:
- `docs.go`
- `swagger.json`
- `swagger.yaml`

### 3. Запуск сервера

```bash
go run main.go
```

Сервер запустится по адресу: **http://localhost:8080**

---

## Документация API (Swagger UI)

Интерактивная документация и песочница:

👉 **http://localhost:8080/swagger/index.html**

Спецификация OpenAPI (JSON):

**http://localhost:8080/docs/swagger.json**

---

## Примеры использования API (cURL)

Базовый адрес API: `http://localhost:8080/api/v1`

### Пользователи (Users)

**1. Получить список всех пользователей**

```bash
curl -X GET http://localhost:8080/api/v1/users -H "Accept: application/json"
```

**2. Создать нового пользователя**

```bash
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{"email": "test@mail.com", "name": "Иван"}'
```

Пример ответа (HTTP 201):

```json
{
  "success": true,
  "data": {
    "id": 1,
    "email": "test@mail.com",
    "name": "Иван",
    "created_at": "2026-03-30T10:00:00Z"
  }
}
```

**3. Получить пользователя по ID**

```bash
curl -X GET http://localhost:8080/api/v1/users/1 -H "Accept: application/json"
```

**4. Обновить пользователя**

```bash
curl -X PUT http://localhost:8080/api/v1/users/1 \
  -H "Content-Type: application/json" \
  -d '{"name": "Иван Иванов"}'
```

**5. Удалить пользователя**

```bash
curl -i -X DELETE http://localhost:8080/api/v1/users/1
```

Ответ: `HTTP 204 No Content` (тело ответа пустое).

---

### Задачи (Todos)

**1. Получить список всех задач**

```bash
curl -X GET http://localhost:8080/api/v1/todos -H "Accept: application/json"
```

**2. Создать новую задачу**

```bash
curl -X POST http://localhost:8080/api/v1/todos \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": 1,
    "title": "Новая задача",
    "description": "Описание задачи",
    "done": false
  }'
```

Пример ответа (HTTP 201):

```json
{
  "success": true,
  "data": {
    "id": 1,
    "user_id": 1,
    "title": "Новая задача",
    "description": "Описание задачи",
    "done": false,
    "created_at": "2026-03-30T10:05:00Z",
    "updated_at": "2026-03-30T10:05:00Z"
  }
}
```

**3. Получить задачу по ID**

```bash
curl -X GET http://localhost:8080/api/v1/todos/1 -H "Accept: application/json"
```

**4. Обновить задачу**

```bash
curl -X PUT http://localhost:8080/api/v1/todos/1 \
  -H "Content-Type: application/json" \
  -d '{"title": "Обновлённая задача", "done": true}'
```

**5. Удалить задачу**

```bash
curl -i -X DELETE http://localhost:8080/api/v1/todos/1
```

Ответ: `HTTP 204 No Content` (тело ответа пустое).

---

## Примеры ответов с ошибками

| Ситуация | HTTP-статус | Текст ответа |
|----------|-------------|--------------|
| Неверный ID (не число) | 400 Bad Request | Неверный ID |
| Задача / пользователь не найден | 404 Not Found | Задача не найдена / Пользователь не найден |
| Невалидный формат JSON | 400 Bad Request | Невалидный запрос |

---

## Структура проекта

```
.
├── main.go              # Точка входа: инициализация маршрутов, запуск сервера
├── docs/                # Сгенерированные файлы Swagger
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── internal/
│   ├── handlers/        # HTTP-обработчики запросов
│   │   ├── todo.go      # CRUD для задач
│   │   └── user.go      # CRUD для пользователей
│   └── models/          # Модели данных и DTO
│       ├── todo.go
│       └── user.go
├── go.mod               # Зависимости проекта
├── go.sum               # Контрольные суммы пакетов
└── README.md            # Документация проекта
```

---

## Лицензия и автор

**Лицензия**: MIT

**Автор**: Когут И.В.

**Репозиторий**: https://github.com/WAKAJAKA/ai-assist-it.git
