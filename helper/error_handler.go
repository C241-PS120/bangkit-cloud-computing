package helper

import (
	"context"
	"errors"

	"github.com/gofiber/fiber/v2"
)

func HandleRequestError(err error) *fiber.Error {
	if errors.Is(err, context.Canceled) {
		return fiber.NewError(fiber.StatusRequestTimeout, "Request was canceled by the client.")
	} else if errors.Is(err, context.DeadlineExceeded) {
		return fiber.NewError(fiber.StatusRequestTimeout, "Request timed out. Please try again.")
	} else {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
}
