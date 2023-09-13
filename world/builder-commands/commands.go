package buildercommands

import (
	"github.com/pwsdc/web-mud/interfaces/iserver/iactor"
	"github.com/pwsdc/web-mud/server/user/actor/command"
)

var BuilderCommands iactor.ICommandGroup

func init() {
	BuilderCommands = command.NewCommandGroup("build")
}
