package controller

import "github.com/gofiber/fiber/v2"

func CreateProduct(c *fiber.Ctx) error {
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"code":    200,
		"message": "CREATED",
	})
}

func GetAllProduct(c *fiber.Ctx) error {
	return c.SendString("all products")
}

func GetOneProduct(c *fiber.Ctx) error {

	return c.SendString("one products " + c.Params("id"))
}

func DeleteProduct(c *fiber.Ctx) error {
	return nil
}
