package command

import "github.com/pwsdc/web-mud/server/actor/base"

type Command struct {
	Name        string
	Alias       []string
	Description string
	OnExecute   func(*base.Actor, string)
}

func NewCommand(name string, description string, alias []string, onExecute func(*base.Actor, string)) *Command {
	cmd := Command{name, alias, description, onExecute}
	return &cmd
}
