package click_fields

const (
	MessageTableName    = "telegram.messages"
	TgIdColumnName      = "tg_id"
	ChatIdColumnName    = "chat_id"
	MessageColumnName   = "message"
	CreatedAtColumnName = "created_at"
)

// nolint:gochecknoglobals // needed
var (
	InsertMessageColumns = []string{
		TgIdColumnName,
		ChatIdColumnName,
		MessageColumnName,
		CreatedAtColumnName,
	}
)
