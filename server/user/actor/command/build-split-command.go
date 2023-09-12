package command

import "github.com/pwsdc/web-mud/interfaces/iserver/iactor"

// This structure provides methods for functionally building a command object.
type splitCommandBuilder struct {
	command *SplitCommand
}

func NewSplitCommand() *splitCommandBuilder {
	cb := splitCommandBuilder{
		command: &SplitCommand{
			name:        "unnamed",
			aliases:     make([]string, 0),
			description: "this command was never defined",
			onExecute:   stubExecute,
			splits:      make(map[string]iactor.ICommand),
		},
	}
	return &cb
}

func (cb *splitCommandBuilder) Name(val string) *splitCommandBuilder {
	cb.command.name = val
	return cb
}
func (cb *splitCommandBuilder) Alias(val string) *splitCommandBuilder {
	cb.command.aliases = append(cb.command.aliases, val)
	return cb
}
func (cb *splitCommandBuilder) Desc(val string) *splitCommandBuilder {
	cb.command.description = val
	return cb
}
func (cb *splitCommandBuilder) OnExec(callb func(iactor.IActor, string)) *splitCommandBuilder {
	cb.command.onExecute = callb
	return cb
}
func (cb *splitCommandBuilder) SetSplit(key string, command iactor.ICommand) *splitCommandBuilder {
	cb.command.splits[key] = command
	return cb
}

func (cb *splitCommandBuilder) Get() iactor.ICommand {
	return cb.command
}
