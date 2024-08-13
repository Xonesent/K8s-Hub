package clickDB

import (
	"context"
	"errors"
	"fmt"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"go.uber.org/zap"
)

type ConfigClickHouse struct {
	Host     string `yaml:"Host" validate:"required"`
	Port     string `yaml:"Port" validate:"required"`
	Database string `yaml:"Database" validate:"required"`
	User     string `yaml:"User" validate:"required"`
	Password string `yaml:"Password" validate:"required"`
}

// nolint:ireturn // Connector was taken from clickhouse repository (Probably best practices)
func NewClickhouseDB(cfg *ConfigClickHouse) (driver.Conn, error) {
	ctx := context.Background()

	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)},
		Auth: clickhouse.Auth{
			Database: cfg.Database,
			Username: cfg.User,
			Password: cfg.Password,
		},
	})
	if err != nil {
		return nil, err
	}

	if err := conn.Ping(ctx); err != nil {
		var exception *clickhouse.Exception

		if errors.As(err, &exception) {
			zap.L().Info(fmt.Sprintf("Exception [%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace))
		}

		return nil, err
	}

	return conn, nil
}
