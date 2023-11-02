package klaus

import (
	"slices"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func FilterNewMessage() Filter {
	return func(upd tg.Update) bool {
		return upd.Message != nil
	}
}

func FilterEditMessage() Filter {
	return func(upd tg.Update) bool {
		return upd.EditedMessage != nil
	}
}

func FilterFromAdmin(admins []int64) Filter {
	return func(upd tg.Update) bool {
		if upd.Message == nil {
			return false
		}

		if upd.Message.From == nil {
			return false
		}

		return slices.Contains(admins, upd.Message.From.ID)
	}
}

func FilterCommands(command []string) Filter {
	return func(upd tg.Update) bool {
		if upd.Message == nil {
			return false
		}

		return slices.Contains(command, upd.Message.Command())
	}
}
