package handlers

import (
	"errors"
	"strconv"

	"github.com/dartt0n/klaus/klaus"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func AddPrefsMenuHandler(k *klaus.Klaus) {
	k.AddHandler(
		func(bot *tg.BotAPI, upd tg.Update) error {
			user_key := strconv.FormatInt(upd.SentFrom().ID, 10)
			user, err := k.Storage.Get(user_key)
			if err != nil {
				return errors.New("Unknown user")
			}

			msgconf := klaus.ReplyMessage(
				upd.Message,
				`Nice! Enter next preference:`,
			)
			msgconf.ReplyMarkup = tg.NewRemoveKeyboard(true)

			if _, err := bot.Send(msgconf); err != nil {
				return err
			}

			user.State = klaus.StateEnterPref
			k.Storage.Put(user_key, user)

			return nil
		},

		klaus.FilterNewMessage(),
		klaus.FilterUserState(k, klaus.StateContinuePref),
		klaus.FilterMsgText(EnterPrefButtonContinue),
	)

	k.AddHandler(
		func(bot *tg.BotAPI, upd tg.Update) error {
			user_key := strconv.FormatInt(upd.SentFrom().ID, 10)
			user, err := k.Storage.Get(user_key)
			if err != nil {
				return errors.New("Unknown user")
			}

			msgconf := klaus.ReplyMessage(
				upd.Message,
				`Cool! We made all preparations! 

Now you should wait for the 11th of December! That day, at 11:00 AM, I'll send you the name of the person you're preparing a gift for! Good luck, my friend!`,
			)
			msgconf.ReplyMarkup = tg.NewRemoveKeyboard(true)

			if _, err := bot.Send(msgconf); err != nil {
				return err
			}

			user.State = klaus.StateWait
			k.Storage.Put(user_key, user)

			return nil
		},

		klaus.FilterNewMessage(),
		klaus.FilterUserState(k, klaus.StateContinuePref),
		klaus.FilterMsgText(EnterPrefButtonEnd),
	)

	k.AddHandler(
		func(bot *tg.BotAPI, upd tg.Update) error {
			user_key := strconv.FormatInt(upd.SentFrom().ID, 10)
			user, err := k.Storage.Get(user_key)
			if err != nil {
				return errors.New("Unknown user")
			}

			user.Prefs = user.Prefs[:len(user.Prefs)-1]

			prefsList := ""
			for _, pref := range user.Prefs {
				prefsList += "• " + pref
			}

			msgconf := klaus.ReplyMessage(
				upd.Message,
				"Ok! Your current list of preferences is:\n\n"+prefsList,
			)
			msgconf.ReplyMarkup = EnterPrefKeyboard

			if _, err := bot.Send(msgconf); err != nil {
				return err
			}

			user.State = klaus.StateContinuePref
			k.Storage.Put(user_key, user)

			return nil
		},

		klaus.FilterNewMessage(),
		klaus.FilterUserState(k, klaus.StateContinuePref),
		klaus.FilterMsgText(EnterPrefButtonRemove),
	)

	k.AddHandler(
		func(bot *tg.BotAPI, upd tg.Update) error {
			msgconf := klaus.ReplyMessage(
				upd.Message,
				`Please, use keyboard buttons :)`,
			)
			msgconf.ReplyMarkup = EnterPrefKeyboard

			if _, err := bot.Send(msgconf); err != nil {
				return err
			}

			return nil
		},

		klaus.FilterNewMessage(),
		klaus.FilterUserState(k, klaus.StatePreferences),
	)
}
