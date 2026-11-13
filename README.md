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

## Запуск проекта

Если вам нужно просто запустить сервис, готовая документация уже сгенерирована в папке `docs/`:

```bash
go run main.go
```

Интерфейс Swagger UI откроется по адресу: **http://localhost:8080/swagger/index.html**

---

### Для разработчиков: обновление документации Swagger

Если вы изменили хэндлеры, модели или Swagger-аннотации, перегенерируйте спецификацию:

1. Установите утилиту `swag` (если ещё не установлена):

```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

2. Сгенерируйте обновлённые файлы документации:

```bash
swag init -g main.go --parseInternal
```

> Если терминал пишет `swag: command not found`, используйте `$(go env GOPATH)/bin/swag init -g main.go --parseInternal`

Команда обновит файлы в каталоге `docs/`:

- `docs.go`
- `swagger.json`
- `swagger.yaml`

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

## Известные ограничения текущей версии

* **Конкурентный доступ в `TodoHandler`:** В отличие от `UserHandler` (где используется `sync.RWMutex`), хранилище задач в `TodoHandler` реализовано на базе стандартной `map` без синхронизации потоков. Эндпоинты `/api/v1/todos/` в текущей версии предназначены только для однопоточного (демонстрационного) использования.
* **Формат пустых списков (`null` вместо `[]`):** При отсутствии записей в памяти эндпоинты `GET /api/v1/todos/` и `GET /api/v1/users/` возвращают `"data": null`.
* **Отсутствие серверной валидации полей и границ ID:** Сервер выполняет только базовый парсинг JSON и целочисленных идентификаторов без проверки длины строк, формата email и положительных значений ID (`id >= 1`).
* **Формат ответов об ошибках:** Все ошибки возвращаются в виде обычного текста.

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
├── static/              # Статические файлы
│   └── index.html       # Welcome-страница (http://localhost:8080/)
├── go.mod               # Зависимости проекта
├── go.sum               # Контрольные суммы пакетов
└── README.md            # Документация проекта
```

### Welcome-страница

При переходе на `http://localhost:8080/` открывается welcome-страница (`static/index.html`), которая содержит:
- Подтверждение успешного запуска сервиса
- Ссылки на основные эндпоинты API
- Ссылку на Swagger UI

---

## Лицензия и автор

**Лицензия**: MIT

**Автор**: Когут И.В.

**Репозиторий**: https://github.com/WAKAJAKA/-ai-assist-it
