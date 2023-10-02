package commands

import (
	"github.com/pwsdc/web-mud/interfaces/iserver/iactor"
	"github.com/pwsdc/web-mud/server/user/actor/command"
	"github.com/pwsdc/web-mud/server/user/actor/message"
)

var SenseCommands iactor.ICommandGroup

func init() {
	SenseCommands = command.NewCommandGroup("sense")
	SenseCommands.RegisterCommand(command.NewCommand().Name("look").Alias("l").Desc("turns your eyes to the world around you").OnExec(being(look)).Get())
}

func look(actor iactor.IActor, msg string) {
	if len(msg) == 0 {
		actor.Being().SeeRoom()
		return
	}
	if msg == "me" || msg == "self" {
		msgb := message.New().Text("You see yourself.").NewLine(1).Next()
		msgb.Text(actor.Being().Desc())
		actor.Message(msgb.Bytes())
		return
	}
	// implement looking at objects around
	actor.ErrorMessage("You are currently blind.")
}
