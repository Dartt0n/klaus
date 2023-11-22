package handlers

import (
	loc "github.com/dartt0n/klaus/loc"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func StartKeyboard(loc loc.Localization) tg.ReplyKeyboardMarkup {
	return tg.NewReplyKeyboard(
		tg.NewKeyboardButtonRow(
			tg.NewKeyboardButton(loc.StartButtonYes()),
		),
	)
}

func RulesKeyboard(loc loc.Localization) tg.ReplyKeyboardMarkup {
	return tg.NewReplyKeyboard(
		tg.NewKeyboardButtonRow(
			tg.NewKeyboardButton(loc.RulesButtonYes()),
		),
	)
}

func EnterPrefKeyboard(loc loc.Localization) tg.ReplyKeyboardMarkup {
	return tg.NewReplyKeyboard(
		tg.NewKeyboardButtonRow(
			tg.NewKeyboardButton(loc.EnterPrefButtonContinue()),
		),
		tg.NewKeyboardButtonRow(
			tg.NewKeyboardButton(loc.EnterPrefButtonEnd()),
		),
		tg.NewKeyboardButtonRow(
			tg.NewKeyboardButton(loc.EnterPrefButtonRemove()),
		),
	)
}

func EnterPrefKeyboardEmpty(loc loc.Localization) tg.ReplyKeyboardMarkup {
	return tg.NewReplyKeyboard(
		tg.NewKeyboardButtonRow(
			tg.NewKeyboardButton(loc.EnterPrefButtonContinue()),
		),
	)
}
