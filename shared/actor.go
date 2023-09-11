package shared

type ISendsMessages interface {
	Message(m []byte)
	MessageSimplef(format string, a ...any)
	MessageSimple(txt string)
	ErrorMessage(txt string)
	Errorf(format string, a ...any)
}

type IExecutesCommands interface {
	StartQuestioning(IInterogator)
}

type IActor interface {
	ISendsMessages
	IExecutesCommands
}
