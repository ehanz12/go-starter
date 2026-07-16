package middleware

import (
	"strings"

	"{{MODULE_NAME}}/config"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/golang-jwt/jwt/v5"
)

// JWTClaims defines the structure of claims inside a JWT token.
type JWTClaims struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

// JWTProtected returns a middleware that enforces JWT authentication.
func JWTProtected() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:  jwtware.SigningKey{Key: []byte(config.AppConfig.JWTSecret)},
		ContextKey: "user",
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"status":  "error",
				"message": "Unauthorized: " + err.Error(),
			})
		},
		Filter: func(c *fiber.Ctx) bool {
			// Skip JWT for OPTIONS (preflight)
			return c.Method() == fiber.MethodOptions
		},
	})
}

// GetClaims extracts JWTClaims from the request context.
// Must be called inside a JWTProtected route.
func GetClaims(c *fiber.Ctx) *JWTClaims {
	token, ok := c.Locals("user").(*jwt.Token)
	if !ok {
		return nil
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil
	}
	return &JWTClaims{
		UserID: uint(claims["user_id"].(float64)),
		Email:  claims["email"].(string),
		Role:   claims["role"].(string),
	}
}

// RoleRequired ensures the authenticated user has one of the given roles.
func RoleRequired(roles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		claims := GetClaims(c)
		if claims == nil {
			return fiber.ErrUnauthorized
		}
		for _, r := range roles {
			if strings.EqualFold(claims.Role, r) {
				return c.Next()
			}
		}
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"status":  "error",
			"message": "Forbidden: insufficient permissions",
		})
	}
}
