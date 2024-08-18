package admin_repository

import (
	"context"

	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	sq "github.com/Masterminds/squirrel"
	"github.com/Xonesent/K8s-Hub/admin-panel/config"
	models "github.com/Xonesent/K8s-Hub/admin-panel/internal/business_models"
	click_fields "github.com/Xonesent/K8s-Hub/admin-panel/internal/database_stores/clickhouse_fields"
	errlst "github.com/Xonesent/K8s-Hub/admin-panel/pkg/predefined_responses/error_list"
)

type AdminCHRepository struct {
	cfg        *config.Config
	clickhouse driver.Conn
}

func NewAdminCHRepository(cfg *config.Config, clickhouse driver.Conn) *AdminCHRepository {
	return &AdminCHRepository{
		cfg:        cfg,
		clickhouse: clickhouse,
	}
}

func (r *AdminCHRepository) GetUserStatistics(ctx context.Context, tgId models.TgId) (models.UserLog, error) {
	query, args, err := sq.Select(click_fields.MessageColumns...).
		From(click_fields.MessageTableName).
		Where(sq.Eq{click_fields.TgIdColumnName: tgId}).
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return models.UserLog{}, errlst.HttpServerError
	}

	var userMessages []UserMessage

	if err := r.clickhouse.Select(ctx, &userMessages, query, args...); err != nil {
		return models.UserLog{}, errlst.HttpServerError
	}

	if len(userMessages) == 0 {
		return models.UserLog{}, errlst.HttpErrNotFound
	}

	return toUserLog(userMessages), nil
}
