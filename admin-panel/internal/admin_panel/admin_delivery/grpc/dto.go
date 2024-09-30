package grpc_admin

import (
	models "github.com/Xonesent/K8s-Hub/admin-panel/internal/business_models"
	adminProto "github.com/Xonesent/K8s-Hub/admin-panel/pkg/api/admin_proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func toProtoUserStatistics(userStatistics models.UserLog) *adminProto.UserStatistics {
	protoMessages := make([]*adminProto.Message, 0)
	for _, message := range userStatistics.Messages {
		protoMessages = append(protoMessages, &adminProto.Message{
			Message: message.Message, CreatedAt: timestamppb.New(message.CreatedAt),
		})
	}

	return &adminProto.UserStatistics{
		TgId:     uint32(userStatistics.TgId),
		ChatId:   userStatistics.ChatId,
		Messages: protoMessages,
	}
}
