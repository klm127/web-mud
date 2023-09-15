package iworld

import (
	"github.com/pwsdc/web-mud/db/dbg"
)

type IBeing interface {
	IExists
	IInRoom
	IDatabase[*dbg.MudBeing]
	// Being sees the room they are in.
	SeeRoom()
	// Called before being is deleted from a map so it can do whatever it needs to do
	Removing()
	// Called by a room when a sound is emitted
	SoundHear(ISound)
	// Called by a room when a sight is emitted
	SightSee(ISeen)
}
