package shared

type ICommand interface {
	GetName() string
	GetAliases() []string
	GetDescription() string
	Execute(IActor, string)
}

type ICommandSet interface {
	GetName() string
	GetCommands() map[string]ICommand
	GetKeys() []string
	HasCommandOrAlias(name string) bool
	Getcommand(name string) ICommand
	Execute(IActor, name string, value string)
	RegisterCommand(ICommand) error
}
