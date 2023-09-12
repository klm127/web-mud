package room

import (
	"context"

	"github.com/pwsdc/web-mud/db"
	"github.com/pwsdc/web-mud/db/dbg"
	"github.com/pwsdc/web-mud/interfaces/iworld"
)

type room struct {
	room_base
	beingsHere map[int64]iworld.IBeing
}

func _initRoom(data *dbg.MudRoom) *room {
	rb := _initRoomBase(data)
	rn := room{
		*rb,
		make(map[int64]iworld.IBeing),
	}
	return &rn
}

func NewRoom(id int64) (iworld.IRoom, error) {
	room, err := db.Store.Query.GetRoom(context.Background(), id)
	if err != nil {
		return nil, err
	}
	room_instance := _initRoom(&room)
	return room_instance, nil

}

// IRoom

func (r *room) GetHere() []iworld.IExists {
	exists := make([]iworld.IExists, 0, len(r.beingsHere))
	for _, v := range r.beingsHere {
		exists = append(exists, v.(iworld.IExists))
	}
	return exists
}

func (r *room) GetBeingsHere() []iworld.IBeing {
	here := make([]iworld.IBeing, 0, len(r.beingsHere))
	for _, v := range r.beingsHere {
		here = append(here, v)
	}
	return here
}

func (r *room) AddBeing(new_being iworld.IBeing) {
	old_room := new_being.GetRoom()
	if old_room == r {
		return
	}
	if old_room != nil {
		old_room.RemoveBeing(new_being)
	}
	r.beingsHere[new_being.GetId()] = new_being
	new_being.SetRoom(r)
}

func (r *room) RemoveBeing(to_remove iworld.IBeing) {
	delete(r.beingsHere, to_remove.GetId())
}
