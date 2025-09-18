# 🇷🇺 Gosuslugi API

![CI](https://github.com/pavel97go/gosuslugi/actions/workflows/ci.yml/badge.svg)

Микросервис для подачи заявок (паспорта, справки и т.д.), написанный на **Go (Fiber + pgx/pgxpool)** с поддержкой миграций через **goose** и запуском в **Docker Compose**.  
Проект реализует базовый CRUD и готов для продакшена.

---

## 🚀 Быстрый старт

### 1. Запуск через Docker Compose
```bash
docker compose up -d --build
```

### 2. Проверка статуса
```bash
curl http://localhost:8080/health/live
curl http://localhost:8080/health/ready
```

---

## 📑 API эндпоинты

### Applications
- **POST** `/v1/applications` — создать заявку  
- **GET** `/v1/applications/:id` — получить заявку по ID  
- **GET** `/v1/applications?status=&document_type=&limit=&offset=` — список заявок с фильтрами  
- **PUT** `/v1/applications/:id` — обновить заявку  
- **DELETE** `/v1/applications/:id` — удалить заявку  

Пример POST-запроса:
```json
POST /v1/applications
{
  "citizen_name": "Иван Иванов",
  "document_type": "passport",
  "data": { "series": "1234", "number": "567890" }
}
```

---

## 🗄️ БД

Миграции выполняются с помощью [goose](https://github.com/pressly/goose):

```bash
goose -dir db/migrations postgres "postgres://postgres:postgres@localhost:5432/gosuslugi?sslmode=disable" up
```

---

## 🛠️ Технологии

- [Go](https://go.dev/) **1.25**
- [Fiber](https://gofiber.io/) — веб-фреймворк
- [pgx/pgxpool](https://github.com/jackc/pgx) — драйвер PostgreSQL
- [goose](https://github.com/pressly/goose) — миграции
- [zerolog](https://github.com/rs/zerolog) — логирование
- [Docker & Compose](https://docs.docker.com/) — инфраструктура
- [GitHub Actions](https://docs.github.com/en/actions) — CI/CD

---

##  Тесты
Запуск юнит-тестов:
```bash
go test -race ./...
```

##  Сборка и запуск вручную
```bash
go build ./cmd/api
./api
```

## 📄 Лицензия
Проект распространяется под лицензией [MIT](LICENSE).
