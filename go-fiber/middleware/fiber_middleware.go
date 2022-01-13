package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func InitFiberBuiltinMiddleware(app *fiber.App) {
	app.Use(
		cors.New(),
		logger.New(),
	)
}
