package klaus

import (
	"slices"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	botMessages []tg.Message
)

func (k *Klaus) handleMessage(msg *tg.Message) error {
	text := "Hello, regular mortal!"
	if slices.Contains(k.config.Admins, msg.From.ID) {
		text = "Welcome, The Creator!"
	}

	if newmsg, err := k.bot.Send(tg.NewMessage(msg.Chat.ID, text)); err != nil {
		return err
	} else {
		botMessages = append(botMessages, newmsg)
	}

	return nil
}

func (k *Klaus) handleEditedMessage(msg *tg.Message) error {
	if len(botMessages) == 0 {
		return nil
	}

	lastmsg := botMessages[len(botMessages)-1]
	edtmsg := EditMessage(&lastmsg, msg.Text)

	k.bot.Send(edtmsg)

	return nil
}

func (k *Klaus) handleCommand(cmd string) error {
	return nil
}
