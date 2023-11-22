package handlers

import (
	"errors"
	"strconv"

	"github.com/dartt0n/klaus/klaus"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func AddEnterPrefsHandler(k *klaus.Klaus) {
	k.AddHandler(
		func(bot *tg.BotAPI, upd tg.Update) error {
			user_key := strconv.FormatInt(upd.SentFrom().ID, 10)
			user, err := k.Storage.Get(user_key)
			if err != nil {
				return errors.New("Unknown user")
			}

			if upd.Message.Text == EnterPrefButtonContinue ||
				upd.Message.Text == EnterPrefButtonEnd ||
				upd.Message.Text == EnterPrefButtonRemove ||
				upd.Message.Text == RulesButtonYes {
				return errors.New("Multiple button clicks detected")
			}

			user.Prefs = append(user.Prefs, upd.Message.Text)

			prefsList := ""
			for _, pref := range user.Prefs {
				prefsList += "• " + pref + "\n"
			}

			msgconf := klaus.ReplyMessage(
				upd.Message,
				EnterPrefReplySuccess+prefsList,
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
		klaus.FilterNonEmptyMessage(),
		klaus.FilterUserState(k, klaus.StateEnterPref),
	)

	k.AddHandler(
		func(bot *tg.BotAPI, upd tg.Update) error {
			msgconf := klaus.ReplyMessage(
				upd.Message,
				EnterPrefReplyUnkwonMessage,
			)
			msgconf.ReplyMarkup = tg.NewRemoveKeyboard(true)

			if _, err := bot.Send(msgconf); err != nil {
				return err
			}

			return nil
		},

		klaus.FilterNewMessage(),
		klaus.FilterEmptyMessage(),
		klaus.FilterUserState(k, klaus.StateEnterPref),
	)
}
