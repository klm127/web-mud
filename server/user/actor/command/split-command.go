package command

import (
	"strings"

	"github.com/pwsdc/web-mud/interfaces/iserver/iactor"
)

/*

	Need an abstraction for splitting the strings

	"SplitCommand"

	baseCommand
		map[string] > SplitCommand | Command | nil

	execute traverses the map
		if it reaches no more strings to split off... it tries to execute the command its on

		if it runs out of commands to traverse, it executes whatever its on with the leftover string

*/

type SplitCommand struct {
	name        string
	aliases     []string
	description string
	onExecute   func(iactor.IActor, string)
	splits      map[string]iactor.ICommand
}

func (c *SplitCommand) Execute(actor iactor.IActor, msg string) {
	msg = strings.ToLower(strings.TrimSpace(msg))
	splits := strings.SplitN(msg, " ", 2)
	branch, ok := c.splits[splits[0]]
	if !ok {
		c.onExecute(actor, msg)
		return
	}
	if len(splits) > 1 {
		branch.Execute(actor, splits[1])
	} else {
		branch.Execute(actor, "")
	}

}

func (c *SplitCommand) GetName() string {
	return c.name
}

func (c *SplitCommand) GetAliases() *[]string {
	return &c.aliases
}

func (c *SplitCommand) GetDescription() string {
	return c.description
}

func (c *SplitCommand) OnExecute(cb func(actor iactor.IActor, msg string)) {
	c.onExecute = cb
}
