package main

import (
	"fmt"
	"go-fiber/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
)

func main() {

	fmt.Println("hello")

	app := fiber.New()

	router.InitRoutes(app)
	app.Use(cors.New())
	app.Use(csrf.New())

	app.Listen(":3000")
}
