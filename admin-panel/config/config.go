package config

import (
	fiberApp "github.com/Xonesent/K8s-Hub/admin-panel/pkg/dependency_connectors/fiber"
	grpcServer "github.com/Xonesent/K8s-Hub/admin-panel/pkg/dependency_connectors/grpc"

	clickDB "github.com/Xonesent/K8s-Hub/admin-panel/pkg/dependency_connectors/clickhouse"
	"github.com/go-playground/validator/v10"
	"github.com/kelseyhightower/envconfig"
)

const path = "./config"

type Config struct {
	ClickHouse clickDB.ConfigClickHouse `validate:"required"`
	Fiber      fiberApp.FiberConfig     `validate:"required"`
	Grpc       grpcServer.GRPCConfig    `validate:"required"`
}

func LoadConfig() (cfg Config, err error) {
	if err := envconfig.Process("", &cfg); err != nil {
		return cfg, err
	}

	if err := validator.New().Struct(&cfg); err != nil {
		return cfg, err
	}

	return
}
