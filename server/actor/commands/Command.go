package commands

import "github.com/pwsdc/web-mud/shared"

// implements ICommand
type Command struct {
	name        string
	aliases     []string
	description string
	onExecute   func(shared.IActor, string)
}

func NewCommand(name string, description string, alias []string, onExecute func(shared.IActor, string)) *Command {
	cmd := Command{name, alias, description, onExecute}
	return &cmd
}

func (c *Command) GetName() string {
	return c.name
}

func (c *Command) GetAliases() []string {
	return c.aliases
}
func (c *Command) GetDescription() string {
	return c.description
}
func (c *Command) Execute(actor shared.IActor, val string) {
	c.onExecute(actor, val)
}
