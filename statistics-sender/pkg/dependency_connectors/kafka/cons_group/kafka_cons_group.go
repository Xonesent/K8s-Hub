package cons_group

import (
	"fmt"
	"github.com/IBM/sarama"
	event_dropper "github.com/Xonesent/K8s-Hub/statistics-sender/internal/statistics_dropper/dropper_delivery/event"
	"go.uber.org/zap"
)

type Consumer struct {
	Ready      chan bool
	dropperHDL event_dropper.DropperHDL
}

func NewConsumerHandler(dropperHDL event_dropper.DropperHDL) Consumer {
	return Consumer{
		Ready:      make(chan bool),
		dropperHDL: dropperHDL,
	}
}

func (c *Consumer) Setup(sarama.ConsumerGroupSession) error {
	close(c.Ready)
	return nil
}

func (c *Consumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (c *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for {
		select {
		case msg, ok := <-claim.Messages():
			if !ok {
				zap.L().Info("message channel was closed")
				return nil
			}
			zap.L().Info(fmt.Sprintf("Message claimed: value = %s, timestamp = %v, topic = %s", string(msg.Value), msg.Timestamp, msg.Topic))

			if err := event_dropper.MapDropperEvents(session.Context(), msg, c.dropperHDL); err != nil {
				zap.L().Info("Error handling message", zap.Error(err))
			}

			session.MarkMessage(msg, "")
		case <-session.Context().Done():
			return nil
		}
	}
}
