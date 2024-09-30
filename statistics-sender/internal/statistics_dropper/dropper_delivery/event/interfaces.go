package event_dropper

import "context"

type DropperUC interface {
	SendUserStatistics(ctx context.Context, chatId uint32) error
}
