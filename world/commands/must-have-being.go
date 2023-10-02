package commands

import (
	"strings"

	"github.com/pwsdc/web-mud/interfaces/iserver/iactor"
)

// performs a check to ensure the actor has a being before making the callback. Also trims the message.
func being(cb func(actor iactor.IActor, msg string)) func(actor iactor.IActor, msg string) {
	return func(actor iactor.IActor, msg string) {
		b := actor.Being()
		if b == nil {
			actor.ErrorMessage("You can't find your body.")
			return
		}
		cb(actor, strings.TrimSpace(msg))
	}
}
