package admin_usecase

import (
	"context"

	models "github.com/Xonesent/K8s-Hub/admin-panel/internal/business_models"
)

type AdminCHRepo interface {
	GetUserStatistics(ctx context.Context, tgId models.TgId) (models.UserLog, error)
}
