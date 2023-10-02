package world

import (
	"github.com/pwsdc/web-mud/interfaces/iworld"
	"github.com/pwsdc/web-mud/shared"
	"github.com/pwsdc/web-mud/world/room"
)

type rooms struct {
	shared.HasResults
	rooms map[int64]iworld.IRoom
}

var Rooms rooms

func init() {
	Rooms = rooms{
		rooms: make(map[int64]iworld.IRoom),
	}
	shared.HasResultsInit(&Rooms)
}

func (r *rooms) Get(id int64) iworld.IRoom {
	room_loaded, ok := r.rooms[id]
	if !ok {
		//fmt.Println("room", id, "doesnt exist, initting")
		new_room, err := room.NewRoom(id)
		if err != nil {
			r.Errorf("error id: %d %s", id, err.Error())
			room_loaded = nil
			return nil
		}
		r.rooms[id] = new_room
		// fmt.Println("set map at id", id)
		r.Logf("Room with id %d and name '%s' initialized.", new_room.GetId(), new_room.Name())
		return new_room
	}
	// fmt.Println("said ok?", ok)
	// fmt.Println("returning room", room_loaded.Name())
	return room_loaded
}

func (r *rooms) SaveDirty() {
	for _, v := range r.rooms {
		i := 0
		if v.IsDirty() {
			v.Save()
			i++
		}
		if i > 0 {
			r.Logf("Saved %d rooms.", i)
		}
	}
}
