package db

import (
	"context"

	"github.com/pwsdc/web-mud/db/dbg"
)

func (store *tStore) CreateEmptyRoom(name string, description string) (*dbg.MudRoom, error) {
	room, err := store.Query.CreateUnlinkedRoom(context.Background(), &dbg.CreateUnlinkedRoomParams{
		Name:        name,
		Description: description,
		Objects:     make([]int64, 0),
	})
	if err != nil {
		store.Error("Failed to create room " + name + " error: " + err.Error())
		return nil, err
	}
	return &room, err
}
