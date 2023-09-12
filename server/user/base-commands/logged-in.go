package basecommands

import (
	"github.com/pwsdc/web-mud/interfaces/iserver/iactor"
	"github.com/pwsdc/web-mud/server/user/actor/command"
)

var UserLoggedInCommands iactor.ICommandGroup

func init() {
	UserLoggedInCommands = command.NewCommandGroup("user")
	logout_cmd := command.NewCommand().Name("logout").Desc("logs you out of your account.").OnExec(logout).Get()
	UserLoggedInCommands.RegisterCommand(logout_cmd)
}

func logout(actor iactor.IActor, msg string) {
	actor.RemoveUser()
	actor.SetCommandGroup(UserLoggedOutCommands)
	actor.MessageSimplef("You have been logged out.")
}
