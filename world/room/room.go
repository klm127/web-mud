package room

import (
	"context"
	"strings"

	"github.com/pwsdc/web-mud/db"
	"github.com/pwsdc/web-mud/db/dbg"
	"github.com/pwsdc/web-mud/interfaces/iworld"
	"github.com/pwsdc/web-mud/shared/enum"
	"github.com/pwsdc/web-mud/util/language"
)

type troom struct {
	room_base
	beingsHere map[int64]iworld.IBeing
	builder    iworld.IRoomBuilder
}

func _initRoom(data *dbg.MudRoom) *troom {
	rb := _initRoomBase(data)
	rn := troom{
		*rb,
		make(map[int64]iworld.IBeing),
		nil,
	}
	rn.builder = newRoomBuilder(&rn)
	return &rn
}

func NewRoom(id int64) (iworld.IRoom, error) {
	room, err := db.Store.Query.GetRoom(context.Background(), id)
	if err != nil {
		return nil, err
	}
	room_instance := _initRoom(&room)
	return room_instance, nil

}

// IRoom

func (r *troom) GetHere() []iworld.IExists {
	exists := make([]iworld.IExists, 0, len(r.beingsHere))
	for _, v := range r.beingsHere {
		exists = append(exists, v.(iworld.IExists))
	}
	return exists
}

func (r *troom) GetBeingsHere() []iworld.IBeing {
	here := make([]iworld.IBeing, 0, len(r.beingsHere))
	for _, v := range r.beingsHere {
		here = append(here, v)
	}
	return here
}

func (r *troom) AddBeing(new_being iworld.IBeing) {
	old_room := new_being.GetRoom()
	if old_room == r {
		return
	}
	if old_room != nil {
		old_room.RemoveBeing(new_being)
	}
	r.beingsHere[new_being.GetId()] = new_being
	new_being.SetRoom(r)
}

func (r *troom) RemoveBeing(to_remove iworld.IBeing) {
	delete(r.beingsHere, to_remove.GetId())
}

func (r *troom) SoundEmit(sound iworld.ISound) {
	for _, v := range r.beingsHere {
		v.SoundHear(sound)
	}
}

func (r *troom) GetDirectionList() string {
	possible := make([]string, 0, 10)
	if r.data.N.Valid {
		possible = append(possible, language.ParseDirectionFull(enum.North))
	}
	if r.data.S.Valid {
		possible = append(possible, language.ParseDirectionFull(enum.South))
	}
	if r.data.E.Valid {
		possible = append(possible, language.ParseDirectionFull(enum.East))
	}
	if r.data.W.Valid {
		possible = append(possible, language.ParseDirectionFull(enum.West))
	}
	if r.data.Ne.Valid {
		possible = append(possible, language.ParseDirectionFull(enum.NorthEast))
	}
	if r.data.Se.Valid {
		possible = append(possible, language.ParseDirectionFull(enum.SouthEast))
	}
	if r.data.Nw.Valid {
		possible = append(possible, language.ParseDirectionFull(enum.NorthWest))
	}
	if r.data.Sw.Valid {
		possible = append(possible, language.ParseDirectionFull(enum.SouthWest))
	}
	if r.data.Se.Valid {
		possible = append(possible, language.ParseDirectionFull(enum.In))
	}
	if r.data.U.Valid {
		possible = append(possible, language.ParseDirectionFull(enum.Up))
	}
	if r.data.D.Valid {
		possible = append(possible, language.ParseDirectionFull(enum.Down))
	}
	return strings.Join(possible, ", ")
}

func (r *troom) GetBuilder() iworld.IRoomBuilder {
	return r.builder
}
