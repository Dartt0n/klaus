package klaus

import (
	"log"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Klaus struct {
	config *Config
	bot    *tg.BotAPI
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
		config: config,
		bot:    bot,
	}

	return k, nil
}

func (k *Klaus) Run() {
	updateConf := tg.NewUpdate(0)
	updateConf.Timeout = 30

	updates := k.bot.GetUpdatesChan(updateConf)

	for upd := range updates {
		if upd.Message != nil {
			if err := k.HandleMessage(upd.Message); err != nil {
				log.Printf("Error handling message: %s", err.Error())
			}
		}

		if upd.EditedMessage != nil {
			if err := k.HandleEditedMessage(upd.EditedMessage); err != nil {
				log.Printf("Error handling edited message: %s", err.Error())
			}
		}

	}
}
