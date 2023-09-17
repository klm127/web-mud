package iserver

/** Could be a regular gorilla websocket connection or a psuedo-socket HTTP connection */
type IConnection interface {
	ReadMessage() (int, []byte, error)
	WriteMessage(int, []byte) error
	Close() error
}
