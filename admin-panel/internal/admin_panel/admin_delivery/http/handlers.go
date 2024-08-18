package http_admin

import (
	"strconv"

	"github.com/Xonesent/K8s-Hub/admin-panel/config"
	models "github.com/Xonesent/K8s-Hub/admin-panel/internal/business_models"
	errlst "github.com/Xonesent/K8s-Hub/admin-panel/pkg/predefined_responses/error_list"
	"github.com/gofiber/fiber/v2"
)

type AdminHandler struct {
	cfg     *config.Config
	adminUC AdminUC
}

func NewAdminHandler(cfg *config.Config, adminUC AdminUC) *AdminHandler {
	return &AdminHandler{
		cfg:     cfg,
		adminUC: adminUC,
	}
}

func (h *AdminHandler) GetUserStatistics() fiber.Handler {
	return func(c *fiber.Ctx) error {
		tgId, err := strconv.Atoi(c.Params("tg_id"))
		if err != nil {
			return errlst.HttpErrInvalidRequest
		}

		userStatistics, err := h.adminUC.GetUserStatistics(c.Context(), models.TgId(tgId))
		if err != nil {
			return err
		}

		return c.JSON(toUserStatistics(userStatistics))
	}
}
