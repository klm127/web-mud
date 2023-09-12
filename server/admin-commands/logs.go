package admincommands

import (
	"github.com/pwsdc/web-mud/db"
	"github.com/pwsdc/web-mud/interfaces/iserver/iactor"
	"github.com/pwsdc/web-mud/server/user/actor/command"
	"github.com/pwsdc/web-mud/server/user/actor/message"
	"github.com/pwsdc/web-mud/world"
)

func init() {
	log_cmd := command.NewSplitCommand().Name("log").OnExec(log_base)

	log_db_cmd := command.NewCommand().Name("db").OnExec(logs_db).Get()
	log_beings_cmd := command.NewCommand().Name("being").OnExec(logs_beings).Get()
	log_rooms_cmd := command.NewCommand().Name("room").OnExec(logs_rooms).Get()
	log_cmd.SetSplit("db", log_db_cmd)
	log_cmd.SetSplit("beings", log_beings_cmd)
	log_cmd.SetSplit("rooms", log_rooms_cmd)

	AdminCommands.RegisterCommand(log_cmd.Get())
}

func log_base(actor iactor.IActor, msg string) {
	actor.MessageSimple("Try 'log db', 'log beings', or 'log rooms'")
}

func logs_db(actor iactor.IActor, msg string) {
	mb := message.New().Text("Database logs:")
	actor.Message(mb.Bytes())
	db.Store.MessageResults(actor)
}

func logs_beings(actor iactor.IActor, msg string) {
	mb := message.New().Text("Beings logs:")
	actor.Message(mb.Bytes())
	world.Beings.MessageResults(actor)
}

func logs_rooms(actor iactor.IActor, msg string) {
	mb := message.New().Text("Room logs:")
	actor.Message(mb.Bytes())
	world.Rooms.MessageResults(actor)
}
