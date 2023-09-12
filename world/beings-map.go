package world

import (
	"github.com/pwsdc/web-mud/interfaces/iserver/iactor"
	"github.com/pwsdc/web-mud/interfaces/iworld"
	"github.com/pwsdc/web-mud/shared"
	"github.com/pwsdc/web-mud/world/being"
)

type beings struct {
	shared.HasResults
	beings map[int64]iworld.IBeing
}

var Beings beings

func init() {
	Beings = beings{
		beings: make(map[int64]iworld.IBeing),
	}
	shared.HasResultsInit(&Beings)
}

func (b *beings) GetHuman(id int64, actor iactor.IActor) iworld.IBeing {
	if actor == nil {
		b.Error("Get human called with nil parameter.")
	}
	being_instanced, ok := b.beings[id]
	if ok {
		being_instanced.Save()
	}
	b_new, err := being.NewHumanBeing(id, actor)
	if err != nil {
		b.Error(err.Error())
		return nil
	}
	b.beings[id] = b_new
	room_id := b_new.GetData().Room
	room := Rooms.Get(room_id)
	if room == nil {
		b.Errorf("Failed to initialize room %d for being %s.", room_id, b_new.Name())
		return nil
	}
	room.AddBeing(b_new)
	return b_new
}

func (b *beings) Remove(being iworld.IBeing) {
	if being == nil {
		b.Error("Remove called with nil parameter.")
		return
	}
	id := being.GetId()
	being.Removing()
	delete(b.beings, id)
}
