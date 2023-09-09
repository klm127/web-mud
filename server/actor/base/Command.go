package base

type Command struct {
	Name        string
	Alias       []string
	Description string
	OnExecute   func(*Actor, string)
}

func NewCommand(name string, description string, alias []string, onExecute func(*Actor, string)) *Command {
	cmd := Command{name, alias, description, onExecute}
	return &cmd
}
