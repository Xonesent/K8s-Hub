package timer_reminder

import (
	"log"
	"time"

	utils "github.com/Xonesent/K8s-Hub/statistics-sender/pkg/utilities"
	"go.uber.org/zap"
)

type ReminderHDL interface {
	PublishStatisticsEvent()
}

func SetRemindTimer(triggers []string, h *ReminderHandler) {
	for _, trigger := range triggers {
		parsedTime, err := utils.ParseTimeStr(trigger)
		if err != nil {
			zap.L().Info(err.Error())
			return
		}

		duration, err := utils.ValidateTimer(parsedTime)
		if err != nil {
			zap.L().Info(err.Error())
			return
		}

		timer := time.NewTimer(duration)

		go func(t string, timer *time.Timer) {
			<-timer.C
			h.PublishStatisticsEvent()
			log.Printf("Время %s наступило! Выполняем действие...\n", t)
		}(trigger, timer)
	}
}
