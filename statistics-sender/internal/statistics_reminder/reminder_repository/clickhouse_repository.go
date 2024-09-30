package reminder_repository

import (
	"context"

	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	sq "github.com/Masterminds/squirrel"
	"github.com/Xonesent/K8s-Hub/statistics-sender/config"
	click_fields "github.com/Xonesent/K8s-Hub/statistics-sender/internal/database_stores/clickhouse_fields"
)

type ReminderCHRepository struct {
	cfg        *config.Config
	clickhouse driver.Conn
}

func NewReminderCHRepository(cfg *config.Config, clickhouse driver.Conn) *ReminderCHRepository {
	return &ReminderCHRepository{
		cfg:        cfg,
		clickhouse: clickhouse,
	}
}

func (r *ReminderCHRepository) GetUniqueChatIds(ctx context.Context) ([]UniqueChatId, error) {
	query, args, err := sq.Select("DISTINCT " + click_fields.ChatIdColumnName).
		From(click_fields.MessageTableName).
		Where(sq.Expr(click_fields.TgIdColumnName + " = " + click_fields.ChatIdColumnName)).
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return []UniqueChatId{}, err
	}

	var chatIds []UniqueChatId

	if err := r.clickhouse.Select(ctx, &chatIds, query, args...); err != nil {
		return []UniqueChatId{}, err
	}

	return chatIds, nil
}
