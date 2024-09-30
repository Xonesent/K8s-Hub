package server

import (
	"github.com/Xonesent/K8s-Hub/telegram-bot/internal/error_handler"
	grpc_tg "github.com/Xonesent/K8s-Hub/telegram-bot/internal/tg_bot_buttons/buttons_delivery/grpc"
	tg_buttons "github.com/Xonesent/K8s-Hub/telegram-bot/internal/tg_bot_buttons/buttons_delivery/telegram"
	"github.com/Xonesent/K8s-Hub/telegram-bot/internal/tg_bot_buttons/buttons_repository"
	"github.com/Xonesent/K8s-Hub/telegram-bot/internal/tg_bot_buttons/buttons_usecase"
	middleware "github.com/Xonesent/K8s-Hub/telegram-bot/internal/tg_bot_middleware"
	tgProto "github.com/Xonesent/K8s-Hub/telegram-bot/pkg/api/tg_proto"
)

func (s *Server) MapHandlers() {
	buttonsCHRepo := buttons_repository.NewButtonsCHRepository(s.cfg, s.clickhouse)

	buttonsUC := buttons_usecase.NewButtonsUseCase(s.cfg, buttonsCHRepo, s.tgBot)

	errHandler := error_handler.NewErrorHandler(s.cfg)

	buttonsHDL := tg_buttons.NewButtonsHandler(s.cfg, buttonsUC, errHandler)

	tgGrpc := grpc_tg.NewGrpcTgHandler(buttonsUC)
	tgProto.RegisterTgServiceServer(s.gRPCServer, tgGrpc)

	mw := middleware.NewMDWManager()

	tg_buttons.MapButtonsRoutes(s.tgBot, buttonsHDL, mw)
}
