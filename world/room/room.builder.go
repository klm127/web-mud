package room

import (
	"database/sql"
	"fmt"

	"github.com/pwsdc/web-mud/interfaces/iworld"
	"github.com/pwsdc/web-mud/shared/enum"
	"github.com/pwsdc/web-mud/util/language"
)

type roomBuilder struct {
	room *troom
}

func newRoomBuilder(room *troom) iworld.IRoomBuilder {
	return &roomBuilder{room: room}
}

func (rb *roomBuilder) Name(to string) iworld.IRoomBuilder {
	rb.room.dirty = true
	rb.room.data.Name = to
	return rb
}

func (rb *roomBuilder) Desc(to string) iworld.IRoomBuilder {
	rb.room.dirty = true
	rb.room.data.Description = to
	return rb
}

func (rb *roomBuilder) Link(to iworld.IRoom, dirraw string) string {
	dir, ok := language.ParseDirection(dirraw)
	if !ok {
		return fmt.Sprintf("I don't know which direction %s is in.", dirraw)
	}
	rb.room.dirty = true
	to_id := sql.NullInt64{Valid: true, Int64: to.GetId()}

	dir_full := language.ParseDirectionFull(dir)

	switch dir {
	case enum.North:
		if rb.room.data.N.Valid {
			return fmt.Sprintf("It would seem that %s already has reality manifest.", dir_full)
		}
		rb.room.data.N = to_id
	case enum.South:
		if rb.room.data.S.Valid {
			return fmt.Sprintf("It would seem that %s already has reality manifest.", dir_full)
		}
		rb.room.data.S = to_id
	case enum.East:
		if rb.room.data.E.Valid {
			return fmt.Sprintf("It would seem that %s already has reality manifest.", dir_full)
		}
		rb.room.data.E = to_id
	case enum.West:
		if rb.room.data.W.Valid {
			return fmt.Sprintf("It would seem that %s already has reality manifest.", dir_full)
		}
		rb.room.data.W = to_id
	case enum.NorthEast:
		if rb.room.data.Ne.Valid {
			return fmt.Sprintf("It would seem that %s already has reality manifest.", dir_full)
		}
		rb.room.data.Ne = to_id
	case enum.NorthWest:
		if rb.room.data.Nw.Valid {
			return fmt.Sprintf("It would seem that %s already has reality manifest.", dir_full)
		}
		rb.room.data.Nw = to_id
	case enum.SouthEast:
		if rb.room.data.Se.Valid {
			return fmt.Sprintf("It would seem that %s already has reality manifest.", dir_full)
		}
		rb.room.data.Se = to_id
	case enum.SouthWest:
		if rb.room.data.Sw.Valid {
			return fmt.Sprintf("It would seem that %s already has reality manifest.", dir_full)
		}
		rb.room.data.Sw = to_id
	case enum.Up:
		if rb.room.data.U.Valid {
			return fmt.Sprintf("It would seem that %s already has reality manifest.", dir_full)
		}
		rb.room.data.U = to_id
	case enum.Down:
		if rb.room.data.D.Valid {
			return fmt.Sprintf("It would seem that %s already has reality manifest.", dir_full)
		}
		rb.room.data.D = to_id
	}
	return ""
}
