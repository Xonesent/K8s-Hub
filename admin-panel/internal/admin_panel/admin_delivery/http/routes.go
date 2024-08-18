package http_admin

import (
	"github.com/Xonesent/K8s-Hub/admin-panel/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

type AdminHDL interface {
	GetUserStatistics() fiber.Handler
}

func MapAdminRoutes(group fiber.Router, h AdminHDL, mw *middleware.ManagerMiddleware) {
	group.Get("/user_stats/:tg_id", mw.NoMiddleware(), h.GetUserStatistics())
}
