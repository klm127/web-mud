package iactor

import (
	"time"

	"github.com/pwsdc/web-mud/db/dbg"
	"github.com/pwsdc/web-mud/interfaces/iserver"
	"github.com/pwsdc/web-mud/interfaces/iworld"
)

// An actor performs actions on the server and in the game.
type IActor interface {

	// Gets the underlying websocket connection.
	GetConnection() iserver.IConnection

	// Get the time the connection was open.
	GetTimeOpened() time.Time

	// Get the last time this actor tried to do anything.
	GetTimeLastTalked() time.Time

	// Get the time since actor was last used.
	GetTimeSinceLastTalked() time.Duration

	// Refreshes time last talked.
	RefreshTime()

	// Disconnects this actor
	Disconnect()

	// Function to call when the actor disconnects
	OnDisconnect(cb func(IActor))

	// Called whenever the actor gets some input. Usually called internally.
	ParseInput(input string)

	// Gets the command sets describing the available commands for this actor.
	GetCommandGroups() map[string]ICommandGroup

	// Gets a command group from the map, if possible.
	GetCommandGroup(key string) (ICommandGroup, bool)

	// Adds a command set, making it available to the actor.
	SetCommandGroup(cs ICommandGroup)

	// Removes a command set
	RemoveCommandGroup(cs ICommandGroup)

	// Get the current questioning status, if any.
	GetQuestioning() IQuestionResult

	// Ends an interrogation process, if ongoing.
	EndQuestioning()

	// Starts an interrogation process.
	StartQuestioning(inter IInterrogatory)

	// Associates this connection with a user.
	SetUser(user *dbg.MudUser)

	// Get the underyling user, if logged in.
	GetUser() *dbg.MudUser

	// Removes a user from this actor.
	RemoveUser()

	// Gets the being associated with this actor.
	Being() iworld.IBeing

	// Sends text to the connection.
	MessageSimple(txt string)

	// Sends formatted text to the connection.
	MessageSimplef(format string, a ...any)

	// Sends raw message bytes to the connection.
	Message(m []byte)

	// Sends an error message to the connection.
	ErrorMessage(txt string)

	// Sends a formatted error message.
	Errorf(format string, a ...any)
}
