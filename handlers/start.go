package handlers

import (
	"errors"
	"strconv"

	"github.com/dartt0n/klaus/klaus"
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
				k.Storage.Put(storeKey, user)
			}

			msgconf := klaus.ReplyMessage(
				upd.Message,
				`Hello, my friend! It's me, Santa! And I'm glad to see you here! 🎅

My clever elves decided to help me with presents for kind people in Innopolis University. 

They created this bot where you can participate in sharing wonderful vibes through your gifts.  

Are you ready for a miracle?`,
			)
			msgconf.ReplyMarkup = StartKeyboard

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
