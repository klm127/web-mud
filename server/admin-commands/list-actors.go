package admincommands

import (
	"github.com/pwsdc/web-mud/interfaces/iserver/iactor"
	"github.com/pwsdc/web-mud/server/user/actor"
	"github.com/pwsdc/web-mud/server/user/actor/command"
	"github.com/pwsdc/web-mud/server/user/actor/message"
)

func init() {
	list_cmd := command.NewSplitCommand().Name("list").Desc("lists objects instanced on the server").OnExec(list_base)

	list_actors_cmd := command.NewCommand().Name("actors").Desc("list actors instanced on the server").OnExec(listActors)
	list_cmd.SetSplit("actors", list_actors_cmd.Get())

	AdminCommands.RegisterCommand(list_cmd.Get())
}

// actors disconnect id

// logs beings
// errors beings
// list beings

// logs db
// logs rooms
// errors rooms
// list rooms

func listActors(an_actor iactor.IActor, msg string) {
	traverser := getLister(an_actor)
	actor.Traverse(traverser, false)
}

func getLister(actor iactor.IActor) func(*map[int64]iactor.IActor) {
	return func(amap *map[int64]iactor.IActor) {
		mb := message.New().Text("Active actors logs:").NewLine(1).Next()
		for id, v := range *amap {
			mb.Textf("Actor id: %d ", id)
			if v.Being() != nil {
				mb.Next().Textf("named: %s", v.Being().Name())
			}
			mb.Next().Textf("since: %s", v.GetTimeOpened().Format("2006-01-02 15:04:05")).NewLine(1).Next()
		}
		actor.Message(mb.Bytes())
	}
}
