package server

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/Xonesent/K8s-Hub/admin-panel/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type Server struct {
	cfg        *config.Config
	clickhouse driver.Conn
	fiberApp   *fiber.App
	gRPCServer *grpc.Server
}

func NewServer(
	cfg *config.Config,
	clickhouse driver.Conn,
	fiberApp *fiber.App,
	gRPCServer *grpc.Server,
) *Server {
	return &Server{
		cfg:        cfg,
		clickhouse: clickhouse,
		fiberApp:   fiberApp,
		gRPCServer: gRPCServer,
	}
}

func (s *Server) Run() error {
	zap.L().Info("Trying to run server...")

	s.MapHandlers()

	go func() {
		fiberAddress := fmt.Sprintf("%s:%s", s.cfg.Fiber.Host, s.cfg.Fiber.Port)

		s.fiberApp.Get("/health_check", func(c *fiber.Ctx) error {
			return c.SendStatus(fiber.StatusOK)
		})
		zap.L().Info("Fiber server is started on " + fiberAddress)

		if err := s.fiberApp.Listen(fiberAddress); err != nil {
			log.Fatal(err.Error())
		}
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

	zap.L().Info("Shutting down server...")

	return nil
}
