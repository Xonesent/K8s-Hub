package errlst

import "github.com/gofiber/fiber/v2"

// nolint:gochecknoglobals // cannot be const
var (
	HttpErrInvalidRequest = fiber.NewError(fiber.StatusBadRequest, fiber.ErrBadRequest.Error())
	HttpErrNotFound       = fiber.NewError(fiber.StatusNotFound, fiber.ErrNotFound.Error())
	HttpServerError       = fiber.NewError(fiber.StatusInternalServerError, fiber.ErrInternalServerError.Error())
)
