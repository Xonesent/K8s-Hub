package reminder_usecase

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/IBM/sarama"
	"github.com/Xonesent/K8s-Hub/statistics-sender/config"
	models "github.com/Xonesent/K8s-Hub/statistics-sender/internal/business_models"
	"github.com/Xonesent/K8s-Hub/statistics-sender/pkg/constant"
	"github.com/gammazero/workerpool"
)

type ReminderUseCase struct {
	cfg            *config.Config
	reminderCHRepo ReminderCHRepo
	producer       sarama.SyncProducer
}

func NewReminderUseCase(
	cfg *config.Config,
	reminderCHRepo ReminderCHRepo,
	producer sarama.SyncProducer,
) *ReminderUseCase {
	return &ReminderUseCase{
		cfg:            cfg,
		reminderCHRepo: reminderCHRepo,
		producer:       producer,
	}
}

func (u *ReminderUseCase) PublishStatisticsEvent() error {
	chatIds, err := u.reminderCHRepo.GetUniqueChatIds(context.Background())
	if err != nil {
		return err
	}

	wp := workerpool.New(6)
	errChan := make(chan error, len(chatIds))

	for _, chatId := range chatIds {
		bytes, err := json.Marshal(models.StatisticsEvent{ChatId: chatId.ChatId})
		if err != nil {
			errChan <- fmt.Errorf("json marshal %d %w", chatId.ChatId, err)
			continue
		}

		wp.Submit(func() {
			msg := &sarama.ProducerMessage{
				Topic: constant.TopicK8sHub,
				Value: sarama.ByteEncoder(bytes),
			}

			_, _, err := u.producer.SendMessage(msg)
			if err != nil {
				errChan <- fmt.Errorf("send message %d %w", chatId.ChatId, err)
			}
		})
	}

	wp.StopWait()
	close(errChan)

	var finalErr error

	for err := range errChan {
		finalErr = errors.Join(finalErr, err)
	}

	return finalErr
}
