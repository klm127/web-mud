package being

import (
	"context"
	"time"

	"github.com/pwsdc/web-mud/db"
	"github.com/pwsdc/web-mud/db/dbg"
	"github.com/pwsdc/web-mud/interfaces/iserver/iactor"
	"github.com/pwsdc/web-mud/interfaces/iworld"
	"github.com/pwsdc/web-mud/server/user/actor/message"
	"github.com/pwsdc/web-mud/shared/css"
	"github.com/pwsdc/web-mud/world/being/commands"
)

type beingHuman struct {
	beingBase
	actor iactor.IActor
}

func _initBeingHuman(actor iactor.IActor, data *dbg.MudBeing) *beingHuman {
	base := _initBeingBase(data)
	hum := beingHuman{
		*base,
		actor,
	}
	return &hum
}

func NewHumanBeing(id int64, actor iactor.IActor) (iworld.IBeing, error) {
	b_data, err := db.Store.Query.GetBeingById(context.Background(), id)
	if err != nil {
		return nil, err
	}
	b := _initBeingHuman(actor, &b_data)
	actor.SetCommandGroup(commands.SenseCommands)
	return b, nil
}

// IDatabase[*dbg.MudBeing] partial - GetTimeSinceLastInteraction

func (bh *beingHuman) GetTimeSinceLastInteraction() time.Duration {
	return 0
}

// InRoom partial - setRoom

func (b *beingHuman) SetRoom(room iworld.IRoom) {
	if room == b.room {
		return
	}
	b.room = room
	room.AddBeing(b)
}

// IBeing

func (bh *beingHuman) SeeRoom() {
	if bh.room == nil {
		bh.actor.ErrorMessage("You can't figure out where you are.")
		return
	}
	bh.room.RefreshTime()
	msg := message.New().Text(bh.room.Name()).Class(css.RoomName).NewLine(1).Next()
	msg.Text(bh.room.Desc()).Class(css.RoomDesc)
	bh.actor.Message(msg.Bytes())
}

func (bh *beingHuman) Removing() {
	bh.actor.RemoveCommandGroup(commands.SenseCommands)
	bh.Save()
	// Save inventory, etc?
}
