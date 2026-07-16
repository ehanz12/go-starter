package generator

// DBConfig holds database-specific template values for each supported driver.
type DBConfig struct {
	Driver      string
	Port        string
	Import      string
	Dial        string
	DockerImage string
	DockerEnv   string
	DockerData  string
	HealthCheck string
}

// SupportedDatabases maps a normalised driver name to its config.
var SupportedDatabases = map[string]DBConfig{
	"mysql": {
		Driver: "mysql",
		Port:   "3306",
		Import: `"gorm.io/driver/mysql"`,
		Dial: `mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName,
	))`,
		DockerImage: "mysql:8.0",
		DockerEnv: `MYSQL_ROOT_PASSWORD: "${DB_PASSWORD}"
      MYSQL_DATABASE: "${DB_NAME}"
      MYSQL_USER: "${DB_USER}"
      MYSQL_PASSWORD: "${DB_PASSWORD}"`,
		DockerData:  "mysql",
		HealthCheck: `["CMD", "mysqladmin", "ping", "-h", "localhost"]`,
	},
	"postgres": {
		Driver: "postgres",
		Port:   "5432",
		Import: `"gorm.io/driver/postgres"`,
		Dial: `postgres.Open(fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=UTC",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort, cfg.DBSSLMode,
	))`,
		DockerImage: "postgres:16-alpine",
		DockerEnv: `POSTGRES_USER: "${DB_USER}"
      POSTGRES_PASSWORD: "${DB_PASSWORD}"
      POSTGRES_DB: "${DB_NAME}"`,
		DockerData:  "postgresql",
		HealthCheck: `["CMD-SHELL", "pg_isready -U ${DB_USER} -d ${DB_NAME}"]`,
	},
	"sqlite": {
		Driver: "sqlite",
		Port:   "0",
		Import: `"gorm.io/driver/sqlite"`,
		Dial:   `sqlite.Open(cfg.DBName + ".db")`,
		// SQLite doesn't need a Docker service
		DockerImage: "",
		DockerEnv:   "",
		DockerData:  "",
		HealthCheck: "",
	},
}

// NormaliseDB converts user-facing database names (e.g. "PostgreSQL") to a canonical key.
func NormaliseDB(choice string) string {
	switch choice {
	case "PostgreSQL", "postgres", "postgresql":
		return "postgres"
	case "SQLite", "sqlite":
		return "sqlite"
	default:
		return "mysql"
	}
}
