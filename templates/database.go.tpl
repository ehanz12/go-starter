package database

import (
	"fmt"
	"log"
	"time"

	"{{MODULE_NAME}}/config"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	{{DB_IMPORT}}
)

// DB is the global GORM database instance.
var DB *gorm.DB

// Connect initialises the database connection pool.
// It panics if the connection cannot be established.
func Connect() {
	cfg := config.AppConfig

	gormCfg := &gorm.Config{}
	if cfg.Debug {
		gormCfg.Logger = logger.Default.LogMode(logger.Info)
	} else {
		gormCfg.Logger = logger.Default.LogMode(logger.Silent)
	}

	db, err := gorm.Open({{DB_DIAL}}, gormCfg)
	if err != nil {
		log.Fatalf("❌ Failed to connect to database: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("❌ Failed to get underlying *sql.DB: %v", err)
	}

	// Connection pool settings
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	// Verify connection
	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("❌ Database ping failed: %v", err)
	}

	DB = db
	fmt.Println("✅ Database connected successfully")
}
