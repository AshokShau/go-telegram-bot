package modules

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"log"
)

func start(b *gotgbot.Bot, ctx *ext.Context) error {
	msg := ctx.EffectiveMessage

	text := "Hello! \nMost useless bot ever. \n\nUse /help to see available commands."
	_, err := msg.Reply(b, text, &gotgbot.SendMessageOpts{})
	if err != nil {
		log.Printf("[start] Error sending message: %v", err)
		return err
	}

	return ext.EndGroups
}
