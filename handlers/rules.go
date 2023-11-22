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

			if user.State != klaus.StateRules {
				return errors.New("Bad user state")
			}

			msgconf := klaus.ReplyMessage(
				upd.Message,
				`Great! Then, let me explain the process:

My elves will send you the name of a random person for whom you will prepare a present 🎉
				
How it will work: 
1️⃣ You need to register before the 11th of December (inclusively)
2️⃣ 11th of December at 11:00 AM you'll receive the name of the person you're preparing a gift for
3️⃣ You need to prepare present:
The maximum value of the gift is 500 rubles
You need to make it before 19th of December (inclusive).
Add the note or little postcard, which includes the name of the person for whom the present is.

4️⃣ Bring your gift it to the 319 office

Want to specify preferences about your present?`,
			)
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
			msgconf := klaus.ReplyMessage(
				upd.Message,
				`Please, use keyboard buttons :)`,
			)
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
