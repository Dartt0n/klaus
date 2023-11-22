package handlers

import (
	"errors"
	"strconv"

	"github.com/dartt0n/klaus/klaus"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func AddPrefsHandler(k *klaus.Klaus) {
	k.AddHandler(
		func(bot *tg.BotAPI, upd tg.Update) error {
			user_key := strconv.FormatInt(upd.SentFrom().ID, 10)
			user, err := k.Storage.Get(user_key)
			if err != nil {
				return errors.New("Unknown user")
			}

			msgconf := klaus.ReplyMessage(
				upd.Message,
				`Ho ho ho!

Probably, the person who would receive your name will not know about you anything. 

You need to help him/her to prepare a good gift for you!

So, tell me your 1st preference:`,
			)
			msgconf.ReplyMarkup = tg.NewRemoveKeyboard(true) // clean keyboard

			if _, err := bot.Send(msgconf); err != nil {
				return err
			}

			user.State = klaus.StateEnterPref
			k.Storage.Put(user_key, user)

			return nil
		},

		klaus.FilterNewMessage(),
		klaus.FilterUserState(k, klaus.StatePreferences),
		klaus.FilterMsgText(RulesButtonYes),
	)

	k.AddHandler(
		func(bot *tg.BotAPI, upd tg.Update) error {
			msgconf := klaus.ReplyMessage(
				upd.Message,
				`Please, use keyboard buttons :)`,
			)
			msgconf.ReplyMarkup = RulesKeyboard

			if _, err := bot.Send(msgconf); err != nil {
				return err
			}

			return nil
		},

		klaus.FilterNewMessage(),
		klaus.FilterUserState(k, klaus.StatePreferences),
	)
}
