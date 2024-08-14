package tg_buttons

import (
	middleware "github.com/Xonesent/K8s-Hub/telegram-bot/internal/tg_bot_middleware"
	tgUtils "github.com/Xonesent/K8s-Hub/telegram-bot/pkg/utilities/telegram"
	"github.com/go-telegram/bot"
)

type ButtonsHDL interface {
	DefaultHandler() bot.HandlerFunc
}

func MapButtonsRoutes(b *bot.Bot, h *ButtonsHandler, mw *middleware.MDWManager) {
	b.RegisterHandlerMatchFunc(tgUtils.ValidateDefaultHandler(), mw.DefaultMiddleware(h.DefaultHandler()))
}
