package command

import (
	"fmt"

	"github.com/pwsdc/web-mud/server/actor/base"
)

var BaseCommandSets map[string]*CommandSet

type CommandSet struct {
	Name     string
	commands map[string]*Command
}

func init() {
	BaseCommandSets = make(map[string]*CommandSet)
}

func NewCommandSet(name string) *CommandSet {
	cs := CommandSet{}
	cs.Name = name
	cs.commands = make(map[string]*Command)
	return &cs
}

func RegisterBaseCommandSet(cs *CommandSet) error {
	_, ok := BaseCommandSets[cs.Name]
	if ok {
		return fmt.Errorf("a command named %s is already registered", cs.Name)
	}
	return nil
}

func (cs *CommandSet) HasCommand(name string) bool {
	_, ok := cs.commands[name]
	return ok
}

func (cs *CommandSet) Execute(actor *base.Actor, name string, value string) {
	_, ok := cs.commands[name]
	if ok {
		cs.commands[name].OnExecute(actor, value)
	}
}

func (cs *CommandSet) RegisterCommand(cmd *Command) error {
	c, ok := cs.commands[cmd.Name]
	if ok {
		return fmt.Errorf("a command named %s is already registed to %s on command set %s", c.Name, cmd.Name, cs.Name)
	}
	for _, v := range cmd.Alias {
		c, ok = cs.commands[v]
		if ok {
			return fmt.Errorf("command %s is already registed to alias %s requested by command %s on command set %s", c.Name, v, cmd.Name, cs.Name)
		}
	}
	cs.commands[cmd.Name] = cmd
	for _, v := range cmd.Alias {
		cs.commands[v] = cmd
	}
	return nil
}
