package basecommands

import (
	"time"

	"github.com/pwsdc/web-mud/interfaces/iserver/iactor"
	"github.com/pwsdc/web-mud/server/user/actor/command"
	"github.com/pwsdc/web-mud/server/user/actor/message"
)

var ConnectionCommands iactor.ICommandGroup

func init() {
	ConnectionCommands = command.NewCommandGroup("connection")

	time_cmd := command.NewCommand().Name("ctime").Desc("displays when you connected").OnExec(gettime).Get()
	disconnect_cmd := command.NewCommand().Name("disconnect").Desc("disconnects you from the server").OnExec(disconnect).Get()
	ConnectionCommands.RegisterCommands(time_cmd, disconnect_cmd)
}

func disconnect(actor iactor.IActor, msg string) {
	actor.Disconnect()
}

func gettime(actor iactor.IActor, msg string) {
	opened := actor.GetTimeOpened()
	since := time.Since(opened)
	msgb := message.New().Textf("You connected on %s.", opened.Format("2006-01-02 15:04:05")).NewLine(1).Next()
	msgb.Textf("The current time is %s.", time.Now().Format("2006-01-02 15:04:05")).NewLine(1).Next()
	msgb.Textf("You have been connected for %.2f minutes.", since.Minutes())
	actor.Message(msgb.Bytes())
}
