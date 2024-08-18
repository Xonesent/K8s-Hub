package admin_repository

import (
	"time"

	models "github.com/Xonesent/K8s-Hub/admin-panel/internal/business_models"
)

type UserMessage struct {
	TgId      uint32    `ch:"tg_id"`
	ChatId    uint32    `ch:"chat_id"`
	Message   string    `ch:"message"`
	CreatedAt time.Time `ch:"created_at"`
}

func toUserLog(userMessages []UserMessage) models.UserLog {
	messages := make([]models.Message, 0)
	for _, userMessage := range userMessages {
		messages = append(messages, models.Message{Message: userMessage.Message, CreatedAt: userMessage.CreatedAt})
	}

	return models.UserLog{
		TgId:     models.TgId(userMessages[0].TgId),
		ChatId:   userMessages[0].ChatId,
		Messages: messages,
	}
}
