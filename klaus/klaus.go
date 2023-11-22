package klaus

import (
	"log"

	"github.com/dartt0n/skhron"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Klaus struct {
	Config   *Config
	Bot      *tg.BotAPI
	Handlers []Handler

	Storage *skhron.Skhron[User]
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
		Storage:  skhron.New[User](),
	}

	if err := k.Storage.LoadSnapshot(); err != nil {
		log.Println(err)
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
				break
			}
		}
	}
}
