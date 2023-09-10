package base

import (
	"fmt"
	"strings"
	"time"

	"github.com/gorilla/websocket"
	"github.com/pwsdc/web-mud/server/actor/message"
	"github.com/pwsdc/web-mud/util/re"
)

/*
An actor is a connection associated with commands it is able to perform.
*/
type Actor struct {
	id              int64
	conn            *websocket.Conn
	Commands        map[string]*CommandSet
	close_requested bool
	time_opened     time.Time
	time_lastTalked time.Time
	questioning     *QuestionResult
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
	actor := Actor{id, conn, map[string]*CommandSet{}, false, time.Now(), time.Now(), nil}
	fmt.Println("Creating actor")
	for s, v := range defaultCommandSets {
		fmt.Println("Adding", s, "to default actor commands.")
		actor.Commands[s] = v
	}
	Actors[id] = &actor
	go actor.listenSocket()
}

/*
Goroutine loop that listens for socket messages and dispatches them. Also updates time since last talked.
*/
func (actor *Actor) listenSocket() {
	for !actor.close_requested {
		_, data, err := actor.conn.ReadMessage()
		if err != nil {
			actor.ErrorMessage(err.Error())
			break
		}
		actor.time_lastTalked = time.Now()
		actor.ParseInput(string(data))
	}
}

/*
Determines the appropriate command the user wanted from their input string. Sends an error if no command can be found.
*/
func (actor *Actor) ParseInput(s string) {
	actor.RefreshTime()
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		actor.ErrorMessage("No command entered.")
	}
	// input is redirected to the questioning system, if active
	if actor.questioning != nil {
		actor.questioning.InputReceived(actor, s)
		if actor.questioning != nil {
			actor.questioning.AskNext(actor)
		}
		return
	}

	// test input as a command
	items := strings.SplitN(s, " ", 2)
	if re.HasPeriod.Match([]byte(items[0])) {
		if len(items) > 1 {
			actor.commandWithDomain(items[0], items[1])
		} else {
			actor.commandWithDomain(items[0], "")
		}
	} else {
		if len(items) > 1 {
			actor.anyMatchingCommand(items[0], items[1])
		} else {
			actor.anyMatchingCommand(items[0], "")
		}
	}
}

func (actor *Actor) commandWithDomain(perstring string, rest string) {
	parts := strings.SplitN(perstring, ".", 2)
	cset := parts[0]
	cmd := parts[1]
	for _, v := range actor.Commands {
		if v.Name == cset {
			if v.HasCommandOrAlias(cmd) {
				v.Execute(actor, cmd, rest)
			} else {
				actor.ErrorMessage(fmt.Sprintf("I couldn't find a command named '%s' in '%s'. Sorry.", cmd, cset))
			}
			return
		}
	}
	actor.ErrorMessage(fmt.Sprintf("I couldn't find find a command named '%s' in '%s'. Sorry.", cmd, cset))
}

func (actor *Actor) anyMatchingCommand(comname string, rest string) {
	matches := make([]*CommandSet, 0, 1)
	for _, v := range actor.Commands {
		if v.HasCommandOrAlias(comname) {
			matches = append(matches, v)
		}
	}
	if len(matches) == 1 {
		matches[0].Execute(actor, comname, rest)
	} else if len(matches) < 1 {
		actor.ErrorMessage(fmt.Sprintf("I couldn't find a command named '%s'. Try 'help' to see commands.", comname))
	} else {
		cset_names := make([]string, len(matches))
		for i, v := range matches {
			cset_names[i] = v.Name
		}
		joined := strings.Join(cset_names, ", ")

		actor.ErrorMessage(fmt.Sprintf("There are multiple commands named %s. You'll need to qualify it with one of the following: %s. For example, try %s.%s.", comname, joined, cset_names[0], comname))
	}

}

func (actor *Actor) MessageSimple(txt string) {
	actor.conn.WriteMessage(1, message.New().Text(txt).Bytes())
}

func (actor *Actor) MessageSimplef(format string, a ...any) {
	s := fmt.Sprintf(format, a...)
	actor.conn.WriteMessage(1, message.New().Text(s).Bytes())
}

func (actor *Actor) Message(m []byte) {
	actor.conn.WriteMessage(1, m)
}

func (actor *Actor) ErrorMessage(txt string) {
	actor.conn.WriteMessage(1, message.New().Color("red").Text(txt).Bytes())
}

func (actor *Actor) Errorf(format string, a ...any) {
	s := fmt.Sprintf(format, a...)
	actor.conn.WriteMessage(1, message.New().Color("red").Text(s).Bytes())
}

func (actor *Actor) GetTimeLastTalked() time.Time {
	return actor.time_lastTalked
}

func (actor *Actor) RefreshTime() {
	actor.time_lastTalked = time.Now()
}

func (actor *Actor) GetTimeOpened() time.Time {
	return actor.time_opened
}

func (actor *Actor) GetTimeSinceLastTalked() time.Duration {
	return time.Now().Sub(actor.time_lastTalked)
}

func (actor *Actor) EndQuestioning() {
	actor.questioning = nil
}

func (actor *Actor) StartQuestioning(inter *Interrogator) {
	inter.StartInterragator(actor)
}
