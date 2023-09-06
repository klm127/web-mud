package msg

type Server struct {
	error bool
	msg   string
}

func ServerError(msg string) Server {
	return Server{true, msg}
}
func ServerMsg(msg string) Server {
	return Server{false, msg}
}
