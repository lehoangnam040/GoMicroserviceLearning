package middleware

import (
	"os"

	"github.com/gofiber/fiber/v2"
	gofiberJwtV3 "github.com/gofiber/jwt/v3"
)

func ValidateJwt() func(*fiber.Ctx) error {
	return gofiberJwtV3.New(gofiberJwtV3.Config{
		SigningKey:   []byte(os.Getenv("JWT_SECRET_KEY")),
		ErrorHandler: jwtError,
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"code":    "1000",
		"message": err.Error(),
	})
}
