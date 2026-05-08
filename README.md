# List API


---

## Структура проекта

После реализации проект будет иметь следующую структуру:
tasks-api/
├── cmd/server/main.go
├── internal/handlers/tasks.go
├── internal/models/task.go
├── internal/storage/memory.go
├── internal/storage/storage.go
├── go.mod
└── README.md

REST API для управления списком задач (in‑memory). Лёгкий сервис для тестирования и прототипирования без необходимости настройки базы данных.

## Технологии

* Язык: Go (версия 1.21+)
* Веб‑фреймворк: стандартный `net/http`
* Хранение: in‑memory (синхронизация через `sync.RWMutex`)
* Формат данных: JSON
* Логирование: стандартный пакет `log`


## Запуск проекта


### Предварительные требования

* Установленный Go (версия 1.21 или выше)
* Терминал/командная строка

### Инструкции по запуску

1. Клонируйте репозиторий (если проект не локальный):
   ```bash
   git clone <url-репозитория>
   cd tasks-api
2. Инициализируйте модуль Go:
   ```bash
   go mod init tasks-api
3. Запустите сервер:
   ```bash
   go run cmd/server/main.go
## Эндпоинты API
| Метод | Путь | Описание | Обязательные поля | Успешный статус |
|-------|------|----------|------------------|---------------|
| GET | `/tasks` | Получить список всех задач | — | 200 OK |
| POST | `/tasks` | Создать новую задачу | `title` | 201 Created |
| GET | `/tasks/{id}` | Получить задачу по ID | — | 200 OK |
| PUT | `/tasks/{id}` | Обновить задачу целиком | `title` | 200 OK |
| DELETE | `/tasks/{id}` | Удалить задачу | — | 204 No Content |
| GET | `/health` | Health check сервиса | — | 200 OK |

## Примеры curl
1. Получить все задачи
   ```bash
   curl -X GET http://localhost:8080/tasks
2. Создать задачу
    ```bash
    curl -X POST http://localhost:8080/tasks \
3. Получить задачу по ID
   ```bash
   curl -X GET http://localhost:8080/tasks/1
4. Обновить задачу
   ```bash
   curl -X PUT http://localhost:8080/tasks/1 \
5. Удалить задачу
   ```bash
   curl -X DELETE http://localhost:8080/tasks/1
6. Health check
   ```bash
   curl -X GET http://localhost:8080/health


