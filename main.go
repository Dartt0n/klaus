package main

import (
	"log"

	"github.com/dartt0n/klaus/klaus"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	k, err := klaus.NewKlaus()
	if err != nil {
		log.Fatal(err)
	}
	k.AddHandler(
		func(bot *tg.BotAPI, upd tg.Update) error {
			_, err := bot.Send(klaus.ReplyMessage(
				upd.Message,
				"{start message}",
			))

			if err != nil {
				return err
			}

			return nil
		},

		// React on new message with /start command
		klaus.FilterNewMessage(),
		klaus.FilterCommands([]string{"start"}),
	)

	k.AddHandler(
		func(bot *tg.BotAPI, upd tg.Update) error {
			_, err := bot.Send(klaus.ReplyMessage(
				upd.Message,
				"{message from admin}",
			))

			if err != nil {
				return err
			}

			return nil
		},

		// React on new message from admin without commands
		klaus.FilterNewMessage(),
		klaus.FilterFromAdmin(k.Config.Admins),
		klaus.FilterCommands([]string{""}),
	)

	k.AddHandler(
		func(bot *tg.BotAPI, upd tg.Update) error {
			_, err := bot.Send(klaus.ReplyMessage(
				upd.Message,
				"{edit message from admin}",
			))

			if err != nil {
				return err
			}

			return nil
		},

		klaus.FilterEditMessage(),
		klaus.FilterFromAdmin(k.Config.Admins),
	)

	k.Run()
}
