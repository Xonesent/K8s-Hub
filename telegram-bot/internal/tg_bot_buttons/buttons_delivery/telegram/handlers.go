package tg_buttons

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"

	"github.com/Xonesent/K8s-Hub/telegram-bot/config"
)

type ButtonsHandler struct {
	cfg *config.Config
}

func NewButtonsHandler(cfg *config.Config) ButtonsHandler {
	return ButtonsHandler{
		cfg: cfg,
	}
}

func (h *ButtonsHandler) StartBot() bot.HandlerFunc {
	return func(ctx context.Context, bot *bot.Bot, update *models.Update) {
	}
}
