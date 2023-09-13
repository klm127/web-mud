package buildercommands

import (
	"github.com/pwsdc/web-mud/interfaces/iserver/iactor"
	"github.com/pwsdc/web-mud/server/user/actor/command"
	"github.com/pwsdc/web-mud/util/cmdvalidate"
	"github.com/pwsdc/web-mud/util/language"
)

func init() {
	exit := command.NewCommand().Name("exit").Desc("Creates a new exit for this room.").OnExec(cmdvalidate.BeingTrimLower(createExit)).Get()
	BuilderCommands.RegisterCommand(exit)
}

func createExit(actor iactor.IActor, msg string) {
	if len(msg) == 0 {
		actor.ErrorMessage("You must specify a direction, like n, s, e, w, ne, nw, se, sw, in, out, up, or down.")
		return
	}
	direction, ok := language.ParseDirection(msg)
	if !ok {
		actor.Errorf("%s is not a direction I know.", msg)
		return
	}
	actor.MessageSimplef("Soom, you will be able to create a room to the %s.", direction)
}

/*
	build.set room name <name>
	build.set room desc <desc>
	build.exit <direction>

*/
