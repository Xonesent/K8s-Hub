package dropper_usecase

import (
	"context"
	"github.com/Xonesent/K8s-Hub/statistics-sender/config"
	adminProto "github.com/Xonesent/K8s-Hub/statistics-sender/pkg/api/admin_proto"
	tgProto "github.com/Xonesent/K8s-Hub/statistics-sender/pkg/api/tg_proto"
)

type DropperUseCase struct {
	cfg       *config.Config
	adminGrpc adminProto.AdminServiceClient
	tgGrpc    tgProto.TgServiceClient
}

func NewDropperUseCase(
	cfg *config.Config,
	adminGrpc adminProto.AdminServiceClient,
	tgGrpc tgProto.TgServiceClient,
) *DropperUseCase {
	return &DropperUseCase{
		cfg:       cfg,
		adminGrpc: adminGrpc,
		tgGrpc:    tgGrpc,
	}
}

func (h *DropperUseCase) SendUserStatistics(ctx context.Context, chatId uint32) error {
	userStatistics, err := h.adminGrpc.GetUserStatistics(ctx, toUserInfo(chatId))
	if err != nil {
		return err
	}

	if _, err = h.tgGrpc.BotSendMessage(ctx, toMessageParams(chatId, toTgMessage(userStatistics))); err != nil {
		return err
	}

	return nil
}
