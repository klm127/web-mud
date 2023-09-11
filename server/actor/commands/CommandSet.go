package commands

import (
	"fmt"

	"github.com/pwsdc/web-mud/shared"
)

// implements ICommandSet
type CommandSet struct {
	name               string
	commands_and_alias map[string]shared.ICommand
	commands_only      map[string]shared.ICommand
}

func (cs *CommandSet) GetName() string {
	return cs.name
}
func (cs *CommandSet) GetCommands() map[string]shared.ICommand {
	return cs.commands_only
}
func (cs *CommandSet) GetKeys() []string {
	keys := make([]string, 0, len(cs.commands_only))
	for k := range cs.commands_and_alias {
		keys = append(keys, k)
	}
	return keys
}
func (cs *CommandSet) HasCommandOrAlias(name string) bool {
	_, ok := cs.commands_and_alias[name]
	return ok
}
func (cs *CommandSet) Getcommand(name string) shared.ICommand {
	c := cs.commands_and_alias[name]
	return c

}
func (cs *CommandSet) Execute(actor shared.IActor, name string, value string) {
	_, ok := cs.commands_and_alias[name]
	if ok {
		cs.commands_and_alias[name].Execute(actor, value)
	}
}
func (cs *CommandSet) RegisterCommand(cmd shared.ICommand) error {
	cname := cmd.GetName()
	c, ok := cs.commands_and_alias[cname]
	if ok {
		return fmt.Errorf("a command named %s is already registed to %s on command set %s", c.GetName(), cname, cs.name)
	}
	aliases := cmd.GetAliases()
	for _, v := range aliases {
		c, ok = cs.commands_and_alias[v]
		if ok {
			return fmt.Errorf("command %s is already registed to alias %s requested by command %s on command set %s", c.GetName(), v, cname, cs.name)
		}
	}
	cs.commands_and_alias[cname] = cmd
	cs.commands_only[cname] = cmd
	for _, v := range aliases {
		cs.commands_and_alias[v] = cmd
	}
	return nil

}
