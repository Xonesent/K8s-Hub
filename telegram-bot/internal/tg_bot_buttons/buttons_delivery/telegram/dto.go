package tg_buttons

import (
	model "github.com/Xonesent/K8s-Hub/telegram-bot/internal/business_models"
	"github.com/Xonesent/K8s-Hub/telegram-bot/internal/tg_bot_buttons/buttons_usecase"
	"github.com/go-telegram/bot/models"
)

func toSentMessage(update *models.Update) buttons_usecase.SentMessage {
	return buttons_usecase.SentMessage{
		Sender:  model.TgId(update.Message.From.ID),
		ChatId:  update.Message.Chat.ID,
		Message: update.Message.Text,
	}
}
