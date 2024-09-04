# Reply Keyboard

![reply_keyboard.png](reply_keyboard.png)

## Getting Started

```go
package main

import (
	"context"

	"github.com/yinhe418/telegram-bot"
	"github.com/yinhe418/telegram-bot/models"
	"github.com/yinhe418/telegram-ui/keyboard/reply"
)

var demoReplyKeyboard *reply.ReplyKeyboard

func initReplyKeyboard(b *bot.Bot) {
	demoReplyKeyboard = reply.New(
		b,
		reply.WithPrefix("reply_keyboard"),
		reply.IsSelective(),
		reply.IsOneTimeKeyboard(),
	).
		Button("Button", b, bot.MatchTypeExact, onReplyKeyboardSelect).
		Row().
		Button("Cancel", b, bot.MatchTypeExact, onReplyKeyboardSelect)
}

func handlerReplyKeyboard(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        "Select example command from reply keyboard:",
		ReplyMarkup: demoReplyKeyboard,
	})
}

func onReplyKeyboardSelect(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "You selected: " + string(update.Message.Text),
	})
}

```

## Options

See in [options.go](options.go) file 
