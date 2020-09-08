package mw

import (
	"github.com/gofiber/fiber"
	jwtware "github.com/gofiber/jwt"
)

// Protected protect routes
func Protected() func(*fiber.Ctx) {
	return jwtware.New(jwtware.Config{
		SigningKey:   []byte("secret"),
		ErrorHandler: jwtError,
	})
}

func jwtError(c *fiber.Ctx, err error) {
	if err.Error() == "Missing or malformed JWT" {
		_ = c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "5001", "message": "Missing or malformed JWT", "data": nil})
	} else {
		_ = c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "5002", "message": "Invalid or expired JWT", "data": nil})
	}
}
