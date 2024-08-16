package main

import (
	"log"

	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/Xonesent/K8s-Hub/telegram-bot/config"
	"github.com/Xonesent/K8s-Hub/telegram-bot/internal/server"
	"github.com/Xonesent/K8s-Hub/telegram-bot/pkg/constant"
	clickDB "github.com/Xonesent/K8s-Hub/telegram-bot/pkg/dependency_connectors/clickhouse"
	"github.com/Xonesent/K8s-Hub/telegram-bot/pkg/dependency_connectors/telegram"
	"github.com/Xonesent/K8s-Hub/telegram-bot/pkg/helper_modules/logger"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main() {
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

	telegramBot, err := telegram.NewTelegramBot(cfg.Telegram)

	s := server.NewServer(
		&cfg,
		clickhouseDB,
		telegramBot,
	)
	if err = s.Run(); err != nil {
		zap.L().Fatal("Cannot start server", zap.Error(err))
	}
}
