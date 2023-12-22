package main

import (
	"log"
	"os"
	"os/signal"
	"regexp"

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
		msg := tgbotapi.NewMessage(user.ID, user.Loc.StartupMessage())
		msg.ParseMode = "HTML"
		_, err = k.Bot.Send(msg)
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
