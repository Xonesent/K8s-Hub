package grpc_tg

import (
	"context"

	"github.com/Xonesent/K8s-Hub/telegram-bot/pkg/api/tg_proto"
)

type GrpcTgHandler struct {
	tgProto.UnimplementedTgServiceServer
	tgUC TgUC
}

func NewGrpcTgHandler(tgUC TgUC) *GrpcTgHandler {
	return &GrpcTgHandler{
		tgUC: tgUC,
	}
}

func (h *GrpcTgHandler) BotSendMessage(ctx context.Context, msg *tgProto.MessageParams) (*tgProto.SendResponse, error) {
	sendMessageDTO := toSendMessage(msg)

	err := h.tgUC.BotSendMessage(ctx, sendMessageDTO)

	return &tgProto.SendResponse{}, err
}
