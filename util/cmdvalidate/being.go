package cmdvalidate

import (
	"strings"

	"github.com/pwsdc/web-mud/interfaces/iserver/iactor"
)

// validates actor has a being and room
func BeingAndRoom(cb iactor.CommandFunc) iactor.CommandFunc {
	return func(actor iactor.IActor, msg string) {
		if actor.Being() == nil {
			actor.Errorf("You can't find your body.")
			return
		}
		if actor.Being().GetRoom() == nil {
			actor.Errorf("You can't figure out where you are.")
			return
		}
		cb(actor, msg)
	}
}

// trims message and validates actor has a being and room
func BeingTrimLower(cb iactor.CommandFunc) iactor.CommandFunc {
	return func(actor iactor.IActor, msg string) {
		msg = strings.ToLower(strings.TrimSpace(msg))
		if actor.Being() == nil {
			actor.Errorf("You can't find your body.")
			return
		}
		if actor.Being().GetRoom() == nil {
			actor.Errorf("You can't figure out where you are.")
			return
		}
		cb(actor, msg)
	}
}
