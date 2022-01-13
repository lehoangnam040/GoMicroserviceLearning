package router

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

func InitSwaggerRoutes(app *fiber.App) {
	routes := app.Group("/swagger")

	routes.Get("*", swagger.Handler)
}
