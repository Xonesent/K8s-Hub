package buttons_usecase

import (
	"context"

	"github.com/Xonesent/K8s-Hub/telegram-bot/config"
	tg_resp "github.com/Xonesent/K8s-Hub/telegram-bot/pkg/predefined_responses/telegram"
	tg_utils "github.com/Xonesent/K8s-Hub/telegram-bot/pkg/utilities/telegram"
	"github.com/go-telegram/bot"
)

type ButtonsUseCase struct {
	cfg           *config.Config
	buttonsCHRepo ButtonsCHRepo
	b             *bot.Bot
}

func NewButtonsUseCase(cfg *config.Config, buttonsCHRepo ButtonsCHRepo, b *bot.Bot) *ButtonsUseCase {
	return &ButtonsUseCase{
		cfg:           cfg,
		buttonsCHRepo: buttonsCHRepo,
		b:             b,
	}
}

func (u *ButtonsUseCase) DefaultHandler(ctx context.Context, sentMessage *SentMessage) error {
	sentMessageDTO := sentMessage.toSentMessage()
	if err := u.buttonsCHRepo.InsertLastMessage(ctx, &sentMessageDTO); err != nil {
		return err
	}

	sendMessageDTO := tg_utils.FormSimpleResponse(sentMessage.ChatId, tg_resp.DefaultResponse)
	if _, err := u.b.SendMessage(ctx, &sendMessageDTO); err != nil {
		return err
	}

	return nil
}
