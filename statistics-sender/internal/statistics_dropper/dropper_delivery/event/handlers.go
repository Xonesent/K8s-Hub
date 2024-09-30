package event_dropper

import (
	"context"
	"encoding/json"

	"github.com/IBM/sarama"
	"github.com/Xonesent/K8s-Hub/statistics-sender/config"
	models "github.com/Xonesent/K8s-Hub/statistics-sender/internal/business_models"
)

type DropperHandler struct {
	cfg       *config.Config
	dropperUC DropperUC
}

func NewDropperHandler(cfg *config.Config, dropperUC DropperUC) *DropperHandler {
	return &DropperHandler{
		cfg:       cfg,
		dropperUC: dropperUC,
	}
}

func (h *DropperHandler) SendUserStatistics(ctx context.Context, msg *sarama.ConsumerMessage) error {
	var statisticsEvent models.StatisticsEvent

	if err := json.Unmarshal(msg.Value, &statisticsEvent); err != nil {
		return err
	}

	if err := h.dropperUC.SendUserStatistics(ctx, statisticsEvent.ChatId); err != nil {
		return err
	}

	return nil
}
