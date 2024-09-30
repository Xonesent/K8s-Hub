package main

import (
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/IBM/sarama"
	"github.com/Xonesent/K8s-Hub/statistics-sender/config"
	"github.com/Xonesent/K8s-Hub/statistics-sender/internal/server"
	"github.com/Xonesent/K8s-Hub/statistics-sender/pkg/constant"
	clickDB "github.com/Xonesent/K8s-Hub/statistics-sender/pkg/dependency_connectors/clickhouse"
	"github.com/Xonesent/K8s-Hub/statistics-sender/pkg/helper_modules/logger"
	"go.uber.org/zap"
	"log"
)

func main() {
	if err := logger.Initialize(); err != nil {
		log.Fatalf("Error to init logger: %v\n", err)
	}

	//if err := godotenv.Load(constant.LocalEnvFile); err != nil {
	//	zap.L().Fatal("Error loading env variables", zap.Error(err))
	//}

	cfg, err := config.LoadConfig()
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

	kafkaProducer, err := sarama.NewSyncProducer(cfg.KafkaSettings.Brokers, cfg.KafkaCfg)
	if err != nil {
		zap.L().Fatal("Error creating kafkaProducer", zap.Error(err))
	}
	defer func(kafkaProducer sarama.SyncProducer) {
		if err = kafkaProducer.Close(); err != nil {
			zap.L().Error("KafkaProducer close error", zap.Error(err))
		} else {
			zap.L().Info("KafkaProducer closed properly")
		}
	}(kafkaProducer)

	consumerGroup, err := sarama.NewConsumerGroup(cfg.KafkaSettings.Brokers, constant.ConsumerGroupId, cfg.KafkaCfg)
	if err != nil {
		zap.L().Fatal("Error creating consumerGroup", zap.Error(err))
	}
	defer func(consumerGroup sarama.ConsumerGroup) {
		if err = consumerGroup.Close(); err != nil {
			zap.L().Error("ConsumerGroup close error", zap.Error(err))
		} else {
			zap.L().Info("ConsumerGroup closed properly")
		}
	}(consumerGroup)

	s := server.NewServer(
		&cfg,
		clickhouseDB,
		kafkaProducer,
		consumerGroup,
	)
	if err = s.Run(); err != nil {
		zap.L().Fatal("Cannot start server", zap.Error(err))
	}
}
