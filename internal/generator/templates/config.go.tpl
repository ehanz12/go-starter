package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config holds all application configuration loaded from environment variables.
type Config struct {
	// App
	AppName string
	AppEnv  string
	Port    string
	Debug   bool

	// Database
	DBDriver   string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string

	// JWT
	JWTSecret          string
	JWTExpiryHours     int
	JWTRefreshHours    int
}

// AppConfig is the global singleton configuration.
var AppConfig *Config

// LoadEnv reads the .env file and populates AppConfig.
// Panics on critical failures.
func LoadEnv() {
	// In production, environment variables may be injected directly.
	if err := godotenv.Load(); err != nil {
		fmt.Println("⚠️  .env file not found, using environment variables")
	}

	debug, _ := strconv.ParseBool(getEnv("APP_DEBUG", "false"))
	jwtExpiry, _ := strconv.Atoi(getEnv("JWT_EXPIRY_HOURS", "24"))
	jwtRefresh, _ := strconv.Atoi(getEnv("JWT_REFRESH_HOURS", "168"))

	AppConfig = &Config{
		AppName: getEnv("APP_NAME", "{{PROJECT_NAME}}"),
		AppEnv:  getEnv("APP_ENV", "development"),
		Port:    getEnv("PORT", "8080"),
		Debug:   debug,

		DBDriver:   getEnv("DB_DRIVER", "{{DB_DRIVER}}"),
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "{{DB_PORT}}"),
		DBUser:     getEnv("DB_USER", "root"),
		DBPassword: getEnv("DB_PASSWORD", ""),
		DBName:     getEnv("DB_NAME", "{{PROJECT_NAME}}_db"),
		DBSSLMode:  getEnv("DB_SSLMODE", "disable"),

		JWTSecret:       getEnv("JWT_SECRET", "change-me-in-production"),
		JWTExpiryHours:  jwtExpiry,
		JWTRefreshHours: jwtRefresh,
	}
}

// getEnv returns the environment variable or a fallback default.
func getEnv(key, defaultVal string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defaultVal
}