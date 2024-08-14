package error_handler

import (
	"context"
	"github.com/Xonesent/K8s-Hub/telegram-bot/config"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"go.uber.org/zap"
)

type ErrorHandler struct {
	cfg *config.Config
}

func NewErrorHandler(cfg *config.Config) *ErrorHandler {
	return &ErrorHandler{
		cfg: cfg,
	}
}

func (h *ErrorHandler) HandleError(ctx context.Context, b *bot.Bot, update *models.Update, err error) {
	zap.L().Error(err.Error())
}
