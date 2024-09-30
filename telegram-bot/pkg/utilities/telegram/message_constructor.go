package tg_utils

import "github.com/go-telegram/bot"

func TgSendMessage(chatId int64, response string) bot.SendMessageParams {
	return bot.SendMessageParams{
		ChatID: chatId,
		Text:   response,
	}
}
