package middleware

import (
	"github.com/Xonesent/K8s-Hub/admin-panel/config"
	"github.com/gofiber/fiber/v2"
)

type ManagerMiddleware struct {
	cfg *config.Config
}

func NewMiddlewareManager(cfg *config.Config) *ManagerMiddleware {
	return &ManagerMiddleware{cfg: cfg}
}

func (m *ManagerMiddleware) NoMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Next()
	}
}
