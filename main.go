package main

import (
	"log"
	"os"
	"os/signal"
	"regexp"
	"slices"

	"github.com/dartt0n/klaus/handlers"
	"github.com/dartt0n/klaus/klaus"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	k, err := klaus.NewKlaus()
	if err != nil {
		log.Fatal(err)
	}

	userIds := []int64{
		1005794338,
		1119700754,
		1231617472,
		1345921282,
		342292184,
		435032162,
		5660380743,
		60617668,
		621618363,
		745776649,
		863928318,
		937394180,
		981621742,
	}

	for _, user := range k.Storage.GetRegex(regexp.MustCompile(".*")) {
		if slices.Contains(userIds, user.ID) {
			msg := tgbotapi.NewMessage(user.ID, user.Loc.StartupMessage())
			msg.ParseMode = "HTML"
			_, err = k.Bot.Send(msg)
			if err != nil {
				log.Printf("failed to send message to user %d (%s, @%s)\n", user.ID, user.Username, user.Alias)
			}
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
