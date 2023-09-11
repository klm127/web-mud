package world

import (
	"context"
	"time"

	"github.com/pwsdc/web-mud/db"
	"github.com/pwsdc/web-mud/db/dbg"
)

var LoadedRooms map[int64]*Room

type Room struct {
	data *dbg.MudRoom
	// Whether this object has been changed from the database.
	dirty          bool
	timeSinceVisit time.Time
}

func init() {
	LoadedRooms = make(map[int64]*Room)
}

func GetRoom(id int64) *Room {
	room, ok := LoadedRooms[id]
	if ok {
		return room
	}
	room_data, err := db.Store.Query.GetRoom(context.Background(), id)
	if err != nil {
		return nil
	}
	new_room := Room{&room_data, false, time.Now()}
	LoadedRooms[id] = &new_room
	return &new_room
}

func (r *Room) Description() *string {
	return &r.data.Description
}

func (r *Room) SetDescription(newdesc string) {
	r.data.Description = newdesc
	r.dirty = true
}

func (r *Room) Name() *string {
	return &r.data.Name
}

func (r *Room) SetName(newname string) {
	r.data.Name = newname
	r.dirty = true
}

func (r *Room) Img() *string {
	if r.data.Img.Valid {
		return &r.data.Img.String
	}
	return nil
}

func (r *Room) SetImg(newimg string) {
	r.data.Img.Valid = true
	r.data.Img.String = newimg
	r.dirty = true
}

func (r *Room) UpdateTime() {
	r.timeSinceVisit = time.Now()
}
