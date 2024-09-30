package config

import (
	"github.com/IBM/sarama"
	clickDB "github.com/Xonesent/K8s-Hub/statistics-sender/pkg/dependency_connectors/clickhouse"
	grpcConn "github.com/Xonesent/K8s-Hub/statistics-sender/pkg/dependency_connectors/grpc"
	"github.com/Xonesent/K8s-Hub/statistics-sender/pkg/dependency_connectors/kafka/config"
	"github.com/go-playground/validator/v10"
	"github.com/kelseyhightower/envconfig"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const path = "./config"

type Config struct {
	ClickHouse    clickDB.ConfigClickHouse `validate:"required"`
	KafkaSettings config.ConfigKafka       `validate:"required"`
	KafkaCfg      *sarama.Config           `validate:"required"`
	GRPC          grpcConn.ConfigGrpc
	Timers        []string `envconfig:"TIMERS"`
}

func LoadConfig() (cfg Config, err error) {
	if err := envconfig.Process("", &cfg); err != nil {
		return cfg, err
	}

	cfg.KafkaCfg = config.NewSaramaConfig(&cfg.KafkaSettings)
	cfg.GRPC.GrpcClientConn = make(map[string]*grpc.ClientConn, len(cfg.GRPC.GrpcClientPorts))

	for service, port := range cfg.GRPC.GrpcClientPorts {
		conn, err := grpc.Dial(cfg.GRPC.GrpcClientHosts[service]+":"+port,
			grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			return cfg, err
		}

		cfg.GRPC.GrpcClientConn[service] = conn
	}

	if err := validator.New().Struct(&cfg); err != nil {
		return cfg, err
	}

	return
}
