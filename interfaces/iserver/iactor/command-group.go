package iactor

type ICommandGroup interface {
	// Gets the name of this command set.
	GetName() string

	// Gets the underlying command map
	GetCommands() *map[string]ICommand

	// Gets the names of commands registered to this command map. Does not include aliases.
	GetKeys() *[]string

	// Returns whether a given string is a command or alias of a command in this command map.
	HasCommandOrAlias(name string) bool

	// Given a key, returns the command. Otherwise, null.
	GetCommand(name string) ICommand

	// Executes a given command. Value is the leftover string after the command part has been extracted.
	Execute(actor IActor, name string, value string)

	// Registers a new command to this command map.
	RegisterCommand(cmd ICommand) error

	// Registers multiple commands to this command map.
	RegisterCommands(cmds ...ICommand) error
}
