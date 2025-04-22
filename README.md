# Task Management Service 📋

Это приложение для управления задачами, использующее архитектуру с разделением на слои (`Hexagonal Architecture`). Оно предоставляет возможности для создания и получения задач через `HTTP API`.

## Описание API 📃

Приложение предоставляет следующие эндпоинты для работы с задачами:

`POST /tasks` — Создание новой задачи.

`GET /tasks/{id}` — Получение задачи по ID.

## Структура проекта 🗂️

- `cmd/` - Точка входа в приложение.
- `internal/` - Внутренние пакеты (сервисы, репозитории и т.д.).
  - `domain/` - Модели данных и бизнес-логика.
  - `ports/in/` - Интерфейсы входных портов.
  - `ports/out/` - Интерфейсы выходных портов (репозитории).
  - `service/` - Реализация бизнес-логики.
  - `adapters/repository/` - Реализация репозитория для работы с базой данных.
  - `adapters/transport/` - Обработчики HTTP запросов и роутеры.
- `Dockerfile` - Docker-файл для сборки и запуска приложения.
- `docker-compose.yml` - Docker Compose конфигурация для запуска приложения и базы данных.

## Сборка и запуск приложения 🚀

### 1. Сборка с использованием Docker 🐋

Для того чтобы собрать и запустить приложение с использованием `Docker`, выполните следующие шаги:

1. Убедитесь, что у вас установлен `Docker`.

2. Для удобства запуска приложения с базой данных `PostgreSQL` используйте `Docker Compose`. Для этого создайте файл docker-compose.yml со следующим содержанием:

```yaml
services:
  postgres:
    image: postgres:15
    environment:
      POSTGRES_USER: your_user
      POSTGRES_PASSWORD: your_password
      POSTGRES_DB: your_db
    ports:
      - '5432:5432'
    volumes:
      - postgres_data:/var/lib/postgresql/data

  app:
    build:
      context: .
      dockerfile: Dockerfile
    depends_on:
      - postgres
    ports:
      - '8080:8080'
    environment:
      POSTGRES_USER: your_user
      POSTGRES_PASSWORD: your_password
      POSTGRES_DB: your_db
      POSTGRES_HOST: postgres

volumes:
  postgres_data:
```

3. Соберите и запустите контейнеры:

```bash
docker-compose up --build
```

## Автор ✒️

Дамир Усетов
