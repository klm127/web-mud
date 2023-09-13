package db

import (
	"context"

	"github.com/pwsdc/web-mud/arg"
	"github.com/pwsdc/web-mud/db/dbg"
)

// Checks that the configured start room exists. Creates one if it doesn't.
func (store *tStore) EnsureStartRoom() {
	sroom_id := arg.Config.World.StartRoom()
	_, err := store.Query.GetRoom(context.Background(), sroom_id)
	if err != nil {
		store.Errorf("Couldn't find start room with id %d in database. Creating.", sroom_id)
	} else {
		return
	}
	rm, err := store.Query.CreateUnlinkedRoom(context.Background(), &dbg.CreateUnlinkedRoomParams{
		Name:        "A formless void",
		Description: "Chaos swirls around. Something should have been here, but isn't.",
		Objects:     make([]int64, 0),
	})
	if err != nil {
		store.Error(err.Error())
		return
	}
	arg.Config.World.OverrideStartRoom(rm.ID)

}
