# 🇷🇺 Gosuslugi API

![CI](https://github.com/pavel97go/gosuslugi/actions/workflows/ci.yml/badge.svg)
![Go Version](https://img.shields.io/badge/Go-1.25-blue)
![Fiber](https://img.shields.io/badge/Fiber-🚀-green)
![Postgres](https://img.shields.io/badge/Postgres-16-blue)
![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)

Микросервис для подачи заявок (**паспорта, справки и т.д.**) на **Go (Fiber + pgx/pgxpool)**.  
Поддерживает миграции через **goose**, логирование через **zerolog**, и готов к запуску в **Docker Compose**.  
Проект реализует полный CRUD и настроен для CI/CD через GitHub Actions.

---

## 🚀 Быстрый старт

```bash
git clone git@github.com:pavel97go/gosuslugi.git
cd gosuslugi
docker compose up -d --build
```

Проверка статуса:
```bash
curl http://localhost:8080/health/live   # {"status":"ok"}
curl http://localhost:8080/health/ready  # {"status":"up"}
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

## 🗄️ Архитектура

```
[ Client ] -> [ Fiber API ] -> [ Usecase ] -> [ Repository ] -> [ PostgreSQL ]
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

## 🧪 Тесты
Запуск юнит-тестов:
```bash
go test -race ./...
```

---

## ⚡ Сборка и запуск вручную
```bash
go build ./cmd/api
./api
```

---

## 🗺️ Roadmap
- [ ] Добавить **Swagger/OpenAPI**
- [ ] Подключить **JWT авторизацию**
- [ ] Реализовать **Rate limiting**
- [ ] Добавить **Helm чарты** для Kubernetes
- [ ] Настроить **DockerHub автопубликацию**

---

## 📄 Лицензия
Проект распространяется под лицензией [MIT](LICENSE).
