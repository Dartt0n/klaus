package main

import (
	"log"
	"os"
	"os/signal"
	"regexp"
	"strconv"

	"github.com/dartt0n/klaus/handlers"
	"github.com/dartt0n/klaus/klaus"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	k, err := klaus.NewKlaus()
	if err != nil {
		log.Fatal(err)
	}

	for _, user := range k.Storage.GetRegex(regexp.MustCompile(".*")) {
		giftFor, err := k.Storage.Get(strconv.FormatInt(user.GiftFor, 10))
		if err != nil {
			log.Printf("CRITICAL FAILED TO GET GIFT FOR %v\n", err)
		}

		_, err = k.Bot.Send(tgbotapi.NewMessage(user.ID, user.Loc.GiftForMessage(giftFor.Username, giftFor.Alias, giftFor.Prefs)))
		if err != nil {
			log.Printf("failed to send message to user %d (%s, @%s)\n", user.ID, user.Username, user.Alias)
		}
	}

	handlers.AddLocaleHandler(k)
	handlers.AddStartHandler(k)
	// handlers.AddRulesHandler(k)
	// handlers.AddPrefsHandler(k)
	// handlers.AddEnterPrefsHandler(k)
	// handlers.AddPrefsMenuHandler(k)

	handlers.AddDebugHandler(k)
	// handlers.AddAdminHandlers(k)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		k.Run()
	}()

	<-c
	k.Storage.CreateSnapshot()

}
