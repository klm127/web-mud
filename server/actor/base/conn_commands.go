package base

/*
Commands relating to connecting / disconnecting that are registered default set.
*/
func init() {
	cs := NewCommandSet("connection")
	qcom := NewCommand("disconnect", "close your connection with the server", []string{"quit", "exit"}, disconnect)
	cs.RegisterCommand(qcom)
	RegisterDefaultCommandSet(cs)
}

// Disconnect socket
func disconnect(actor *Actor, extratxt string) {
	actor.Disconnect()
}
