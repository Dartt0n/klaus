package klaus

import (
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func EditMessage(msg *tg.Message, text string) tg.EditMessageTextConfig {
	return tg.EditMessageTextConfig{
		BaseEdit: tg.BaseEdit{
			ChatID:    msg.Chat.ID,
			MessageID: msg.MessageID,
		},
		Text: text,
	}
}
