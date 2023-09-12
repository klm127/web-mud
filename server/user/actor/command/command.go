package command

import "github.com/pwsdc/web-mud/interfaces/iserver/iactor"

type Command struct {
	name        string
	aliases     []string
	description string
	onExecute   func(iactor.IActor, string)
}

func (c *Command) GetName() string {
	return c.name
}

func (c *Command) GetAliases() *[]string {
	return &c.aliases
}

func (c *Command) GetDescription() string {
	return c.description
}

func (c *Command) OnExecute(cb func(actor iactor.IActor, msg string)) {
	c.onExecute = cb
}

func (c *Command) Execute(actor iactor.IActor, msg string) {
	c.onExecute(actor, msg)
}
