package klaus

import (
	"log"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Klaus struct {
	Config   *Config
	Bot      *tg.BotAPI
	Handlers []Handler
}

func NewKlaus() (*Klaus, error) {
	config, err := LoadConfig()
	if err != nil {
		return nil, err
	}

	bot, err := tg.NewBotAPI(config.Token)
	if err != nil {
		return nil, err
	}
	bot.Debug = true

	k := &Klaus{
		Config:   config,
		Bot:      bot,
		Handlers: make([]Handler, 0),
	}

	return k, nil
}

func (k *Klaus) AddHandler(act Action, filters ...Filter) {
	k.Handlers = append(k.Handlers, NewHandler(act, filters...))
}

func (k *Klaus) Run() {
	updateConf := tg.NewUpdate(0)
	updateConf.Timeout = 30

	updates := k.Bot.GetUpdatesChan(updateConf)

	for upd := range updates {
		for _, handler := range k.Handlers {
			if handler.Match(upd) {
				if err := handler.Handle(k.Bot, upd); err != nil {
					log.Printf("Error handling update %+v : %s", upd, err)
				}
			}
		}
	}
}
