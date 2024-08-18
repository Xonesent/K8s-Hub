package models

import "time"

type UserLog struct {
	TgId     TgId
	ChatId   uint32
	Messages []Message
}

type Message struct {
	Message   string
	CreatedAt time.Time
}
