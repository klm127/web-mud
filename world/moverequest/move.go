package moverequest

import (
	"fmt"
	"time"

	"github.com/pwsdc/web-mud/interfaces/iworld"
	"github.com/pwsdc/web-mud/util/language"
	"github.com/pwsdc/web-mud/world/sight"
)

type moveReq struct {
	source       iworld.IBeing
	req_time     time.Time
	time_to_take time.Duration
	// Must be a valid dir constant
	direction         string
	repeats_remaining int8
}

func NewMoveRequest(from iworld.IBeing, time_for_move time.Duration, direction string, num_repeats int8) iworld.IMoveReq {
	return &moveReq{
		source:            from,
		req_time:          time.Now(),
		time_to_take:      time_for_move,
		direction:         direction,
		repeats_remaining: num_repeats,
	}
}

func (mr *moveReq) CanMove() bool {
	return time.Since(mr.req_time) > mr.time_to_take
}

func (mr *moveReq) GetRepeats() int8 {
	return mr.repeats_remaining
}

func (mr *moveReq) GetDirection() string {
	return mr.direction
}

func (mr *moveReq) GetMover() iworld.IBeing {
	return mr.source
}

func (mr *moveReq) Move(to iworld.IRoom) {
	room := mr.source.GetRoom()
	if to == nil {
		msg_self := fmt.Sprintf("You smack into the matter %s and fail to move it.", mr.direction)
		msg_others := fmt.Sprintf("smacks into the matter %s trying to move through it.", mr.direction)
		sight := sight.NewSeen(mr.source, &msg_self, &msg_others)
		room.SightEmit(sight)
		mr.repeats_remaining = 0
		return
	}
	mr.req_time = time.Now()
	mr.repeats_remaining -= 1
	msg_others_new_room := fmt.Sprintf("enters from the %s.", language.ParseOppositeDirection(mr.direction))
	sight_enter := sight.NewSeen(mr.source, nil, &msg_others_new_room)
	to.SightEmit(sight_enter)
	fulldir := language.ParseDirectionFull(mr.direction)
	msg_self := fmt.Sprintf("You move %s.", fulldir)
	msg_others := fmt.Sprintf("heads %s.", fulldir)
	sight_leave := sight.NewSeen(mr.source, &msg_self, &msg_others)
	room.SightEmit(sight_leave)
	to.AddBeing(mr.source)
}
