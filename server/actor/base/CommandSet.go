package base

import (
	"fmt"
)

/*
Command sets in this map will be added to every actor upon creation.
*/
var defaultCommandSets map[string]*CommandSet

/*
Associates a map of commands with a name for the group.
*/
type CommandSet struct {
	Name               string
	commands_and_alias map[string]*Command
	commands_only      map[string]*Command
}

func init() {
	defaultCommandSets = make(map[string]*CommandSet)
}

/*
Get a new command set.
*/
func NewCommandSet(name string) *CommandSet {
	cs := CommandSet{}
	cs.Name = name
	cs.commands_and_alias = make(map[string]*Command)
	cs.commands_only = make(map[string]*Command)
	return &cs
}

/*
Register a command set as a default one that should be added to every actor on creation.
*/
func RegisterDefaultCommandSet(cs *CommandSet) error {
	_, ok := defaultCommandSets[cs.Name]
	if ok {
		return fmt.Errorf("a command named %s is already registered", cs.Name)
	}
	defaultCommandSets[cs.Name] = cs
	return nil
}

/*
Get the map of commands only associated with this command set.

Will be in the form of [commandName] -> command

Aliases are not items in this map.
*/
func (cs *CommandSet) GetCommands() map[string]*Command {
	return cs.commands_only
}

/*
Get only the keys of the map (the names of each command and alias.)
*/
func (cs *CommandSet) GetKeys() []string {
	keys := make([]string, 0, len(cs.commands_only))
	for k := range cs.commands_and_alias {
		keys = append(keys, k)
	}
	return keys
}

/*
Determine whether param name exists as a a key in the the commands map.
*/
func (cs *CommandSet) HasCommandOrAlias(name string) bool {
	_, ok := cs.commands_and_alias[name]
	return ok
}

func (cs *CommandSet) Execute(actor *Actor, name string, value string) {
	_, ok := cs.commands_and_alias[name]
	if ok {
		cs.commands_and_alias[name].OnExecute(actor, value)
	}
}

func (cs *CommandSet) RegisterCommand(cmd *Command) error {
	c, ok := cs.commands_and_alias[cmd.Name]
	if ok {
		return fmt.Errorf("a command named %s is already registed to %s on command set %s", c.Name, cmd.Name, cs.Name)
	}
	for _, v := range cmd.Alias {
		c, ok = cs.commands_and_alias[v]
		if ok {
			return fmt.Errorf("command %s is already registed to alias %s requested by command %s on command set %s", c.Name, v, cmd.Name, cs.Name)
		}
	}
	cs.commands_and_alias[cmd.Name] = cmd
	cs.commands_only[cmd.Name] = cmd
	for _, v := range cmd.Alias {
		cs.commands_and_alias[v] = cmd
	}
	return nil
}
