package handlers

import (
	"errors"
	"strconv"

	"github.com/dartt0n/klaus/klaus"
	loc "github.com/dartt0n/klaus/loc"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func AddStartHandler(k *klaus.Klaus) {
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

				if user.Lang == "ru" {
					user.Loc = &loc.RUS
				} else {
					user.Loc = &loc.ENG
				}

				k.Storage.Put(storeKey, user)
			}

			msgconf := klaus.ReplyMessage(upd.Message, user.Loc.StartMessage())
			msgconf.ReplyMarkup = StartKeyboard(user.Loc)

			if _, err := bot.Send(msgconf); err != nil {
				return err
			}

			user.State = klaus.StateRules
			k.Storage.Put(storeKey, user)

			return nil
		},

		// React on new message with /start command
		klaus.FilterNewMessage(),
		klaus.FilterCommands([]string{"start"}),
	)
}
