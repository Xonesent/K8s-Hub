package dropper_usecase

import (
	"fmt"
	adminProto "github.com/Xonesent/K8s-Hub/statistics-sender/pkg/api/admin_proto"
	tgProto "github.com/Xonesent/K8s-Hub/statistics-sender/pkg/api/tg_proto"
	resp "github.com/Xonesent/K8s-Hub/statistics-sender/pkg/predefined_responces"
)

func toUserInfo(chatId uint32) *adminProto.UserInfo {
	return &adminProto.UserInfo{
		TgId: chatId,
	}
}

func toTgMessage(userStatistics *adminProto.UserStatistics) string {
	messageQuantity := len(userStatistics.Messages)
	if messageQuantity == 0 {
		return resp.StatisticsForm2
	}
	return fmt.Sprintf(resp.StatisticsForm1, userStatistics.TgId, messageQuantity, userStatistics.Messages[0].Message,
		userStatistics.Messages[0].CreatedAt.AsTime().Format("02-01-2006 15:04:05"))
}

func toMessageParams(chatId uint32, message string) *tgProto.MessageParams {
	return &tgProto.MessageParams{
		ChatId:  chatId,
		Message: message,
	}
}
