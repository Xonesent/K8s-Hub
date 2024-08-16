package fiberApp

import (
	"fmt"
	"runtime/debug"

	"github.com/Xonesent/K8s-Hub/admin-panel/internal/error_handler"
	"github.com/Xonesent/K8s-Hub/admin-panel/pkg/utilities/go_utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"go.uber.org/zap"
)

type FiberConfig struct {
	Host string `yaml:"Host" validate:"required"`
	Port string `yaml:"Port" validate:"required"`
}

func NewFiberClient() *fiber.App {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
		ErrorHandler:          error_handler.FiberErrorHandler,
	})

	app.Use(func(c *fiber.Ctx) error {
		reqHeadersStr := fmt.Sprintf("%v", c.GetReqHeaders())
		zap.L().Info(
			"Request received",
			zap.String("method", c.Method()),
			zap.String("route", c.Path()),
			zap.String("ip", c.IP()),
			zap.String("headers", reqHeadersStr[4:len(reqHeadersStr)-1]),
		)

		return c.Next()
	})

	app.Use(recover.New(recover.Config{
		EnableStackTrace: true,
		StackTraceHandler: func(c *fiber.Ctx, e interface{}) {
			fullStackTrace := string(debug.Stack())
			zap.L().Error("Recovered from panic: " + go_utils.LimitStackTrace(fullStackTrace, 1))
		},
	}))

	return app
}
