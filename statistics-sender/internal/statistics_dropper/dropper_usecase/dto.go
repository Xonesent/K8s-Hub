package dropper_usecase

import (
	"fmt"

	adminProto "github.com/Xonesent/K8s-Hub/statistics-sender/pkg/api/admin_proto"
	tgProto "github.com/Xonesent/K8s-Hub/statistics-sender/pkg/api/tg_proto"
	resp "github.com/Xonesent/K8s-Hub/statistics-sender/pkg/predefined_responses"
)

func toUserInfo(chatId uint32) *adminProto.UserInfo {
	return &adminProto.UserInfo{
		TgId: chatId,
	}
}

func toTgMessage(userStatistics *adminProto.UserStatistics) string {
	messageQuantity := len(userStatistics.GetMessages())
	if messageQuantity == 0 {
		return resp.StatisticsForm2
	}

	return fmt.Sprintf(
		resp.StatisticsForm1, userStatistics.GetTgId(), messageQuantity,
		userStatistics.GetMessages()[0].GetMessage(),
		userStatistics.GetMessages()[0].GetCreatedAt().AsTime().Format("02-01-2006 15:04:05"),
	)
}

func toMessageParams(chatId uint32, message string) *tgProto.MessageParams {
	return &tgProto.MessageParams{
		ChatId:  chatId,
		Message: message,
	}
}
