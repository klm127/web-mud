package commands

import (
	"github.com/pwsdc/web-mud/interfaces/iserver/iactor"
	"github.com/pwsdc/web-mud/server/user/actor/command"
)

var SenseCommands iactor.ICommandGroup

func init() {
	SenseCommands = command.NewCommandGroup("sense")
	SenseCommands.RegisterCommand(command.NewCommand().Name("look").Alias("l").Desc("turns your eyes to the world around you").OnExec(look).Get())
}

func look(actor iactor.IActor, msg string) {
	actor.ErrorMessage("You are currently blind.")
}
