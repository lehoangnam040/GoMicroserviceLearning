package controller

import "github.com/gofiber/fiber/v2"

func CreateProduct(c *fiber.Ctx) error {
	return nil
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
