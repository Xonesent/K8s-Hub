package http_admin

import (
	"time"

	models "github.com/Xonesent/K8s-Hub/admin-panel/internal/business_models"
)

type UserStatistics struct {
	TgId       models.TgId `json:"tg_id"`
	ChatId     uint32      `json:"chat_id"`
	TotalCount int64       `json:"total_count"`
	Messages   []Message   `json:"messages"`
}

type Message struct {
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
}

func toUserStatistics(userLogs models.UserLog) UserStatistics {
	messages := make([]Message, 0)
	for _, message := range userLogs.Messages {
		messages = append(messages, Message{message.Message, message.CreatedAt})
	}

	return UserStatistics{
		TgId:       userLogs.TgId,
		ChatId:     userLogs.ChatId,
		TotalCount: int64(len(messages)),
		Messages:   messages,
	}
}
