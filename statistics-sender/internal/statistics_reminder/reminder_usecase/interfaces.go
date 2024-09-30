package reminder_usecase

import (
	"context"

	"github.com/Xonesent/K8s-Hub/statistics-sender/internal/statistics_reminder/reminder_repository"
)

type ReminderCHRepo interface {
	GetUniqueChatIds(ctx context.Context) ([]reminder_repository.UniqueChatId, error)
}
