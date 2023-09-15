package iworld

import "github.com/pwsdc/web-mud/db/dbg"

type IRoom interface {
	IExists
	IDatabase[*dbg.MudRoom]
	// Get everything that exists here
	GetHere() []IExists
	// Get beings here
	GetBeingsHere() []IBeing
	// Add a being to this room
	AddBeing(IBeing)
	// Remove a being from this room
	RemoveBeing(IBeing)
	// Called when a sound is emitted
	SoundEmit(ISound)
	// Called when a sight is emitted
	SightEmit(ISeen)
	// Get a room builder for this room
	GetBuilder() IRoomBuilder
	// Get a list of available directions
	GetDirectionList() string
}

// For anything that's in a room
type IInRoom interface {
	GetRoom() IRoom
	SetRoom(IRoom)
}

type IRoomBuilder interface {
	Name(string) IRoomBuilder
	Desc(string) IRoomBuilder
	// Returns error string for user messaging; 0 length if success
	Link(IRoom, string) string
}
