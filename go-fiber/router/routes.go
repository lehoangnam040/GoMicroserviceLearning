package router

import (
	"go-fiber/controller"

	"github.com/gofiber/fiber/v2"
)

func InitRoutes(app *fiber.App) {

	apiV1 := app.Group("/v1")

	products := apiV1.Group("/products")
	products.Post("/", controller.CreateProduct)
	products.Get("/", controller.GetAllProduct)
	products.Get("/:id", controller.GetOneProduct)
	products.Delete("/:id", controller.DeleteProduct)

	categories := apiV1.Group("/categories")
	categories.Post("/", controller.CreateCategory)
	categories.Get("/", controller.GetAllCategory)
	categories.Get("/:id", controller.GetOneCategory)
	categories.Delete("/:id", controller.DeleteCategory)
}
