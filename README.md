# Task Manager API

Учебный проект на Go (Gin), предоставляющий REST API для управления задачами (to-do list).

## 🚀 Возможности

- **User API**:
  - `GET /user/tasks` — получить список задач
  - `POST /user/task` — добавить задачу (параметр `title`)
  - `PUT /user/task/:id` — отметить задачу выполненной
  - `DELETE /user/task/:id` — удалить задачу

- **Admin API**:
  - `GET /admin/stats` — статистика (количество задач)
  - `POST /admin/reset` — очистить все задачи

## 🛠️ Запуск

1. Клонируй проект или скопируй файлы.
2. Установи зависимости:
   ```bash
   go mod tidy
