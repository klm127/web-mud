package commands

import (
	"time"

	"github.com/pwsdc/web-mud/interfaces/iserver/iactor"
	"github.com/pwsdc/web-mud/interfaces/iworld"
	"github.com/pwsdc/web-mud/server/user/actor/command"
	"github.com/pwsdc/web-mud/util/language"
	"github.com/pwsdc/web-mud/world/moverequest"
)

type tmoveReceiver interface {
	New(iworld.IMoveReq)
}

var MoveCommands iactor.ICommandGroup
var moveReceiver tmoveReceiver

func init() {
	MoveCommands = command.NewCommandGroup("move")
	MoveCommands.RegisterCommand(command.NewCommand().Name("walk").Alias("w").Desc("moves your feet at a normal pace in a given direction").OnExec(being(walk)).Get())
}

func RegisterMoveManager(mr tmoveReceiver) {
	moveReceiver = mr
}

func walk(actor iactor.IActor, msg string) {
	if len(msg) == 0 {
		actor.MessageSimplef("In which direction?")
		return
	}
	dir, ok := language.ParseDirection(msg)
	if !ok {
		actor.MessageSimplef("That's not a direction I've heard of.")
		return
	}

	move_time := time.Duration(1) * time.Second
	move_request := moverequest.NewMoveRequest(actor.Being(), move_time, dir, 1)
	moveReceiver.New(move_request)
	actor.MessageSimplef("You begin moving %s.", dir)

}
