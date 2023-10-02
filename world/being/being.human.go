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
	"github.com/pwsdc/web-mud/world/sight"
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

var appear_msg_self = "You solidify."
var appear_msg_others = "solidifies."

func (b *beingHuman) SetRoom(room iworld.IRoom) {
	first_room := false
	if room == b.room {
		return
	}
	if b.room == nil {
		first_room = true
	}
	b.room = room
	room.AddBeing(b)
	if first_room {
		emit := sight.NewSeen(b, &appear_msg_self, &appear_msg_others)
		room.SightEmit(emit)
	}
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
	dirs := bh.room.GetDirectionList()
	if len(dirs) == 0 {
		bh.actor.MessageSimplef("There seem to be no exits.")
	} else {
		msgdir := message.New().Class(css.RoomDirections).Text("You can go: ").Next().Text(bh.room.GetDirectionList())
		bh.actor.Message(msgdir.Bytes())
	}
}

var remove_sight_msg_self = "You evaporate."
var remove_sight_msg_other = "evaporates."

func (bh *beingHuman) Removing() {
	bh.actor.RemoveCommandGroup(commands.SenseCommands)
	bh.actor.RemoveCommandGroup(commands.VoiceCommands)
	vis := sight.NewSeen(bh, &remove_sight_msg_self, &remove_sight_msg_other)
	bh.room.SightEmit(vis)
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

func (bh *beingHuman) SightSee(sight iworld.ISeen) {
	msg := sight.GetMessage(bh)
	bh.actor.MessageSimplef(msg)
}
