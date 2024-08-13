package telegram

import "github.com/go-telegram/bot"

type ConfigTelegram struct {
	Token string `yaml:"Token" validate:"required"`
}

func NewTelegramBot(cfg ConfigTelegram) (*bot.Bot, error) {
	opts := []bot.Option{}

	b, err := bot.New(cfg.Token, opts...)
	if err != nil {
		return nil, err
	}

	return b, err
}
