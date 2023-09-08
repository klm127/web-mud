package base

import (
	"github.com/gorilla/websocket"
	"github.com/pwsdc/web-mud/server/actor/message"
)

/*
An actor is a connection
*/
type Actor struct {
	id   int64
	conn *websocket.Conn
}

func (actor *Actor) Disconnect() {
	actor.conn.WriteMessage(1, message.Message("Goodbye").String())
	actor.conn.Close()
	removeActor(actor.id)
}

func CreateActor(conn *websocket.Conn) {
	id := nextId()
	actor := Actor{id, conn}
	Actors[id] = &actor
}

func (actor *Actor) MessageSimple(txt string) {
	actor.conn.WriteMessage(1, message.Message(txt).String())
}

func (actor *Actor) Message(msg *message.ActorMessage) {
	actor.conn.WriteMessage(1, msg.String())
}
