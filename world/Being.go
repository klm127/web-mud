package world

import (
	"context"

	"github.com/pwsdc/web-mud/db"
	"github.com/pwsdc/web-mud/db/dbg"
	"github.com/pwsdc/web-mud/shared"
)

type Being struct {
	shared.HasLogs
	data  *dbg.MudBeing
	dirty bool
}

func createBeing(data *dbg.MudBeing) *Being {
	b := Being{
		data:  data,
		dirty: false,
	}
	shared.HasLogsInit(&b)
	return &b
}

func (b *Being) UpdateToDatabase() {
	if b.dirty {
		params := dbg.UpdateBeingParams{
			Description: b.data.Description,
			Room:        b.data.Room,
		}
		err := db.Store.Query.UpdateBeing(context.Background(), &params)
		if err != nil {
			b.dirty = false
		} else {
			b.Logf("Error updating database: %s", err.Error())
		}
	}
}
