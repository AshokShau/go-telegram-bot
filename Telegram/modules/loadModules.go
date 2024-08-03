package modules

import (
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
)

func LoadModules(d *ext.Dispatcher) {
	d.AddHandler(handlers.NewCommand("start", start))
}
