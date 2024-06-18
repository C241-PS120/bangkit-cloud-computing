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

func GetIdFromRequest(ctx *fiber.Ctx) (int, error) {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return 0, fiber.NewError(fiber.StatusBadRequest, "id is not valid, must be an integer")
	}
	if id < 1 {
		return 0, fiber.NewError(fiber.StatusBadRequest, "id is not valid, must be greater than 0")
	}
	return id, nil
}
