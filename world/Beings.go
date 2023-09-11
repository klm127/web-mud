package world

import (
	"context"
	"database/sql"

	"github.com/pwsdc/web-mud/db"
	"github.com/pwsdc/web-mud/db/dbg"
)

var loadedBeings map[int64]*Being

func init() {
	loadedBeings = make(map[int64]*Being)
}

func GetBeing(id int64) *Being {
	being, ok := loadedBeings[id]
	if ok {
		return being
	}
	db_being, err := db.Store.Query.GetBeingById(context.Background(), id)
	if err != nil {
		return nil
	}
	a_being := createBeing(&db_being)
	loadedBeings[id] = a_being
	return a_being
}

func NewBeing(name string, description string, room int64, owner *int64) *int64 {
	params := dbg.CreateBeingParams{
		Name:        name,
		Description: description,
		Room:        room,
		Owner:       sql.NullInt64{},
	}
	if owner == nil {
		params.Owner.Valid = false
	} else {
		params.Owner.Valid = true
		params.Owner.Int64 = *owner
	}
	being_db, err := db.Store.Query.CreateBeing(context.Background(), &params)
	if err != nil {
		return nil
	}
	being := Being{
		data:  &being_db,
		dirty: false,
	}
	loadedBeings[being_db.ID] = &being
	return &being_db.ID
}

func UpdateBeingOwner(id int64, owner int64) {
	being := GetBeing(id)
	if being != nil {
		being.data.ID = owner
		params := dbg.UpdateBeingOwnerParams{
			Owner: sql.NullInt64{Valid: true, Int64: id},
			ID:    id,
		}
		db.Store.Query.UpdateBeingOwner(context.Background(), &params)
	}

}

func DeleteBeing(id int64) {
	delete(loadedBeings, id)
	db.Store.Query.DeleteBeing(context.Background(), id)
}

func OffloadBeing(id int64) {
	b, ok := loadedBeings[id]
	if ok {
		b.UpdateToDatabase()
		delete(loadedBeings, id)
	}
}
