package handlers

import (
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	StartButtonYes = "Let's go!"
)

var (
	StartKeyboard = tg.NewReplyKeyboard(
		tg.NewKeyboardButtonRow(
			tg.NewKeyboardButton(StartButtonYes),
		),
	)
)

var (
	RulesButtonYes = "Yes, I'm ready"
)

var (
	RulesKeyboard = tg.NewReplyKeyboard(
		tg.NewKeyboardButtonRow(
			tg.NewKeyboardButton(RulesButtonYes),
		),
	)
)

var (
	EnterPrefButtonContinue = "I want to enter one more preference"
	EnterPrefButtonEnd      = "That's all"
	EnterPrefButtonRemove   = "Remove last preference"
)

var (
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
)
