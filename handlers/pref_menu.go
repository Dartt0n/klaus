package handlers

import (
	"errors"
	"strconv"

	"github.com/dartt0n/klaus/klaus"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func AddPrefsMenuHandler(k *klaus.Klaus) {
	k.AddHandler(func(bot *tg.BotAPI, upd tg.Update) error {
		user_key := strconv.FormatInt(upd.SentFrom().ID, 10)
		user, err := k.Storage.Get(user_key)
		if err != nil {
			return errors.New("Unknown user")
		}

		switch upd.Message.Text {
		case user.Loc.EnterPrefButtonContinue():
			return enterPrefContinue(k, bot, upd, user)
		case user.Loc.EnterPrefButtonEnd():
			if len(user.Prefs) == 0 {
				return enterPrefsEndEEmpty(k, bot, upd, user)
			}
			return enterPrefsEndFull(k, bot, upd, user)
		case user.Loc.EnterPrefButtonRemove():
			return enterPrefRemove(k, bot, upd, user)
		default:
			return unexpectedMessage(k, bot, upd, user)
		}
	},

		klaus.FilterNewMessage(),
		klaus.FilterUserState(k, klaus.StateContinuePref),
	)
}

func unexpectedMessage(k *klaus.Klaus, bot *tg.BotAPI, upd tg.Update, user klaus.User) error {
	msgconf := klaus.ReplyMessage(upd.Message, user.Loc.UnexpectedMessageText())
	msgconf.ReplyMarkup = EnterPrefKeyboard(user.Loc)

	if _, err := bot.Send(msgconf); err != nil {
		return err
	}

	return nil
}

func enterPrefRemove(k *klaus.Klaus, bot *tg.BotAPI, upd tg.Update, user klaus.User) error {
	user_key := strconv.FormatInt(upd.SentFrom().ID, 10)

	if len(user.Prefs) > 0 {
		user.Prefs = user.Prefs[:len(user.Prefs)-1]
	}

	prefsList := ""
	for _, pref := range user.Prefs {
		prefsList += "• " + pref + "\n"
	}

	msgconf := klaus.ReplyMessage(upd.Message, user.Loc.EnterPrefReplyRemove()+prefsList)
	if len(user.Prefs) == 0 {
		msgconf.ReplyMarkup = EnterPrefKeyboardEmpty(user.Loc)
	} else {
		msgconf.ReplyMarkup = EnterPrefKeyboard(user.Loc)
	}

	if _, err := bot.Send(msgconf); err != nil {
		return err
	}

	user.State = klaus.StateContinuePref
	k.Storage.Put(user_key, user)

	return nil
}

func enterPrefsEndFull(k *klaus.Klaus, bot *tg.BotAPI, upd tg.Update, user klaus.User) error {
	user_key := strconv.FormatInt(upd.SentFrom().ID, 10)
	msgconf := klaus.ReplyMessage(upd.Message, user.Loc.EnterPrefReplyFinish())
	msgconf.ReplyMarkup = tg.NewRemoveKeyboard(true)

	if _, err := bot.Send(msgconf); err != nil {
		return err
	}

	user.State = klaus.StateWait
	k.Storage.Put(user_key, user)

	return nil
}

func enterPrefsEndEEmpty(k *klaus.Klaus, bot *tg.BotAPI, upd tg.Update, user klaus.User) error {
	user_key := strconv.FormatInt(upd.SentFrom().ID, 10)
	msgconf := klaus.ReplyMessage(upd.Message, user.Loc.EnterPrefReplyZeroPref())
	msgconf.ReplyMarkup = EnterPrefKeyboardEmpty(user.Loc)

	if _, err := bot.Send(msgconf); err != nil {
		return err
	}

	user.State = klaus.StateContinuePref
	k.Storage.Put(user_key, user)

	return nil
}

func enterPrefContinue(k *klaus.Klaus, bot *tg.BotAPI, upd tg.Update, user klaus.User) error {
	user_key := strconv.FormatInt(upd.SentFrom().ID, 10)
	msgconf := klaus.ReplyMessage(upd.Message, user.Loc.EnterPrefReplyNext())
	msgconf.ReplyMarkup = tg.NewRemoveKeyboard(true)

	if _, err := bot.Send(msgconf); err != nil {
		return err
	}

	user.State = klaus.StateEnterPref
	k.Storage.Put(user_key, user)

	return nil
}
