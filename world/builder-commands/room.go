package buildercommands

import (
	"github.com/pwsdc/web-mud/db"
	"github.com/pwsdc/web-mud/interfaces/iserver/iactor"
	"github.com/pwsdc/web-mud/server/user/actor/command"
	"github.com/pwsdc/web-mud/util/cmdvalidate"
	"github.com/pwsdc/web-mud/util/language"
	"github.com/pwsdc/web-mud/world"
)

func init() {
	exit := command.NewCommand().Name("exit").Desc("Creates a new exit for this room.").OnExec(cmdvalidate.BeingTrimLower(createExit)).Get()
	BuilderCommands.RegisterCommand(exit)
}

func createExit(actor iactor.IActor, msg string) {
	if len(msg) == 0 {
		actor.ErrorMessage("You must specify a direction, like n, s, e, w, ne, nw, se, sw, in, out, up, or down.")
		return
	}
	direction, ok := language.ParseDirection(msg)
	if !ok {
		actor.Errorf("%s is not a direction I know.", msg)
		return
	}
	room_data, err := db.Store.CreateEmptyRoom("a formless void", "It swirls with purpose.")
	if err != nil {
		actor.Errorf("The room failed to manifest.")
	}
	new_room := world.Rooms.Get(room_data.ID)
	actors_room := actor.Being().GetRoom()
	dir_opposite := language.ParseOppositeDirection(direction)

	er_strn := actors_room.GetBuilder().Link(new_room, direction)
	if len(er_strn) > 0 {
		actor.Errorf(er_strn)
		return
	}
	er_strn = new_room.GetBuilder().Link(actors_room, dir_opposite)
	if len(er_strn) > 0 {
		actor.Errorf(er_strn)
		return
	}

	actor.MessageSimplef("A void has manifested to the %s.", language.ParseDirectionFull(direction))
}

/*
	build.set room name <name>
	build.set room desc <desc>
	build.exit <direction>

*/
