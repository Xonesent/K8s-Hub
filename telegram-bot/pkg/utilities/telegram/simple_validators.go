package tg_utils

import (
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func ValidateDefaultHandler() bot.MatchFunc {
	return func(update *models.Update) bool {
		if update.Message != nil {
			return true
		}
		return false
	}
}
