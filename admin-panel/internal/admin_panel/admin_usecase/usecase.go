package admin_usecase

import (
	"context"

	"github.com/Xonesent/K8s-Hub/admin-panel/config"
	models "github.com/Xonesent/K8s-Hub/admin-panel/internal/business_models"
)

type AdminUseCase struct {
	cfg       *config.Config
	adminRepo AdminCHRepo
}

func NewAdminUseCase(cfg *config.Config, adminRepo AdminCHRepo) *AdminUseCase {
	return &AdminUseCase{
		cfg:       cfg,
		adminRepo: adminRepo,
	}
}

func (u *AdminUseCase) GetUserStatistics(ctx context.Context, tgId models.TgId) (models.UserLog, error) {
	userStatistics, err := u.adminRepo.GetUserStatistics(ctx, tgId)
	if err != nil {
		return models.UserLog{}, err
	}

	return userStatistics, nil
}
