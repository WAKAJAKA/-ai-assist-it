# Todo REST API Service

Сервис для управления задачами (Todo API) и пользователями, написанный на языке Go. Предоставляет программный интерфейс (API) для работы по протоколу HTTP и оснащен интерактивной веб-документацией Swagger (OpenAPI 2.0).

---

## Возможности сервиса

- **Пользователи**: полный CRUD — просмотр списка, получение по ID, создание, обновление, удаление
- **Задачи**: полный CRUD — просмотр списка, получение по ID, создание, обновление, удаление
- **Веб-документация**: просмотр описания всех маршрутов и отправка тестовых запросов через браузер
- **Единый формат ответов**: все успешные данные возвращаются в объекте с отметкой успеха (`"success": true`) и полезной нагрузкой (`"data"`)

---

## Требования к окружению

Для сборки и запуска проекта потребуется:

1. **Go** — версии `1.22` или новее
2. **Git** — для клонирования репозитория

Установка Go:
- Скачайте с [golang.org/dl](https://golang.org/dl) и установите
- Проверьте установку: `go version`

---

## Установка и запуск

### 1. Клонирование и установка зависимостей

```bash
git clone https://github.com/WAKAJAKA/ai-assist-it.git
cd ai-assist-it
go mod download
```

### 2. Генерация Swagger-документации

```bash
swag init
```

В каталоге `docs/` будут созданы файлы:
- `docs.go`
- `swagger.json`
- `swagger.yaml`

### 3. Запуск сервера

```bash
go run main.go
```

Сервер запустится на **http://localhost:8080**

---

## Документация API (Swagger UI)

Интерактивная документация и песочница:
👉 **http://localhost:8080/swagger/index.html**

Сырая спецификация JSON:
**http://localhost:8080/swagger/doc.json**

---

## Примеры использования API (cURL)

Базовый адрес: `http://localhost:8080`

### Задачи (Todos)

**Получить список всех задач**
```bash
curl -X GET http://localhost:8080/todos -H "Accept: application/json"
```

**Создать новую задачу**
```bash
curl -X POST http://localhost:8080/todos \
  -H "Content-Type: application/json" \
  -d '{"user_id": 1, "title": "Новая задача", "description": "Описание", "done": false}'
```

**Получить задачу по ID**
```bash
curl -X GET http://localhost:8080/todos/1 -H "Accept: application/json"
```

**Обновить задачу**
```bash
curl -X PUT http://localhost:8080/todos/1 \
  -H "Content-Type: application/json" \
  -d '{"title": "Обновлённая задача", "done": true}'
```

**Удалить задачу**
```bash
curl -i -X DELETE http://localhost:8080/todos/1
```

### Пользователи (Users)

**Получить список всех пользователей**
```bash
curl -X GET http://localhost:8080/users -H "Accept: application/json"
```

**Создать нового пользователя**
```bash
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{"email": "test@mail.com", "name": "Иван"}'
```

**Получить пользователя по ID**
```bash
curl -X GET http://localhost:8080/users/1 -H "Accept: application/json"
```

**Обновить пользователя**
```bash
curl -X PUT http://localhost:8080/users/1 \
  -H "Content-Type: application/json" \
  -d '{"name": "Иван Иванов"}'
```

**Удалить пользователя**
```bash
curl -i -X DELETE http://localhost:8080/users/1
```

---

## Примеры ошибок

| Ситуация | Статус | Ответ |
|----------|--------|-------|
| Неверный ID (не число) | 400 Bad Request | `Неверный ID` |
| Задача/пользователь не найдена | 404 Not Found | `Задача не найдена` / `Пользователь не найден` |
| Невалидный JSON | 400 Bad Request | `Невалидный запрос` |

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

## Лицензия

**MIT**

Автор: Когут И.В.

Репозиторий: https://github.com/WAKAJAKA/ai-assist-it.git
