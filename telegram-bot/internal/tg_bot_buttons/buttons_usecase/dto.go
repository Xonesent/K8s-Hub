package buttons_usecase

import (
	model "github.com/Xonesent/K8s-Hub/telegram-bot/internal/business_models"
	"github.com/Xonesent/K8s-Hub/telegram-bot/internal/tg_bot_buttons/buttons_repository"
)

type SentMessage struct {
	Sender  model.TgId
	Message string
	ChatId  int64
}

func (d *SentMessage) toSentMessage() buttons_repository.SentMessage {
	return buttons_repository.SentMessage{
		Sender:  d.Sender,
		Message: d.Message,
		ChatId:  d.ChatId,
	}
}

type SendMessage struct {
	Message string
	ChatId  uint32
}
