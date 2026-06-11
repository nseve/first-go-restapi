# Project Manager

Full Stack приложение для управления проектами и задачами.

## Возможности

### Аутентификация

* Регистрация пользователей
* Авторизация пользователей
* JWT аутентификация
* Хэширование паролей через bcrypt

### Проекты

* Создание проекта
* Получение списка проектов пользователя
* Получение проекта по ID
* Обновление проекта
* Удаление проекта

### Задачи

* Создание задачи внутри проекта
* Получение списка задач проекта
* Получение задачи по ID
* Обновление задачи
* Удаление задачи

### Ownership

Каждый пользователь имеет доступ только к своим проектам и задачам.

## Технологии

### Backend

* Go
* Chi Router
* GORM
* PostgreSQL
* JWT
* bcrypt
* godotenv

### Frontend

* React
* Vite
* React Router
* Fetch API
* CSS

### Infrastructure

* Docker
* Docker Compose
* Nginx

---

## Архитектура

Backend построен по многослойной архитектуре:

```text
Handler
  ↓
Service
  ↓
Repository
  ↓
PostgreSQL
```

### Handler

Обработка HTTP запросов и ответов.

### Service

Бизнес-логика и валидация данных.

### Repository

Работа с базой данных через GORM.

### Middleware

JWT авторизация и обработка запросов.

---

## Структура проекта

```text
.
├── cmd
│   └── app
│       └── main.go
│
├── internal
│   ├── db
│   ├── handler
│   ├── middleware
│   ├── models
│   ├── repository
│   ├── response
│   └── service
│
├── frontend
│   ├── src
│   ├── Dockerfile
│   └── nginx.conf
│
├── Dockerfile
├── docker-compose.dev.yml
├── docker-compose.yml
└── README.md
```

---

## Запуск через Docker

Сборка и запуск:

```bash
docker compose up --build
```

После запуска:

Backend:

```text
http://localhost:8080
```

Frontend:

```text
http://localhost:3000
```

PostgreSQL:

```text
localhost:5433
```

---

## Переменные окружения

Пример .env представлен в файле .env.example

---

## API

### Регистрация

POST

```text
/auth/register
```

Body:

```json
{
  "email": "user@example.com",
  "password": "password123"
}
```

---

### Авторизация

POST

```text
/auth/login
```

Body:

```json
{
  "email": "user@example.com",
  "password": "password123"
}
```

Response:

```json
{
  "token": "jwt_token"
}
```

---

## Защищённые маршруты

Для всех защищённых маршрутов необходимо передавать JWT токен:

```http
Authorization: Bearer <token>
```

---

### Проекты

Получить проекты:

```text
GET /projects
```

Создать проект:

```text
POST /projects
```

Получить проект:

```text
GET /projects/{id}
```

Обновить проект:

```text
PUT /projects/{id}
```

Удалить проект:

```text
DELETE /projects/{id}
```

---

### Задачи

Получить задачи проекта:

```text
GET /projects/{projectId}/tasks
```

Создать задачу:

```text
POST /projects/{projectId}/tasks
```

Получить задачу:

```text
GET /tasks/{id}
```

Обновить задачу:

```text
PUT /tasks/{id}
```

Удалить задачу:

```text
DELETE /tasks/{id}
```

---

## Frontend

Frontend предоставляет:

* регистрацию
* авторизацию
* управление проектами
* управление задачами
* поиск по ID
* inline редактирование
* выбор проекта и задачи
* защищённый доступ через JWT