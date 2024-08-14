package buttons_repository

import (
	"context"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	sq "github.com/Masterminds/squirrel"
	"github.com/Xonesent/K8s-Hub/telegram-bot/config"
	click_fields "github.com/Xonesent/K8s-Hub/telegram-bot/internal/database_stores/clickhouse_fields"
	"time"
)

type ButtonsCHRepository struct {
	cfg        *config.Config
	clickhouse driver.Conn
}

func NewButtonsCHRepository(cfg *config.Config, clickhouse driver.Conn) *ButtonsCHRepository {
	return &ButtonsCHRepository{
		cfg:        cfg,
		clickhouse: clickhouse,
	}
}

func (r *ButtonsCHRepository) InsertLastMessage(ctx context.Context, sentMessage *SentMessage) error {
	query, args, err := sq.Insert(click_fields.MessageTableName).
		Columns(click_fields.InsertMessageColumns...).
		Values(
			sentMessage.Sender,
			sentMessage.ChatId,
			sentMessage.Message,
			time.Now(),
		).
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return err
	}

	if err := r.clickhouse.AsyncInsert(ctx, query, false, args...); err != nil {
		return err
	}

	return nil
}
