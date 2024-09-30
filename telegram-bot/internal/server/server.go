package server

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"
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
	gRPCServer *grpc.Server
}

func NewServer(
	cfg *config.Config,
	clickhouse driver.Conn,
	tgBot *bot.Bot,
	gRPCServer *grpc.Server,
) *Server {
	return &Server{
		cfg:        cfg,
		clickhouse: clickhouse,
		tgBot:      tgBot,
		gRPCServer: gRPCServer,
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

	go func() {
		grpcAddress := fmt.Sprintf("%s:%s", s.cfg.Grpc.Host, s.cfg.Grpc.Port)

		listener, err := net.Listen("tcp", grpcAddress)
		if err != nil {
			zap.L().Fatal("Failed to listen", zap.Error(err))
		}

		zap.L().Info("Grpc Server is started on " + grpcAddress)

		if err := s.gRPCServer.Serve(listener); err != nil {
			zap.L().Fatal("Failed to GRPC serve", zap.Error(err))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	zap.L().Info("Server is closing")

	return nil
}
