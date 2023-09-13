package iactor

type ICommand interface {
	// Gets the name of this command, used to execute it.
	GetName() string
	// Gets aliases for this command, used to execute it.
	GetAliases() *[]string
	// Gets the description of what executing this command does.
	GetDescription() string
	// Sets the function to be called when this command is executed.
	OnExecute(cb func(actor IActor, msg string))
	// Execute this command. Msg is any extra text with the command-executing sentence.
	Execute(actor IActor, msg string)
}

type CommandFunc func(IActor, string)
