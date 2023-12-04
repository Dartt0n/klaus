package handlers

import (
	"errors"
	"strconv"

	"github.com/dartt0n/klaus/klaus"
	loc "github.com/dartt0n/klaus/loc"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func AddLocaleHandler(k *klaus.Klaus) {
	k.AddHandler(
		func(bot *tg.BotAPI, upd tg.Update) error {
			tguser := upd.SentFrom()
			if tguser == nil {
				return errors.New("Empty user")
			}

			storeKey := strconv.FormatInt(tguser.ID, 10)

			var user klaus.User

			if value, err := k.Storage.Get(storeKey); err == nil {
				user = value
			} else {
				user = klaus.NewUser(upd)

				k.Storage.Put(storeKey, user)
			}

			msgconf := klaus.ReplyMessage(upd.Message, "Please select a language:")
			msgconf.ReplyMarkup = tg.NewInlineKeyboardMarkup(
				tg.NewInlineKeyboardRow(
					tg.NewInlineKeyboardButtonData("🇷🇺 Russian", loc.RUS.Lang),
				),
				tg.NewInlineKeyboardRow(
					tg.NewInlineKeyboardButtonData("🇺🇸 English", loc.ENG.Lang),
				),
			)

			if _, err := bot.Send(msgconf); err != nil {
				return err
			}

			user.Prefs = []string{}
			user.State = klaus.StateLocale
			k.Storage.Put(storeKey, user)

			return nil
		},

		klaus.FilterNewMessage(),
		klaus.FilterCommands([]string{"start"}),
	)
}
