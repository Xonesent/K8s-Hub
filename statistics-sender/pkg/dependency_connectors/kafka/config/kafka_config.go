package config

import (
	"strings"
	"time"

	"github.com/IBM/sarama"
	"go.uber.org/zap"
)

type ConfigKafka struct {
	Brokers  []string       `envconfig:"KAFKA_BROKERS" validate:"required"`
	Producer ConfigProducer `validate:"required"`
	Consumer ConfigConsumer `validate:"required"`
}

type ConfigProducer struct {
	RequiredAcks    int  `envconfig:"KAFKA_PRODUCER_ACKS" validate:"required"`
	Retries         int  `envconfig:"KAFKA_PRODUCER_RETRIES" validate:"required"`
	ReturnSuccesses bool `envconfig:"KAFKA_PRODUCER_RETURN" validate:"required"`
}

type ConfigConsumer struct {
	Group            ConfigGroup `validate:"required"`
	Retries          int         `envconfig:"KAFKA_CONSGROUP_RETRIES" validate:"required"`
	OffsetAutoCommit bool        `envconfig:"KAFKA_CONSGROUP_AUTOCOMMIT" validate:"required"`
}

type ConfigGroup struct {
	RebalanceStrategy []string      `envconfig:"KAFKA_CONSGROUP_REBALANCE" validate:"required"`
	OffsetsInitial    string        `envconfig:"KAFKA_CONSGROUP_OFFSETS" validate:"required"`
	SessionTimeoutMs  time.Duration `envconfig:"KAFKA_CONSGROUP_SESSION" validate:"required"`
}

func NewSaramaConfig(kafkaConfig *ConfigKafka) *sarama.Config {
	config := sarama.NewConfig()

	config.Producer.RequiredAcks = sarama.RequiredAcks(kafkaConfig.Producer.RequiredAcks)
	config.Producer.Retry.Max = kafkaConfig.Producer.Retries
	config.Producer.Return.Successes = kafkaConfig.Producer.ReturnSuccesses

	config.Consumer.Group.Rebalance.GroupStrategies = getRebalanceStrategies(kafkaConfig.Consumer.Group.RebalanceStrategy)
	config.Consumer.Offsets.Initial = getOffsetInitial(kafkaConfig.Consumer.Group.OffsetsInitial)
	config.Consumer.Group.Session.Timeout = kafkaConfig.Consumer.Group.SessionTimeoutMs
	config.Consumer.Offsets.Retry.Max = kafkaConfig.Consumer.Retries
	config.Consumer.Offsets.AutoCommit.Enable = kafkaConfig.Consumer.OffsetAutoCommit

	return config
}

func getRebalanceStrategies(strategies []string) []sarama.BalanceStrategy {
	var balanceStrategies []sarama.BalanceStrategy

	for _, strategy := range strategies {
		switch strings.ToLower(strategy) {
		case "roundrobin":
			balanceStrategies = append(balanceStrategies, sarama.NewBalanceStrategyRoundRobin())
		case "range":
			balanceStrategies = append(balanceStrategies, sarama.NewBalanceStrategyRange())
		case "sticky":
			balanceStrategies = append(balanceStrategies, sarama.NewBalanceStrategySticky())
		default:
			zap.L().Info("Неподдерживаемая стратегия, заменено на RoundRobin", zap.String("input", strategy))
			balanceStrategies = append(balanceStrategies, sarama.NewBalanceStrategyRoundRobin())
		}
	}

	return balanceStrategies
}

func getOffsetInitial(offset string) int64 {
	switch offset {
	case "oldest":
		return sarama.OffsetOldest
	case "newest":
		return sarama.OffsetNewest
	default:
		return sarama.OffsetNewest
	}
}
