package handler

import "github.com/gofiber/fiber/v2"

func NotFoundHandler(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"success": false,
		"message": "The requested resource or endpoint is not found.",
	})
}
