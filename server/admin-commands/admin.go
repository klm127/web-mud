package admincommands

import (
	"github.com/pwsdc/web-mud/interfaces/iserver/iactor"
	"github.com/pwsdc/web-mud/server/user/actor/command"
)

var AdminCommands iactor.ICommandGroup

func init() {
	AdminCommands = command.NewCommandGroup("admin")
}

func list_base(actor iactor.IActor, msg string) {
	actor.MessageSimple("Try 'list actors', 'list rooms', or 'list beings'.")
}
