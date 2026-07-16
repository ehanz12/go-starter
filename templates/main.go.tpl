package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"{{MODULE_NAME}}/config"
	"{{MODULE_NAME}}/database"
	"{{MODULE_NAME}}/internal/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	config.LoadEnv()
	database.Connect()

	app := fiber.New(fiber.Config{
		AppName:      "{{PROJECT_NAME}} v1.0.0",
		ErrorHandler: routes.ErrorHandler,
	})

	// Middleware global
	app.Use(recover.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,PUT,DELETE,PATCH,OPTIONS",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowCredentials: false,
	}))

	routes.SetupRoutes(app)

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := app.Listen(":" + config.AppConfig.Port); err != nil {
			log.Fatalf("Server error: %v", err)
		}
	}()

	log.Printf("🚀 {{PROJECT_NAME}} running on port %s", config.AppConfig.Port)
	<-quit
	log.Println("🛑 Shutting down server...")
	_ = app.Shutdown()
	log.Println("✅ Server exited cleanly")
}