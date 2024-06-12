package handler

import (
	"github.com/C241-PS120/bangkit-cloud-computing/dto"
	"github.com/gofiber/fiber/v2"
)

func NotFoundHandler(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotFound).JSON(dto.Envelope{
		Success: false,
		Message: "The requested resource or endpoint is not found",
	})
}
