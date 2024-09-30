package grpc_tg

import (
	"github.com/Xonesent/K8s-Hub/telegram-bot/internal/tg_bot_buttons/buttons_usecase"
	tgProto "github.com/Xonesent/K8s-Hub/telegram-bot/pkg/api/tg_proto"
)

func toSendMessage(msg *tgProto.MessageParams) *buttons_usecase.SendMessage {
	return &buttons_usecase.SendMessage{
		Message: msg.GetMessage(),
		ChatId:  msg.GetChatId(),
	}
}
