package routes

import (
	"time"

	"{{MODULE_NAME}}/internal/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

// ErrorHandler is a custom global error handler for Fiber.
func ErrorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	msg := "Internal Server Error"

	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
		msg = e.Message
	}

	return c.Status(code).JSON(fiber.Map{
		"status":  "error",
		"message": msg,
	})
}

// SetupRoutes registers all application routes.
func SetupRoutes(app *fiber.App) {
	// Rate limiter: 100 req/min per IP
	apiLimiter := limiter.New(limiter.Config{
		Max:        100,
		Expiration: 1 * time.Minute,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP()
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"status":  "error",
				"message": "Too many requests. Please try again later.",
			})
		},
	})

	// Health check (no rate limit)
	app.Get("/health", handlers.HealthCheck)

	// API v1 group
	api := app.Group("/api/v1", apiLimiter)

	// Public routes
	api.Get("/", handlers.Welcome)

	// ===========================================================
	// 📌 Add your routes below
	// ===========================================================
	// Example:
	// auth := api.Group("/auth")
	// auth.Post("/register", handlers.Register)
	// auth.Post("/login", handlers.Login)
	//
	// protected := api.Group("/", middleware.JWTProtected())
	// protected.Get("/profile", handlers.GetProfile)
	// ===========================================================
}
