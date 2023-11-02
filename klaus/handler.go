package klaus

import (
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Filter = func(upd tg.Update) bool
type Action = func(bot *tg.BotAPI, upd tg.Update) error

type Handler interface {
	Match(tg.Update) bool
	Handle(*tg.BotAPI, tg.Update) error
}

type handler struct {
	filters []Filter
	action  func(*tg.BotAPI, tg.Update) error
}

func NewHandler(action Action, filters ...Filter) *handler {
	h := &handler{}
	h.filters = filters
	h.action = action
	return h
}

func (h *handler) Match(upd tg.Update) bool {
	for _, f := range h.filters {
		if !f(upd) {
			return false
		}
	}

	return true
}

func (h *handler) Handle(bot *tg.BotAPI, upd tg.Update) error {
	return h.action(bot, upd)
}
