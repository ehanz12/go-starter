package middleware

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/gofiber/fiber/v2"
)

// Log is the global logger instance.
var Log zerolog.Logger

// Init initialises zerolog with a pretty console writer for dev
// and a JSON writer for production.
func Init(env string) {
	zerolog.TimeFieldFormat = time.RFC3339

	if env == "development" {
		Log = zerolog.New(zerolog.ConsoleWriter{
			Out:        os.Stderr,
			TimeFormat: "15:04:05",
		}).With().Timestamp().Caller().Logger()
	} else {
		Log = zerolog.New(os.Stdout).With().Timestamp().Logger()
	}

	// Override the global log package
	log.Logger = Log
}

// Middleware returns a Fiber middleware that logs each HTTP request.
func Middleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()
		err := c.Next()
		duration := time.Since(start)

		status := c.Response().StatusCode()
		event := Log.Info()
		if status >= 500 {
			event = Log.Error()
		} else if status >= 400 {
			event = Log.Warn()
		}

		event.
			Str("method", c.Method()).
			Str("path", c.Path()).
			Int("status", status).
			Str("ip", c.IP()).
			Dur("latency", duration).
			Str("user_agent", c.Get("User-Agent")).
			Msg("request")

		return err
	}
}
