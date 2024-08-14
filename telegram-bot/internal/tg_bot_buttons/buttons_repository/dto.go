package buttons_repository

import model "github.com/Xonesent/K8s-Hub/telegram-bot/internal/business_models"

type SentMessage struct {
	Sender  model.TgId
	Message string
	ChatId  int64
}
