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

			msgconf := klaus.ReplyMessage(upd.Message, EnterPrefReplyNext)
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

			if len(user.Prefs) == 0 {
				msgconf := klaus.ReplyMessage(upd.Message, EnterPrefReplyZeroPref)
				msgconf.ReplyMarkup = EnterPrefKeyboardEmpty

				if _, err := bot.Send(msgconf); err != nil {
					return err
				}

				user.State = klaus.StateContinuePref
				k.Storage.Put(user_key, user)

				return nil
			}

			msgconf := klaus.ReplyMessage(upd.Message, EnterPrefReplyFinish)
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

			if len(user.Prefs) > 0 {
				user.Prefs = user.Prefs[:len(user.Prefs)-1]
			}

			prefsList := ""
			for _, pref := range user.Prefs {
				prefsList += "• " + pref
			}

			msgconf := klaus.ReplyMessage(upd.Message, EnterPrefReplyRemove+prefsList)
			if len(user.Prefs) == 0 {
				msgconf.ReplyMarkup = EnterPrefKeyboardEmpty
			} else {
				msgconf.ReplyMarkup = EnterPrefKeyboard
			}

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
			msgconf := klaus.ReplyMessage(upd.Message, UnexpectedMessageText)
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
