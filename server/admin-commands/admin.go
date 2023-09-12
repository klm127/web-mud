package admincommands

import (
	"strings"

	"github.com/pwsdc/web-mud/interfaces/iserver/iactor"
	"github.com/pwsdc/web-mud/server/user/actor/command"
)

var AdminCommands iactor.ICommandGroup

func init() {
	AdminCommands = command.NewCommandGroup("admin")
	list_cmd := command.NewCommand().Name("list").Desc("lists objects instanced on the server").OnExec(list)
	AdminCommands.RegisterCommand(list_cmd.Get())
}

func list(actor iactor.IActor, msg string) {
	msg = strings.ToLower(strings.TrimSpace(msg))
	if len(msg) == 0 {
		actor.MessageSimple("Try 'list actors', 'list rooms', or 'list beings'.")
		return
	}
	front := strings.SplitN(msg, " ", 2)
	if front[0] == "actors" {
		if len(front) > 1 {
			listActors(actor, front[1])
		} else {
			listActors(actor, "")
		}
		return
	}
	actor.Errorf("I didn't understand that.")
}
