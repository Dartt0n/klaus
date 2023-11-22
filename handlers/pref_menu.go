package handlers

import (
	"errors"
	"strconv"

	"github.com/dartt0n/klaus/klaus"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// todo: refactor
func AddPrefsMenuHandler(k *klaus.Klaus) {
	k.AddHandler(
		func(bot *tg.BotAPI, upd tg.Update) error {
			user_key := strconv.FormatInt(upd.SentFrom().ID, 10)
			user, err := k.Storage.Get(user_key)
			if err != nil {
				return errors.New("Unknown user")
			}

			switch upd.Message.Text {
			case user.Loc.EnterPrefButtonContinue():
				msgconf := klaus.ReplyMessage(upd.Message, user.Loc.EnterPrefReplyNext())
				msgconf.ReplyMarkup = tg.NewRemoveKeyboard(true)

				if _, err := bot.Send(msgconf); err != nil {
					return err
				}

				user.State = klaus.StateEnterPref
				k.Storage.Put(user_key, user)

				return nil
			case user.Loc.EnterPrefButtonEnd():
				if len(user.Prefs) == 0 {
					msgconf := klaus.ReplyMessage(upd.Message, user.Loc.EnterPrefReplyZeroPref())
					msgconf.ReplyMarkup = EnterPrefKeyboardEmpty(user.Loc)

					if _, err := bot.Send(msgconf); err != nil {
						return err
					}

					user.State = klaus.StateContinuePref
					k.Storage.Put(user_key, user)

					return nil
				} else {
					msgconf := klaus.ReplyMessage(upd.Message, user.Loc.EnterPrefReplyFinish())
					msgconf.ReplyMarkup = tg.NewRemoveKeyboard(true)

					if _, err := bot.Send(msgconf); err != nil {
						return err
					}

					user.State = klaus.StateWait
					k.Storage.Put(user_key, user)

					return nil
				}
			case user.Loc.EnterPrefButtonRemove():
				if len(user.Prefs) > 0 {
					user.Prefs = user.Prefs[:len(user.Prefs)-1]
				}

				prefsList := ""
				for _, pref := range user.Prefs {
					prefsList += "• " + pref
				}

				msgconf := klaus.ReplyMessage(upd.Message, user.Loc.EnterPrefReplyRemove()+prefsList)
				if len(user.Prefs) == 0 {
					msgconf.ReplyMarkup = EnterPrefKeyboardEmpty(user.Loc)
				} else {
					msgconf.ReplyMarkup = EnterPrefKeyboard(user.Loc)
				}

				if _, err := bot.Send(msgconf); err != nil {
					return err
				}

				user.State = klaus.StateContinuePref
				k.Storage.Put(user_key, user)

				return nil
			default:
				msgconf := klaus.ReplyMessage(upd.Message, user.Loc.UnexpectedMessageText())
				msgconf.ReplyMarkup = EnterPrefKeyboard(user.Loc)

				if _, err := bot.Send(msgconf); err != nil {
					return err
				}

				return nil
			}
		},

		klaus.FilterNewMessage(),
		klaus.FilterUserState(k, klaus.StateContinuePref),
	)
}
