# ============================================================
# {{PROJECT_NAME}}
# ============================================================
APP_NAME={{PROJECT_NAME}}
APP_ENV=development
APP_DEBUG=true
PORT=8080

# ---- Database -----------------------------------------------
DB_DRIVER={{DB_DRIVER}}
DB_HOST=localhost
DB_PORT={{DB_PORT}}
DB_USER=root
DB_PASSWORD=secret
DB_NAME={{PROJECT_NAME}}_db
DB_SSLMODE=disable

# ---- JWT ----------------------------------------------------
JWT_SECRET=change-me-to-a-strong-random-secret
JWT_EXPIRY_HOURS=24
JWT_REFRESH_HOURS=168