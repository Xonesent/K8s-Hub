package tg_buttons

import (
	middleware "github.com/Xonesent/K8s-Hub/telegram-bot/internal/tg_bot_middleware"
	"github.com/go-telegram/bot"
)

func MapButtonsRoutes(b *bot.Bot, h *ButtonsHandler, mw *middleware.MDWManager) {
	b.RegisterHandler(bot.HandlerTypeMessageText, "/start", bot.MatchTypeExact, mw.DefaultMiddleware(h.StartBot()))
}
