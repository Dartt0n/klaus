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

			if upd.Message.Text != user.Loc.RulesButtonYes() {
				msgconf := klaus.ReplyMessage(upd.Message, user.Loc.UnexpectedMessageText())
				msgconf.ReplyMarkup = RulesKeyboard(user.Loc)

				if _, err := bot.Send(msgconf); err != nil {
					return err
				}

				return nil
			} else {
				msgconf := klaus.ReplyMessage(upd.Message, user.Loc.PrefIntroMessage())
				msgconf.ReplyMarkup = tg.NewRemoveKeyboard(true) // clean keyboard

				if _, err := bot.Send(msgconf); err != nil {
					return err
				}

				user.State = klaus.StateEnterPref
				k.Storage.Put(user_key, user)

				return nil
			}
		},

		klaus.FilterNewMessage(),
		klaus.FilterUserState(k, klaus.StatePreferences),
	)
}
