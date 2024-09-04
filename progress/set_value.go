package progress

import (
	"context"

	"github.com/yinhe418/telegram-bot"
	"github.com/yinhe418/telegram-bot/models"
)

func (p *Progress) SetValue(ctx context.Context, b *bot.Bot, value float64) {
	if p.canceled {
		return
	}

	p.value = value

	editParams := &bot.EditMessageTextParams{
		ChatID:    p.message.Chat.ID,
		MessageID: p.message.ID,
		Text:      p.renderTextFunc(p.value),
		ParseMode: models.ParseModeMarkdown,
	}
	if p.onCancel != nil {
		editParams.ReplyMarkup = p.buildKeyboard()
	}

	m, err := b.EditMessageText(ctx, editParams)
	if err != nil {
		p.onError(err)
	}

	p.message = m
}
