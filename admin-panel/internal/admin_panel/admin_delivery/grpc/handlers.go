package grpc_admin

import (
	"context"
	models "github.com/Xonesent/K8s-Hub/admin-panel/internal/business_models"
	adminProto "github.com/Xonesent/K8s-Hub/admin-panel/pkg/api/admin_proto"
)

type GrpcAdminHandler struct {
	adminProto.UnimplementedAdminServiceServer
	adminUC AdminUC
}

func NewGrpcAdminHandler(adminUC AdminUC) *GrpcAdminHandler {
	return &GrpcAdminHandler{
		adminUC: adminUC,
	}
}

func (h *GrpcAdminHandler) GetUserStatistics(ctx context.Context,
	userInfo *adminProto.UserInfo,
) (*adminProto.UserStatistics, error) {
	userStatistics, err := h.adminUC.GetUserStatistics(ctx, models.TgId(userInfo.TgId))
	if err != nil {
		return &adminProto.UserStatistics{}, err
	}

	return toProtoUserStatistics(userStatistics), err
}
