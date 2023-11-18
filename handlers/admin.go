package handlers

import (
	"fmt"
	"regexp"

	"github.com/dartt0n/klaus/klaus"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func AddAdminHandlers(k *klaus.Klaus) {
	k.AddHandler(
		func(bot *tg.BotAPI, upd tg.Update) error {
			users := k.Storage.GetRegex(regexp.MustCompile("^\\d+$"))
			ids := make([]int64, 0, len(users))
			for _, user := range users {
				ids = append(ids, user.ID)
			}

			_, err := bot.Send(klaus.ReplyMessage(
				upd.Message,
				fmt.Sprintf("{message from admin, would be resend to %v}", ids),
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
				upd.EditedMessage,
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
}
