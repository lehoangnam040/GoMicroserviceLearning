package controller

import "github.com/gofiber/fiber/v2"

func CreateCategory(c *fiber.Ctx) error {
	return nil
}

func GetAllCategory(c *fiber.Ctx) error {
	return c.SendString("all categories")
}

func GetOneCategory(c *fiber.Ctx) error {
	return c.SendString("one categories")
}

func DeleteCategory(c *fiber.Ctx) error {
	return nil
}
