package command

import "github.com/pwsdc/web-mud/interfaces/iserver/iactor"

// This structure provides methods for functionally building a command object.
type commandBuilder struct {
	command *Command
}

func stubExecute(actor iactor.IActor, msg string) {
	actor.Errorf("This command was not properly set up.")
}

func NewCommand() *commandBuilder {
	cb := commandBuilder{
		command: &Command{
			name:        "unnamed",
			aliases:     make([]string, 0),
			description: "this command was never defined",
			onExecute:   stubExecute,
		},
	}
	return &cb
}

func (cb *commandBuilder) Name(val string) *commandBuilder {
	cb.command.name = val
	return cb
}
func (cb *commandBuilder) Alias(val string) *commandBuilder {
	cb.command.aliases = append(cb.command.aliases, val)
	return cb
}
func (cb *commandBuilder) Desc(val string) *commandBuilder {
	cb.command.description = val
	return cb
}
func (cb *commandBuilder) OnExec(callb func(iactor.IActor, string)) *commandBuilder {
	cb.command.onExecute = callb
	return cb
}
func (cb *commandBuilder) Get() iactor.ICommand {
	return cb.command
}
