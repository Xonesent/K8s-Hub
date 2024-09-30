package grpc_tg

import (
	"context"

	"github.com/Xonesent/K8s-Hub/telegram-bot/internal/tg_bot_buttons/buttons_usecase"
)

type TgUC interface {
	BotSendMessage(ctx context.Context, sendMessage *buttons_usecase.SendMessage) error
}
