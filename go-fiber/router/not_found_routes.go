package router

import "github.com/gofiber/fiber/v2"

func InitNotFoundRoutes(app *fiber.App) {
	app.Use(
		func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"code":    404,
				"message": "Not found",
			})
		},
	)
}
