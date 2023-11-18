package handlers

import (
	"encoding/json"
	"strconv"

	"github.com/dartt0n/klaus/klaus"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func AddDebugHandler(k *klaus.Klaus) {
	k.AddHandler(
		func(bot *tg.BotAPI, upd tg.Update) error {
			key := strconv.FormatInt(upd.SentFrom().ID, 10)

			k.Storage.CreateSnapshot()

			if value, err := k.Storage.Get(key); err != nil {
				bot.Send(klaus.ReplyMessage(
					upd.Message,
					"User not found",
				))
			} else {
				jsonBytes, _ := json.Marshal(value)

				bot.Send(klaus.ReplyMessage(
					upd.Message,
					"```json\n"+string(jsonBytes)+"```",
				))
			}

			return nil
		},

		// React on new message with /start command
		klaus.FilterNewMessage(),
		klaus.FilterCommands([]string{"debug"}),
	)
}
