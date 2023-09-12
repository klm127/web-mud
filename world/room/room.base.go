package room

import (
	"time"

	"github.com/pwsdc/web-mud/db/dbg"
	"github.com/pwsdc/web-mud/shared"
)

type room_base struct {
	shared.HasResults
	data             *dbg.MudRoom
	dirty            bool
	loaded_time      time.Time
	last_interaction time.Time
}

func _initRoomBase(data_room *dbg.MudRoom) *room_base {
	rb := room_base{
		data:             data_room,
		dirty:            false,
		loaded_time:      time.Now(),
		last_interaction: time.Now(),
	}
	shared.HasResultsInit(&rb)
	return &rb
}

// IDatabase[*dbg.MudRoom]

func (r *room_base) GetData() *dbg.MudRoom {
	return r.data
}
func (r *room_base) Save() {
	// todo: update query
}
func (r *room_base) IsDirty() bool {
	return r.dirty
}
func (r *room_base) GetId() int64 {
	return r.data.ID
}
func (r *room_base) GetInstancedTime() time.Time {
	return r.loaded_time
}
func (r *room_base) GetLastInteractionTime() time.Time {
	return r.last_interaction
}
func (r *room_base) GetTimeSinceLastInteraction() time.Duration {
	return time.Since(r.last_interaction)
}
func (r *room_base) RefreshTime() {
	r.last_interaction = time.Now()
}

// IExists

func (r *room_base) Name() string {
	return r.data.Name
}
func (r *room_base) Desc() string {
	return r.data.Description
}
func (r *room_base) Img() *string {
	if r.data.Img.Valid {
		return &r.data.Img.String
	} else {
		return nil
	}
}
