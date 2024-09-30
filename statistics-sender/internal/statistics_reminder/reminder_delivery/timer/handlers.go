package timer_reminder

import (
	"github.com/Xonesent/K8s-Hub/statistics-sender/config"
	"go.uber.org/zap"
)

type ReminderHandler struct {
	cfg        *config.Config
	reminderUC ReminderUC
}

func NewReminderHandler(cfg *config.Config, reminderUC ReminderUC) *ReminderHandler {
	return &ReminderHandler{
		cfg:        cfg,
		reminderUC: reminderUC,
	}
}

func (h *ReminderHandler) PublishStatisticsEvent() {
	err := h.reminderUC.PublishStatisticsEvent()
	if err != nil {
		zap.L().Error(err.Error())
	}
}
