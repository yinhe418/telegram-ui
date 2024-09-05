package inline

import (
	"strconv"

	"github.com/yinhe418/telegram-bot/models"
)

func (kb *Keyboard) Row() *Keyboard {
	if len(kb.markup[len(kb.markup)-1]) > 0 {
		kb.markup = append(kb.markup, []models.InlineKeyboardButton{})
	}
	return kb
}

func (kb *Keyboard) Button(text string, data []byte, handler OnSelect) *Keyboard {
	kb.handlers = append(kb.handlers, handlerData{
		Handler: handler,
		data:    data,
	})

	kb.markup[len(kb.markup)-1] = append(kb.markup[len(kb.markup)-1], models.InlineKeyboardButton{
		Text:         text,
		CallbackData: kb.prefix + strconv.Itoa(len(kb.handlers)-1),
	})

	return kb
}

func (kb *Keyboard) URL(text string, url string) *Keyboard {
	kb.markup[len(kb.markup)-1] = append(kb.markup[len(kb.markup)-1], models.InlineKeyboardButton{
		Text: text,
		URL:  url,
	})

	return kb
}
