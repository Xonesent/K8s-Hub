package server

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/Xonesent/K8s-Hub/admin-panel/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"go.uber.org/zap"
)

type Server struct {
	cfg        *config.Config
	clickhouse driver.Conn
	fiberApp   *fiber.App
}

func NewServer(
	cfg *config.Config,
	clickhouse driver.Conn,
	fiberApp *fiber.App,
) *Server {
	return &Server{
		cfg:        cfg,
		clickhouse: clickhouse,
		fiberApp:   fiberApp,
	}
}

func (s *Server) Run() error {
	zap.L().Info("Trying to run server...")

	s.MapHandlers()

	go func() {
		s.fiberApp.Get("/health_check", func(c *fiber.Ctx) error {
			return c.SendStatus(fiber.StatusOK)
		})
		zap.L().Info(fmt.Sprintf("Server is started %s:%s", s.cfg.Fiber.Host, s.cfg.Fiber.Port))

		if err := s.fiberApp.Listen(fmt.Sprintf("%s:%s", s.cfg.Fiber.Host, s.cfg.Fiber.Port)); err != nil {
			log.Fatal(err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	return nil
}
