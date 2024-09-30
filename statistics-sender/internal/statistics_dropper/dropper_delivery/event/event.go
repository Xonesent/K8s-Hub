package event_dropper

import (
	"context"
	"github.com/IBM/sarama"
)

type DropperHDL interface {
	SendUserStatistics(ctx context.Context, msg *sarama.ConsumerMessage) error
}

func MapDropperEvents(ctx context.Context, msg *sarama.ConsumerMessage, h DropperHDL) error {
	err := h.SendUserStatistics(ctx, msg)
	if err != nil {
		return err
	}

	return nil
}
