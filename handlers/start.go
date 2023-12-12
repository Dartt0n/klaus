package handlers

import (
	"log"
	"strconv"

	"github.com/dartt0n/klaus/klaus"
	"github.com/dartt0n/klaus/loc"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func AddStartHandler(k *klaus.Klaus) {
	k.AddHandler(
		func(bot *tg.BotAPI, upd tg.Update) error {
			userKey := strconv.FormatInt(upd.SentFrom().ID, 10)
			user, _ := k.Storage.Get(userKey)

			if upd.CallbackQuery.Data == loc.ENG.Lang {
				user.Lang = loc.ENG.Lang
				user.Loc = &loc.ENG
			} else if upd.CallbackQuery.Data == loc.RUS.Lang {
				user.Lang = loc.RUS.Lang
				user.Loc = &loc.RUS
			} else {
				user.Lang = loc.ENG.Lang
				user.Loc = &loc.ENG
			}

			if _, err := bot.Request(tg.NewCallback(upd.CallbackQuery.ID, upd.CallbackQuery.Data)); err != nil {
				return err
			}

			if len(user.Prefs) == 0 {
				msgconf := tg.NewMessage(upd.CallbackQuery.Message.Chat.ID, user.Loc.RegistrationClosed())
				msgconf.ReplyMarkup = tg.NewRemoveKeyboard(true)
				if _, err := bot.Send(msgconf); err != nil {
					return err
				}

				return nil
			}

			giftFor, err := k.Storage.Get(strconv.FormatInt(user.GiftFor, 10))
			if err != nil {
				return err
			}
			log.Printf("START")

			msgconf := tg.NewMessage(upd.CallbackQuery.Message.Chat.ID, user.Loc.InfoGiftMessage(giftFor.Username, giftFor.Alias, giftFor.Prefs))
			msgconf.ReplyMarkup = tg.NewRemoveKeyboard(true)
			if _, err := bot.Send(msgconf); err != nil {
				return err
			}

			return nil
		},

		klaus.FilterCallbackQuery(),
		klaus.FilterUserState(k, klaus.StateLocale),
	)
}
