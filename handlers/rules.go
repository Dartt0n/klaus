package handlers

import (
	"errors"
	"strconv"

	"github.com/dartt0n/klaus/klaus"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func AddRulesHandler(k *klaus.Klaus) {
	k.AddHandler(
		func(bot *tg.BotAPI, upd tg.Update) error {
			tguser := upd.SentFrom()
			if tguser == nil {
				return errors.New("Empty user")
			}

			user_key := strconv.FormatInt(tguser.ID, 10)
			user, err := k.Storage.Get(user_key)
			if err != nil {
				return errors.New("Unknown user")
			}

			msgconf := klaus.ReplyMessage(upd.Message, RulesMessage)
			msgconf.ReplyMarkup = RulesKeyboard

			if _, err := bot.Send(msgconf); err != nil {
				return err
			}

			user.State = klaus.StatePreferences
			k.Storage.Put(user_key, user)

			return nil
		},

		klaus.FilterNewMessage(),
		klaus.FilterUserState(k, klaus.StateRules),
		klaus.FilterMsgText(StartButtonYes),
	)

	k.AddHandler(
		func(bot *tg.BotAPI, upd tg.Update) error {
			msgconf := klaus.ReplyMessage(upd.Message, UnexpectedMessageText)
			msgconf.ReplyMarkup = StartKeyboard

			if _, err := bot.Send(msgconf); err != nil {
				return err
			}

			return nil
		},

		klaus.FilterNewMessage(),
		klaus.FilterUserState(k, klaus.StateRules),
	)
}
