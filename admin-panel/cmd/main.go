package main

import (
	"context"
	"log"

	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/Xonesent/K8s-Hub/admin-panel/config"
	"github.com/Xonesent/K8s-Hub/admin-panel/internal/server"
	"github.com/Xonesent/K8s-Hub/admin-panel/pkg/constant"
	clickDB "github.com/Xonesent/K8s-Hub/admin-panel/pkg/dependency_connectors/clickhouse"
	fiberApp "github.com/Xonesent/K8s-Hub/admin-panel/pkg/dependency_connectors/fiber"
	"github.com/Xonesent/K8s-Hub/admin-panel/pkg/helper_modules/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main() {
	ctx := context.Background()

	if err := logger.Initialize(); err != nil {
		log.Fatalf("Error to init logger: %v\n", err)
	}

	if err := godotenv.Load(constant.EnvFile); err != nil {
		zap.L().Fatal("Error loading env variables", zap.Error(err))
	}

	cfg, err := config.LoadConfig(constant.DevConfig)
	if err != nil {
		zap.L().Fatal("Error loading config", zap.Error(err))
	}

	clickhouseDB, err := clickDB.NewClickhouseDB(&cfg.ClickHouse)
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

	s := server.NewServer(
		&cfg,
		clickhouseDB,
		fiberClient,
	)
	if err = s.Run(); err != nil {
		zap.L().Fatal("Cannot start server", zap.Error(err))
	}
}
