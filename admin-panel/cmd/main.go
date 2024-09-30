package main

import (
	"context"
	"log"

	grpcServer "github.com/Xonesent/K8s-Hub/admin-panel/pkg/dependency_connectors/grpc"
	"google.golang.org/grpc"

	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/Xonesent/K8s-Hub/admin-panel/config"
	"github.com/Xonesent/K8s-Hub/admin-panel/internal/server"
	clickDB "github.com/Xonesent/K8s-Hub/admin-panel/pkg/dependency_connectors/clickhouse"
	fiberApp "github.com/Xonesent/K8s-Hub/admin-panel/pkg/dependency_connectors/fiber"
	"github.com/Xonesent/K8s-Hub/admin-panel/pkg/helper_modules/logger"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"

	_ "github.com/Xonesent/K8s-Hub/admin-panel/cmd/docs"
)

func main() {
	ctx := context.Background()

	if err := logger.Initialize(); err != nil {
		log.Fatalf("Error to init logger: %v\n", err)
	}

	// nolint:gocritic // needed for local compilation
	//if err := godotenv.Load(constant.LocalEnvFile); err != nil {
	//	zap.L().Fatal("Error loading env variables", zap.Error(err))
	//}

	cfg, err := config.LoadConfig()
	if err != nil {
		zap.L().Fatal("Error loading config", zap.Error(err))
	}

	run(ctx, &cfg)
}

func run(ctx context.Context, cfg *config.Config) {
	clickhouseDB, err := clickDB.NewClickhouseDB(&cfg.ClickHouse) // nolint:contextcheck // not needed
	if err != nil {
		zap.L().Fatal("Error connecting clickhouse", zap.Error(err))
	}
	defer func(clickhouseDB *driver.Conn) {
		if err = (*clickhouseDB).Close(); err != nil {
			zap.L().Error("Clickhouse close error", zap.Error(err))
		} else {
			zap.L().Info("Clickhouse closed properly")
		}
	}(&clickhouseDB)

	fiberClient := fiberApp.NewFiberClient()
	defer func(fiberClient *fiber.App) {
		if err := fiberClient.ShutdownWithContext(ctx); err != nil {
			zap.L().Error("Fiber close error", zap.Error(err))
		} else {
			zap.L().Info("Fiber closed properly")
		}
	}(fiberClient)

	gRPCServer := grpcServer.NewGRPCServer()
	defer func(gRPCServer *grpc.Server) {
		gRPCServer.GracefulStop()
		zap.L().Info("Grpc Server closed properly")
	}(gRPCServer)

	s := server.NewServer(
		cfg,
		clickhouseDB,
		fiberClient,
		gRPCServer,
	)
	if err = s.Run(); err != nil {
		zap.L().Fatal("Cannot start server", zap.Error(err))
	}
}
