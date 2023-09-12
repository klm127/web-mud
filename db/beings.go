package db

import (
	"context"
	"database/sql"

	"github.com/pwsdc/web-mud/arg"
	"github.com/pwsdc/web-mud/db/dbg"
)

func (s *tStore) NewBeingInStartRoom(name string) (*dbg.MudBeing, error) {
	params := dbg.CreateBeingParams{
		Name:        name,
		Description: "",
		Room:        arg.Config.World.StartRoom(),
	}
	being, err := s.Query.CreateBeing(context.Background(), &params)
	if err != nil {
		s.Error(err.Error())
		return nil, err
	}
	s.Logf("Created being for %s.", name)
	return &being, err
}

func (s *tStore) SetBeingOwner(being *dbg.MudBeing, ownerID int64) error {
	params := dbg.UpdateBeingOwnerParams{
		ID: being.ID,
		Owner: sql.NullInt64{
			Valid: true,
			Int64: ownerID,
		},
	}
	err := s.Query.UpdateBeingOwner(context.Background(), &params)
	return err
}

func (s *tStore) DeleteBeingEntry(being_id int64) error {
	return nil
}
