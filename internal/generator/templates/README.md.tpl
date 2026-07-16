# {{PROJECT_NAME}}

> A production-ready Go REST API built with [Fiber](https://gofiber.io/) and [GORM](https://gorm.io/).

---

## 📋 Tech Stack

| Layer      | Technology                   |
|------------|------------------------------|
| Framework  | [Fiber v2](https://gofiber.io/) |
| ORM        | [GORM](https://gorm.io/)    |
| Database   | {{DB_DRIVER}}                |
| Auth       | JWT (golang-jwt)             |
| Logger     | Zerolog                      |
| Config     | godotenv                     |
| Container  | Docker + Compose             |

---

## 🚀 Getting Started

### Prerequisites

- Go 1.21+
- Docker & Docker Compose (optional)
- `make` (optional but recommended)

### 1. Clone & enter the project

```bash
cd {{PROJECT_NAME}}
```

### 2. Configure environment

```bash
cp .env.example .env   # edit values as needed
```

### 3. Run locally

```bash
make run        # uses Air for hot-reload
# or
go run ./cmd/api/main.go
```

### 4. Run with Docker

```bash
make docker-up
```

---

## 📦 Project Structure

```
{{PROJECT_NAME}}/
├── cmd/api/          # Entry point
├── config/           # Environment & app config
├── database/         # DB connection & migrations
├── internal/
│   ├── handlers/     # HTTP handlers
│   ├── middleware/   # Fiber middleware (JWT, Logger…)
│   ├── models/       # GORM models
│   ├── repositories/ # Data access layer
│   ├── requests/     # Request DTOs
│   ├── responses/    # Response DTOs
│   ├── routes/       # Route registration
│   ├── services/     # Business logic
│   ├── utils/        # Shared helpers
│   └── validators/   # Input validation
├── docs/             # Swagger / API docs
├── logs/             # Log files (gitignored)
├── storage/          # Uploaded files (gitignored)
├── scripts/          # DB seed / helper scripts
├── pkg/              # Reusable packages
├── Makefile
├── Dockerfile
├── docker-compose.yml
└── .env
```

---

## 🛠️ Available Commands

```bash
make run          # Hot-reload dev server
make build        # Compile production binary → bin/{{PROJECT_NAME}}
make test         # Run tests with race detector
make lint         # Run golangci-lint
make tidy         # go mod tidy + verify
make docker-up    # Start containers
make docker-down  # Stop containers
make migrate      # Run DB migrations
make help         # Show all targets
```

---

## 🔒 Authentication

JWT is used for authentication. Obtain a token via `POST /api/v1/auth/login` and include it in the `Authorization: Bearer <token>` header on protected routes.

---

## 🌐 API Endpoints

| Method | Path             | Description        |
|--------|------------------|--------------------|
| GET    | /health          | Health check       |
| GET    | /api/v1/         | Welcome message    |

---

## 📄 License

MIT © {{PROJECT_NAME}}
