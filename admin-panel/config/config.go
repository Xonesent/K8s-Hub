package config

import (
	"strings"

	fiberApp "github.com/Xonesent/K8s-Hub/admin-panel/pkg/dependency_connectors/fiber"

	clickDB "github.com/Xonesent/K8s-Hub/admin-panel/pkg/dependency_connectors/clickhouse"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

const path = "./config"

type Config struct {
	ClickHouse clickDB.ConfigClickHouse `yaml:"ClickHouse" validate:"required"`
	Fiber      fiberApp.FiberConfig     `yaml:"Fiber" validate:"required"`
}

func LoadConfig(configName string) (cfg Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(configName)
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		return cfg, err
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		return cfg, err
	}

	if err := validator.New().Struct(&cfg); err != nil {
		return cfg, err
	}

	return
}
