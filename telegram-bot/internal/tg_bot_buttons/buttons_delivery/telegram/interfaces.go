package tg_buttons

import "github.com/go-telegram/bot"

type ButtonsHDL interface {
	StartBot() bot.HandlerFunc
}
