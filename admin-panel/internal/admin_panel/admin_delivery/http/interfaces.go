package http_admin

import (
	"context"

	models "github.com/Xonesent/K8s-Hub/admin-panel/internal/business_models"
)

type AdminUC interface {
	GetUserStatistics(ctx context.Context, tgId models.TgId) (models.UserLog, error)
}
