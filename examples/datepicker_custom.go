package main

import (
	"context"
	bot "github.com/yinhe418/telegram-bot"
	"time"

	"github.com/yinhe418/telegram-bot/models"
	"github.com/yinhe418/telegram-ui/datepicker"
)

var demoDatePickerCustom *datepicker.DatePicker

func initDatePickerCustom(b *bot.Bot) {
	excludeDays := []time.Time{
		makeTime(2020, 1, 10),
		makeTime(2020, 1, 13),
		makeTime(2019, 12, 27),
		makeTime(2019, 12, 28),
		makeTime(2019, 12, 29),
	}

	dpOpts := []datepicker.Option{
		datepicker.StartFromSunday(),
		datepicker.CurrentDate(makeTime(2020, 1, 15)),
		datepicker.From(makeTime(2019, 12, 15)),
		datepicker.To(makeTime(2020, 1, 25)),
		datepicker.OnCancel(onDatepickerCustomCancel),
		datepicker.Language("ru"),
		datepicker.Dates(datepicker.DateModeExclude, excludeDays),
		datepicker.WithPrefix("datepicker-custom"),
	}

	demoDatePickerCustom = datepicker.New(b, onDatepickerCustomSelect, dpOpts...)
}

func makeTime(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)
}

func handlerDatepickerCustom(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        "Select the date",
		ReplyMarkup: demoDatePickerCustom,
	})
}

func onDatepickerCustomCancel(ctx context.Context, b *bot.Bot, mes models.MaybeInaccessibleMessage) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: mes.Message.Chat.ID,
		Text:   "Canceled",
	})
}

func onDatepickerCustomSelect(ctx context.Context, b *bot.Bot, mes models.MaybeInaccessibleMessage, date time.Time) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: mes.Message.Chat.ID,
		Text:   "You select " + date.Format("2006-01-02"),
	})
}
