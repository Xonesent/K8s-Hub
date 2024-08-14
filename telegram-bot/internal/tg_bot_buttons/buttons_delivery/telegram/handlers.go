package tg_buttons

import (
	"context"
	"github.com/Xonesent/K8s-Hub/telegram-bot/internal/error_handler"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"

	"github.com/Xonesent/K8s-Hub/telegram-bot/config"
)

type ButtonsHandler struct {
	cfg        *config.Config
	buttonsUC  ButtonsUC
	errHandler error_handler.ErrHandler
}

func NewButtonsHandler(cfg *config.Config, buttonsUC ButtonsUC, errHandler error_handler.ErrHandler) *ButtonsHandler {
	return &ButtonsHandler{
		cfg:        cfg,
		buttonsUC:  buttonsUC,
		errHandler: errHandler,
	}
}

func (h *ButtonsHandler) DefaultHandler() bot.HandlerFunc {
	return func(ctx context.Context, b *bot.Bot, update *models.Update) {
		sentMessageDTO := toSentMessage(update)

		if err := h.buttonsUC.DefaultHandler(ctx, &sentMessageDTO); err != nil {
			h.errHandler.HandleError(ctx, b, update, err)
		}
	}
}
