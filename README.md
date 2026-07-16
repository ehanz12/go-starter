# go-starter 🚀

> A professional CLI tool that scaffolds production-ready Go REST API projects in seconds.

---

## ✨ Features

| Feature | Description |
|---|---|
| 🗄️ **Multi-database** | MySQL, PostgreSQL, SQLite — choose during setup |
| 🎨 **Interactive TUI** | Beautiful survey wizard powered by `AlecAivazis/survey` |
| 🔐 **JWT Auth** | Optional JWT middleware with role-based access control |
| 📋 **Zerolog Logger** | Structured JSON logging with pretty dev console |
| 🐳 **Docker** | Multi-stage Dockerfile + docker-compose with DB health checks |
| 🛠️ **Makefile** | `run`, `build`, `test`, `lint`, `docker-up`, `migrate` targets |
| 📖 **README** | Auto-generated README with project structure & commands |
| 🌿 **Git Init** | Optional git repository initialization |
| ♻️ **Graceful Shutdown** | OS signal handling for clean server shutdown |
| 🏥 **Health Check** | `/health` endpoint with uptime & runtime metrics |

---

## 📦 Installation

```bash
git clone https://github.com/ehanz12/go-starter
cd go-starter
go build -o go-starter .
```

Or install directly:

```bash
go install github.com/ehanz12/go-starter@latest
```

---

## 🚀 Usage

### Interactive wizard (recommended)

```bash
go-starter new my-api
```

This launches a TUI where you select:
- Module name
- Database driver (MySQL / PostgreSQL / SQLite)
- Optional features (JWT, Docker, Logger, Makefile, README)
- Whether to initialise a Git repository

### With flags (skip wizard questions)

```bash
# Specify database driver upfront
go-starter new my-api --db postgres

# Full example with all flags
go-starter new my-api \
  --module github.com/yourusername/my-api \
  --db postgres \
  --git
```

### List available databases & features

```bash
go-starter list
```

---

## 📁 Generated Project Structure

```
my-api/
├── cmd/api/          # Application entry point (main.go)
├── config/           # Config struct loaded from .env
├── database/         # DB connection + migrations folder
├── internal/
│   ├── handlers/     # HTTP handlers (Welcome + HealthCheck)
│   ├── middleware/   # JWT auth, Logger (if selected)
│   ├── models/       # GORM models
│   ├── repositories/ # Data access layer
│   ├── requests/     # Request DTOs / input structs
│   ├── responses/    # Response DTOs
│   ├── routes/       # Route registration + error handler
│   ├── services/     # Business logic layer
│   ├── utils/        # Shared helpers
│   └── validators/   # Input validation
├── docs/             # Swagger / API documentation
├── logs/             # Log files (git-ignored)
├── storage/          # Uploaded files (git-ignored)
├── scripts/          # DB seed scripts
├── pkg/              # Reusable packages
├── Dockerfile        # (if Docker selected)
├── docker-compose.yml # (if Docker selected)
├── Makefile          # (if Makefile selected)
├── README.md         # (if README selected)
├── .env              # Environment variables
└── .env.example      # Environment template (safe to commit)
```

---

## 🛠️ CLI Reference

```
go-starter [command]

Commands:
  new     Create a new Go project (interactive wizard)
  list    List available databases and optional features

Flags for `new`:
  -m, --module   Module name (default: github.com/yourusername/<name>)
  -d, --db       Database driver: mysql | postgres | sqlite
      --git      Initialize a Git repository
```

---

## 🤝 Contributing

PRs welcome! Please open an issue first to discuss what you'd like to change.

---

## 📄 License

MIT © ehanz12
