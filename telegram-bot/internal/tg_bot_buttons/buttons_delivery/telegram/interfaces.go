package tg_buttons

import (
	"context"
	"github.com/Xonesent/K8s-Hub/telegram-bot/internal/tg_bot_buttons/buttons_usecase"
)

type ButtonsUC interface {
	DefaultHandler(ctx context.Context, sentMessage *buttons_usecase.SentMessage) error
}
