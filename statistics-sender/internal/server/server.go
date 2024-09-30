package server

import (
	"context"
	"errors"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/Xonesent/K8s-Hub/statistics-sender/pkg/constant"

	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/IBM/sarama"
	"github.com/Xonesent/K8s-Hub/statistics-sender/config"
	"go.uber.org/zap"
)

type Server struct {
	cfg           *config.Config
	clickhouse    driver.Conn
	kafkaProducer sarama.SyncProducer
	consumerGroup sarama.ConsumerGroup
}

func NewServer(
	cfg *config.Config,
	clickhouse driver.Conn,
	kafkaProducer sarama.SyncProducer,
	consumerGroup sarama.ConsumerGroup,
) *Server {
	return &Server{
		cfg:           cfg,
		clickhouse:    clickhouse,
		kafkaProducer: kafkaProducer,
		consumerGroup: consumerGroup,
	}
}

func (s *Server) Run() error {
	zap.L().Info("Trying to run server...")

	consGroup, err := s.MapHandlers()
	if err != nil {
		log.Fatalf("map handlers: %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()

		for {
			if err := s.consumerGroup.Consume(ctx, []string{constant.TopicK8sHub}, &consGroup); err != nil {
				if errors.Is(err, sarama.ErrClosedConsumerGroup) {
					return
				}

				log.Fatalf("Error from consumer: %v", err)
			}

			if ctx.Err() != nil {
				return
			}

			consGroup.Ready = make(chan bool)
		}
	}()

	<-consGroup.Ready

	zap.L().Info("Sarama consumer up and running!...")

	zap.L().Info("Server is running")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	cancel()
	wg.Wait()

	zap.L().Info("Server is closing")

	return nil
}
