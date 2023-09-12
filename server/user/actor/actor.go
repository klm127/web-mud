package actor

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/pwsdc/web-mud/db/dbg"
	"github.com/pwsdc/web-mud/interfaces/iserver/iactor"
	"github.com/pwsdc/web-mud/interfaces/iworld"
	"github.com/pwsdc/web-mud/server/user/actor/message"
	"github.com/pwsdc/web-mud/util/re"
	"github.com/pwsdc/web-mud/world"
)

// implements interface
type Actor struct {
	id              int64
	conn            *websocket.Conn
	commands        map[string]iactor.ICommandGroup
	close_requested bool
	time_opened     time.Time
	time_lastTalked time.Time
	questioning     iactor.IQuestionResult
	user            *dbg.MudUser
	being           iworld.IBeing
	onDisconnect    func(iactor.IActor)
	mutex           *sync.Mutex
}

func newActor(actor_id int64, connect *websocket.Conn) *Actor {
	actor := Actor{
		id:              actor_id,
		conn:            connect,
		commands:        make(map[string]iactor.ICommandGroup),
		close_requested: false,
		time_opened:     time.Now(),
		time_lastTalked: time.Now(),
		questioning:     nil,
		user:            nil,
		being:           nil,
		onDisconnect:    nil,
		mutex:           &sync.Mutex{},
	}
	go actor.listenToSocketMessages()
	return &actor
}

func (actor *Actor) listenToSocketMessages() {
	for !actor.close_requested {
		actor.mutex.Lock()
		_, data, err := actor.conn.ReadMessage()
		if err != nil {
			actor.ErrorMessage(err.Error())
			actor.mutex.Unlock()
			break
		}
		actor.RefreshTime()
		actor.ParseInput(string(data))
		actor.mutex.Unlock()
	}
}

func (actor *Actor) GetConnection() *websocket.Conn {
	return actor.conn
}
func (actor *Actor) GetTimeOpened() time.Time {
	return actor.time_opened
}
func (actor *Actor) GetTimeLastTalked() time.Time {
	return actor.time_lastTalked
}
func (actor *Actor) GetTimeSinceLastTalked() time.Duration {
	return time.Since(actor.time_lastTalked)
}
func (actor *Actor) RefreshTime() {
	actor.time_lastTalked = time.Now()
}
func (actor *Actor) Disconnect() {
	actor.conn.WriteMessage(1, message.New().Text("Goodbye").Bytes())
	actor.close_requested = true
	if actor.onDisconnect != nil {
		actor.onDisconnect(actor)
	}
	if actor.being != nil {
		world.Beings.Remove(actor.being)
	}
	// do anything with the user?
	actor.conn.Close()
}
func (actor *Actor) OnDisconnect(cb func(iactor.IActor)) {
	actor.onDisconnect = cb
}

func (actor *Actor) ParseInput(input string) {
	actor.RefreshTime()
	input = strings.TrimSpace(input)
	if len(input) == 0 {
		actor.ErrorMessage("Do you have something to stay?")
	}
	if actor.questioning != nil {
		actor.questioning.InputReceived(input)
		if actor.questioning != nil {
			actor.questioning.AskNext()
		}
		return
	}

	items := strings.SplitN(input, " ", 2)
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
func (actor *Actor) GetCommandGroups() map[string]iactor.ICommandGroup {
	return actor.commands
}

func (actor *Actor) GetCommandGroup(key string) (iactor.ICommandGroup, bool) {
	grp, ok := actor.commands[key]
	return grp, ok
}

func (actor *Actor) SetCommandGroup(cs iactor.ICommandGroup) {
	actor.commands[cs.GetName()] = cs
}
func (actor *Actor) RemoveCommandGroup(cs iactor.ICommandGroup) {
	delete(actor.commands, cs.GetName())
}
func (actor *Actor) GetQuestioning() iactor.IQuestionResult {
	return actor.questioning
}
func (actor *Actor) EndQuestioning() {
	actor.questioning = nil
}
func (actor *Actor) StartQuestioning(inter iactor.IInterrogatory) {
	actor.questioning = inter.GetQuestioner(actor)
	inter.SendIntro(actor)
	actor.questioning.AskNext()
}
func (actor *Actor) SetUser(user *dbg.MudUser) {
	if user == nil {
		actor.ErrorMessage("You are trying to act with a user that doesn't exist.")
		return
	}
	actor.user = user
	actor.being = world.Beings.GetHuman(user.Being, actor)
}
func (actor *Actor) GetUser() *dbg.MudUser {
	return actor.user
}
func (actor *Actor) RemoveUser() {
	world.Beings.Remove(actor.being)
	actor.user = nil
}
func (actor *Actor) Being() iworld.IBeing {
	if actor.being != nil {
		return actor.being
	}
	return nil
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
