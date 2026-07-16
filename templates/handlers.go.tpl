package handlers

import (
	"runtime"
	"time"

	"github.com/gofiber/fiber/v2"
)

// startTime records when the server started (for uptime calculation).
var startTime = time.Now()

// Welcome handles the root API endpoint.
func Welcome(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status":  "ok",
		"message": "Welcome to {{PROJECT_NAME}} API",
		"version": "v1.0.0",
	})
}

// HealthCheck returns the current health status of the application.
func HealthCheck(c *fiber.Ctx) error {
	uptime := time.Since(startTime)

	return c.JSON(fiber.Map{
		"status":  "ok",
		"uptime":  uptime.String(),
		"go":      runtime.Version(),
		"os":      runtime.GOOS,
		"goroutines": runtime.NumGoroutine(),
	})
}
