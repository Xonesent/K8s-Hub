package error_handler

import (
	"database/sql"
	"errors"

	"github.com/Xonesent/K8s-Hub/admin-panel/pkg/constant"
	"github.com/Xonesent/K8s-Hub/admin-panel/pkg/utilities/go_utils"
	"github.com/gofiber/fiber/v2"
)

func FiberErrorHandler(ctx *fiber.Ctx, err error) error {
	setStatusCode(ctx, err)

	if go_utils.InStringSlice(constant.Host, constant.DevHosts) {
		return ctx.JSON(fiber.Map{
			"error": err.Error(),
			"data":  nil,
		})
	}

	return ctx.JSON(fiber.Map{
		"data": nil,
	})
}

func setStatusCode(ctx *fiber.Ctx, err error) {
	statusCode := fiber.StatusInternalServerError

	var e *fiber.Error
	if errors.As(err, &e) {
		statusCode = e.Code
	}

	if errors.Is(err, sql.ErrNoRows) {
		statusCode = fiber.StatusNotFound
	}

	ctx.Status(statusCode)
}
