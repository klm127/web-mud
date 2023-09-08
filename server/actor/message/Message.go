package message

import "encoding/json"

/*
	A message which will be sent to the actors
*/
type ActorMessage struct {
	text string
}

func Message(msg string) *ActorMessage {
	m := ActorMessage{msg}
	return &m
}

func (self *ActorMessage) String() []byte {
	bytes, err := json.Marshal(self)
	if err != nil {
		bytes = []byte("Error processing server message")
	}
	return bytes
}
