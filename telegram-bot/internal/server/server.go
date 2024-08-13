package server

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/go-telegram/bot"
	"go.uber.org/zap"

	"github.com/Xonesent/K8s-Hub/telegram-bot/config"
)

type Server struct {
	cfg        *config.Config
	clickhouse driver.Conn
	tgBot      *bot.Bot
}

func NewServer(
	cfg *config.Config,
	clickhouse driver.Conn,
	tgBot *bot.Bot,
) *Server {
	return &Server{
		cfg:        cfg,
		clickhouse: clickhouse,
		tgBot:      tgBot,
	}
}

func (s *Server) Run() error {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	defer cancel()

	zap.L().Info("Trying to run server...")

	s.MapHandlers()

	go func() {
		s.tgBot.Start(ctx)
	}()

	zap.L().Info("Server is running")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	zap.L().Info("Server is closing")

	return nil
}
