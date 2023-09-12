package being

import (
	"context"
	"time"

	"github.com/pwsdc/web-mud/db"
	"github.com/pwsdc/web-mud/db/dbg"
	"github.com/pwsdc/web-mud/interfaces/iworld"
	"github.com/pwsdc/web-mud/shared"
)

// Implements iworld.IExists
type beingBase struct {
	shared.HasResults
	data            *dbg.MudBeing
	time_instanced  time.Time
	time_lastTalked time.Time
	dirty           bool
	room            iworld.IRoom
}

func _initBeingBase(data *dbg.MudBeing) *beingBase {
	b := beingBase{
		data:            data,
		dirty:           false,
		time_instanced:  time.Now(),
		time_lastTalked: time.Now(),
	}
	shared.HasResultsInit(&b)
	return &b
}

// IDatabase[*dbg.MudBeing] partial - omit GetTimeSinceLastInteraction

func (b *beingBase) GetData() *dbg.MudBeing {
	return b.data
}

func (b *beingBase) Save() {
	if b.dirty {
		params := dbg.UpdateBeingParams{
			Description: b.data.Description,
			Room:        b.data.Room,
		}
		err := db.Store.Query.UpdateBeing(context.Background(), &params)
		if err != nil {
			b.dirty = false
		} else {
			b.Logf("Error updating database: %s", err.Error())
		}
	}

}

func (b *beingBase) IsDirty() bool {
	return b.dirty
}

func (b *beingBase) GetId() int64 {
	return b.data.ID
}
func (b *beingBase) GetInstancedTime() time.Time {
	return b.time_instanced
}
func (b *beingBase) GetLastInteractionTime() time.Time {
	return b.time_lastTalked
}
func (b *beingBase) RefreshTime() {
	b.time_lastTalked = time.Now()
}

// IExists

func (b *beingBase) Name() string {
	return b.data.Name
}

func (b *beingBase) Desc() string {
	return b.data.Description
}

func (b *beingBase) Img() *string {
	return nil
}

// InRoom partial; SetRoom omitted

func (b *beingBase) GetRoom() iworld.IRoom {
	return b.room
}
