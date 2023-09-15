package world

import "github.com/pwsdc/web-mud/interfaces/iworld"

type movements struct {
	actively_moving map[iworld.IBeing]iworld.IMoveReq
}

var Movements movements

func init() {
	Movements = movements{
		actively_moving: make(map[iworld.IBeing]iworld.IMoveReq),
	}
}

func (m *movements) New(move iworld.IMoveReq) {
	m.actively_moving[move.GetMover()] = move
}

func (m *movements) Process() {
	for k, v := range m.actively_moving {
		if v.GetRepeats() <= 0 {
			delete(m.actively_moving, k)
			continue
		}
		if v.CanMove() {
			room := v.GetMover().GetRoom()
			adjacent_id := room.GetAdjacentID(v.GetDirection())
			adjac := Rooms.Get(*adjacent_id)
			v.Move(adjac)
		}
	}
}
