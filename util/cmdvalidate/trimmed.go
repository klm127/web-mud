package cmdvalidate

import (
	"strings"

	"github.com/pwsdc/web-mud/interfaces/iserver/iactor"
)

func Trimmed(cb iactor.CommandFunc) iactor.CommandFunc {
	return func(actor iactor.IActor, msg string) {
		cb(actor, strings.TrimSpace(msg))
	}
}

func Lower(cb iactor.CommandFunc) iactor.CommandFunc {
	return func(actor iactor.IActor, msg string) {
		cb(actor, strings.ToLower(msg))
	}
}

func LowerTrimmed(cb iactor.CommandFunc) iactor.CommandFunc {
	return func(actor iactor.IActor, msg string) {
		cb(actor, strings.ToLower(strings.TrimSpace(msg)))
	}
}
