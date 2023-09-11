package user

import (
	"github.com/pwsdc/web-mud/server/actor/base"
	"github.com/pwsdc/web-mud/server/actor/prompts"
)

var LoggedOutCommands *base.CommandSet
var LoggedInCommands *base.CommandSet

func init() {
	lo := base.NewCommandSet("user")
	reg_cmd := base.NewCommand("register", "registers an account", nil, register)
	login_cmd := base.NewCommand("login", "logs in to your account", nil, login)
	lo.RegisterCommand(reg_cmd)
	lo.RegisterCommand(login_cmd)
	LoggedOutCommands = lo
	base.RegisterDefaultCommandSet(lo)
	prompts.LoggedOutCommands = lo

	li := base.NewCommandSet("user")
	logout_cmd := base.NewCommand("logout", "logout of your account", nil, logout)
	li.RegisterCommand(logout_cmd)
	LoggedInCommands = li
	prompts.LoggedInCommands = li
}

func register(actor *base.Actor, msg string) {
	actor.StartQuestioning(prompts.RegisterQuestions)
}

func login(actor *base.Actor, msg string) {
	actor.StartQuestioning(prompts.LoginQuestions)
}

func logout(actor *base.Actor, msg string) {
	actor.RemoveUser()
	actor.SetCommandGroup("user", LoggedOutCommands)
	actor.MessageSimplef("You are logged out.")
}
