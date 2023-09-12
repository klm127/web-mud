package command

import (
	"errors"
	"fmt"

	"github.com/pwsdc/web-mud/interfaces/iserver/iactor"
)

// Implements iactor.ICommandSet
type commandGroup struct {
	name                 string
	commands_and_aliases map[string]iactor.ICommand
	commands_only        map[string]iactor.ICommand
}

func NewCommandGroup(cset_name string) iactor.ICommandGroup {
	return &commandGroup{
		name:                 cset_name,
		commands_and_aliases: make(map[string]iactor.ICommand),
		commands_only:        make(map[string]iactor.ICommand),
	}
}

func (cs *commandGroup) GetName() string {
	return cs.name
}

func (cs *commandGroup) GetCommands() *map[string]iactor.ICommand {
	return &cs.commands_only
}

func (cs *commandGroup) GetKeys() *[]string {
	keys := make([]string, 0, len(cs.commands_only))
	for k := range cs.commands_and_aliases {
		keys = append(keys, k)
	}
	return &keys

}
func (cs *commandGroup) HasCommandOrAlias(name string) bool {
	_, ok := cs.commands_and_aliases[name]
	return ok

}

func (cs *commandGroup) GetCommand(name string) iactor.ICommand {
	c := cs.commands_and_aliases[name]
	return c
}

func (cs *commandGroup) Execute(actor iactor.IActor, name string, value string) {
	_, ok := cs.commands_and_aliases[name]
	if ok {
		cs.commands_and_aliases[name].Execute(actor, value)
	}

}
func (cs *commandGroup) RegisterCommand(cmd iactor.ICommand) error {
	cname := cmd.GetName()
	c, ok := cs.commands_and_aliases[cname]
	if ok {
		fname := c.GetName()
		return fmt.Errorf("a command named %s is already registed to %s on command set %s", fname, cname, cs.name)
	}
	aliases := *cmd.GetAliases()
	for _, v := range aliases {
		c, ok = cs.commands_and_aliases[v]
		if ok {
			return fmt.Errorf("command %s is already registed to alias %s requested by command %s on command set %s", cname, v, cname, cs.name)
		}
	}
	cs.commands_and_aliases[cname] = cmd
	cs.commands_only[cname] = cmd
	for _, v := range aliases {
		cs.commands_and_aliases[v] = cmd
	}
	return nil
}

func (cs *commandGroup) RegisterCommands(cmds ...iactor.ICommand) error {
	erstring := ""
	for _, v := range cmds {
		err := cs.RegisterCommand(v)
		if err != nil {
			erstring += "..." + err.Error()
		}
	}
	return errors.New(erstring)
}
