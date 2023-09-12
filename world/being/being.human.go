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
	actor.SetCommandGroup(commands.VoiceCommands)
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
	b.SeeRoom()
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
	bh.actor.RemoveCommandGroup(commands.VoiceCommands)
	bh.Save()
	// Save inventory, etc?
}

func (bh *beingHuman) SoundHear(sound iworld.ISound) {
	is_me := true
	name := "You"
	verb := "say"
	if sound.GetSource() != bh {
		is_me = false
		name = sound.GetSourceName()
	}
	loudness := sound.GetLoudness()
	if loudness < 5 {
		verb = "whisper"
	} else if loudness < 15 {
		verb = "say"
	} else {
		verb = "yell"
	}
	if !is_me {
		verb += "s"
	}
	msg := sound.GetMessage()
	last_char := msg[len(msg)-1]
	if last_char != '.' && last_char != '!' && last_char != '?' {
		msg += "."
	}
	bh.actor.MessageSimplef("%s %s, \"%s\"", name, verb, msg)
}
