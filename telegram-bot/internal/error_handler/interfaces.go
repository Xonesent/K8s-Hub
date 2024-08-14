package error_handler

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type ErrHandler interface {
	HandleError(ctx context.Context, b *bot.Bot, update *models.Update, err error)
}
