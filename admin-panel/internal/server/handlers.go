package server

import (
	grpc_admin "github.com/Xonesent/K8s-Hub/admin-panel/internal/admin_panel/admin_delivery/grpc"
	http_admin "github.com/Xonesent/K8s-Hub/admin-panel/internal/admin_panel/admin_delivery/http"
	"github.com/Xonesent/K8s-Hub/admin-panel/internal/admin_panel/admin_repository"
	"github.com/Xonesent/K8s-Hub/admin-panel/internal/admin_panel/admin_usecase"
	"github.com/Xonesent/K8s-Hub/admin-panel/internal/middleware"
	adminProto "github.com/Xonesent/K8s-Hub/admin-panel/pkg/api/admin_proto"
)

func (s *Server) MapHandlers() {
	adminCHRepo := admin_repository.NewAdminCHRepository(s.cfg, s.clickhouse)

	adminUC := admin_usecase.NewAdminUseCase(s.cfg, adminCHRepo)

	adminHDL := http_admin.NewAdminHandler(s.cfg, adminUC)

	adminGrpc := grpc_admin.NewGrpcAdminHandler(adminUC)
	adminProto.RegisterAdminServiceServer(s.gRPCServer, adminGrpc)

	mw := middleware.NewMiddlewareManager(s.cfg)

	adminGroup := s.fiberApp.Group("admin")

	http_admin.MapAdminRoutes(adminGroup, adminHDL, mw)
}
