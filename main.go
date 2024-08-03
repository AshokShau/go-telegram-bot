package main

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"goBot/Telegram/config"
	"goBot/Telegram/modules"
	"log"
	"time"
)

func main() {
	if config.Token == "" {
		log.Fatalln("No token provided")
	}

	b, err := gotgbot.NewBot(config.Token, nil)
	if err != nil {
		log.Fatalln(err.Error())
	}

	dispatcher := ext.NewDispatcher(&ext.DispatcherOpts{
		Error: func(b *gotgbot.Bot, ctx *ext.Context, err error) ext.DispatcherAction {
			log.Println("an error occurred while handling update:", err.Error())
			return ext.DispatcherActionNoop
		},
		MaxRoutines: -1,
	})

	updater := ext.NewUpdater(dispatcher, nil)
	if err = updater.StartPolling(b, &ext.PollingOpts{
		DropPendingUpdates:    true,
		EnableWebhookDeletion: true,
		GetUpdatesOpts: &gotgbot.GetUpdatesOpts{
			AllowedUpdates: []string{"message"},
			Timeout:        9,
			RequestOpts: &gotgbot.RequestOpts{
				Timeout: time.Second * 10,
			},
		},
	}); err != nil {
		log.Fatalln(err.Error())
	}

	modules.LoadModules(dispatcher)

	log.Printf("Bot started as %s", b.User.Username)

	updater.Idle()
}
