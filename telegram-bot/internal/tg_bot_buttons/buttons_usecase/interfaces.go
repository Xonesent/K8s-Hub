package buttons_usecase

import (
	"context"
	"github.com/Xonesent/K8s-Hub/telegram-bot/internal/tg_bot_buttons/buttons_repository"
)

type ButtonsCHRepo interface {
	InsertLastMessage(ctx context.Context, sentMessage *buttons_repository.SentMessage) error
}
