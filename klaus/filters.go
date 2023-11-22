package klaus

import (
	"slices"
	"strconv"

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
		if upd.SentFrom() == nil {
			return false
		}

		return slices.Contains(admins, upd.SentFrom().ID)
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

func FilterMsgText(text string) Filter {
	return func(upd tg.Update) bool {
		if upd.Message == nil {
			return false
		}

		return upd.Message.Text == text
	}
}

func FilterUserState(k *Klaus, state int) Filter {
	return func(upd tg.Update) bool {
		tguser := upd.SentFrom()
		if tguser == nil {
			return false
		}

		user_key := strconv.FormatInt(tguser.ID, 10)
		user, err := k.Storage.Get(user_key)
		if err != nil {
			return false
		}

		return user.State == state
	}
}

func FilterEmptyMessage() Filter {
	return func(upd tg.Update) bool {
		if upd.Message == nil {
			return false
		}

		return upd.Message.Text == ""
	}
}

func FilterNonEmptyMessage() Filter {
	return func(upd tg.Update) bool {
		if upd.Message == nil {
			return false
		}

		return upd.Message.Text != ""
	}
}

func FilterCallbackQuery() Filter {
	return func(upd tg.Update) bool {
		return upd.CallbackQuery != nil
	}
}
