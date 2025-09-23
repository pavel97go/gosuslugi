# 🇷🇺 Gosuslugi API
> From bureaucracy → to API 🚀

![CI](https://github.com/pavel97go/gosuslugi/actions/workflows/ci.yml/badge.svg)
![Go Version](https://img.shields.io/badge/Go-1.25-blue)
![Fiber](https://img.shields.io/badge/Fiber-🚀-green)
![Postgres](https://img.shields.io/badge/Postgres-16-blue)
![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)


Микросервис для подачи заявок (**паспорта, справки и т.д.**) на **Go (Fiber + pgx/pgxpool)**.  
Поддерживает миграции через **goose**, логирование через **zerolog**, и готов к запуску в **Docker Compose**.  
Проект реализует полный CRUD и настроен для CI/CD через GitHub Actions.

> ⚠️ Это **pet-project**: учебная реализация модели госуслуг для демонстрации навыков разработки.  
> В реальном продакшне «Госуслуг» используется гораздо более сложная архитектура и стек технологий.

## 🏛️ Актуальность и практическое применение в условиях современного государства

В современных условиях цифровизации государственные сервисы должны быть:
- **Доступными онлайн** — минимизация очередей и необходимости личного визита в ведомства.  
- **Единообразными** — единый API для интеграции с другими государственными и муниципальными системами.  
- **Прозрачными** — заявитель может отслеживать статус заявки в реальном времени.  
- **Масштабируемыми** — архитектура сервиса позволяет адаптироваться к высоким нагрузкам.  
- **Безопасными** — используется проверенный стек технологий и централизованная авторизация.

Данный микросервис демонстрирует, как можно реализовать **универсальный механизм подачи и обработки заявлений граждан**.  
Он может быть встроен в экосистему государственных услуг (портал «Госуслуги») и применяться для:
- подачи заявлений на выдачу паспорта,  
- получения справок, сертификатов,  
- записи на приём в ведомства,  
- других электронных сервисов.  

Таким образом, проект показывает **практическое решение**, которое отвечает ключевым требованиям цифрового государства: **эффективность, прозрачность, удобство и доверие граждан**.

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
curl http://localhost:8080/health/ready  # {"status":"up","version":"v1.0.1"}
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
