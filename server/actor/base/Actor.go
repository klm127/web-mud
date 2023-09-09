package base

import (
	"fmt"
	"strings"

	"github.com/gorilla/websocket"
	"github.com/pwsdc/web-mud/server/actor/message"
)

/*
An actor is a connection associated with commands it is able to perform.
*/
type Actor struct {
	id              int64
	conn            *websocket.Conn
	Commands        map[string]*CommandSet
	close_requested bool
}

func (actor *Actor) Disconnect() {
	actor.conn.WriteMessage(1, message.New().Text("Goodbye").Bytes())
	actor.close_requested = true
	removeActor(actor.id)
	actor.conn.Close()
}

/*
Creates a new actor associated with a socket connection, then adds that actor to the active actors map.

Also loads the actor with the baseCommandSets - the default commands available to all actors at all times.

Finally it starts the actor.listenSocket() go routine, where input will be processed.
*/
func CreateActor(conn *websocket.Conn) {
	id := nextId()
	actor := Actor{id, conn, map[string]*CommandSet{}, false}
	fmt.Println("Creating actor")
	for s, v := range defaultCommandSets {
		fmt.Println("Adding", s, "to default actor commands.")
		actor.Commands[s] = v
	}
	Actors[id] = &actor
	go actor.listenSocket()
}

func (actor *Actor) listenSocket() {
	for !actor.close_requested {
		_, data, err := actor.conn.ReadMessage()
		if err != nil {
			actor.ErrorMessage(err.Error())
			break
		}
		actor.ParseInput(string(data))
	}
}

func (actor *Actor) ParseInput(s string) {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		actor.ErrorMessage("No command entered.")
	}
	items := strings.SplitN(s, " ", 2)
	matches := make([]*CommandSet, 0, 1)
	for _, v := range actor.Commands {
		if v.HasCommandOrAlias(items[0]) {
			matches = append(matches, v)
		}
	}
	if len(matches) < 1 {
		actor.ErrorMessage(fmt.Sprintf("Command %s not understood.", items[0]))
		return
	}
	if len(matches) == 1 {
		if len(items) > 1 {
			matches[0].Execute(actor, items[0], items[1])
		} else {
			matches[0].Execute(actor, items[0], "")
		}
		return
	} else {
		actor.ErrorMessage("Multiple commands matched that input!")
	}
}

func (actor *Actor) MessageSimple(txt string) {
	actor.conn.WriteMessage(1, message.New().Text(txt).Bytes())
}

func (actor *Actor) Message(m []byte) {
	actor.conn.WriteMessage(1, m)
}

func (actor *Actor) ErrorMessage(txt string) {
	actor.conn.WriteMessage(1, message.New().Color("red").Text(txt).Bytes())
}
