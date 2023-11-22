package handlers

import (
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	StartKeyboard = tg.NewReplyKeyboard(
		tg.NewKeyboardButtonRow(
			tg.NewKeyboardButton(StartButtonYes),
		),
	)

	RulesKeyboard = tg.NewReplyKeyboard(
		tg.NewKeyboardButtonRow(
			tg.NewKeyboardButton(RulesButtonYes),
		),
	)

	EnterPrefKeyboard = tg.NewReplyKeyboard(
		tg.NewKeyboardButtonRow(
			tg.NewKeyboardButton(EnterPrefButtonContinue),
		),
		tg.NewKeyboardButtonRow(
			tg.NewKeyboardButton(EnterPrefButtonEnd),
		),
		tg.NewKeyboardButtonRow(
			tg.NewKeyboardButton(EnterPrefButtonRemove),
		),
	)

	EnterPrefKeyboardEmpty = tg.NewReplyKeyboard(
		tg.NewKeyboardButtonRow(
			tg.NewKeyboardButton(EnterPrefButtonContinue),
		),
	)
)
