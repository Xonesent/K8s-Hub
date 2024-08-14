package middleware

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type MDWManager struct{}

func NewMDWManager() *MDWManager {
	return &MDWManager{}
}

func (m *MDWManager) DefaultMiddleware(next bot.HandlerFunc) bot.HandlerFunc {
	return func(ctx context.Context, bot *bot.Bot, update *models.Update) {
		next(ctx, bot, update)
	}
}
