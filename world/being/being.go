package being

import (
	"context"

	"github.com/pwsdc/web-mud/db"
	"github.com/pwsdc/web-mud/db/dbg"
	"github.com/pwsdc/web-mud/shared"
)

type Being struct {
	shared.HasResults
	data  *dbg.MudBeing
	dirty bool
}

func _InitBeing(data *dbg.MudBeing) *Being {
	b := Being{
		data:  data,
		dirty: false,
	}
	shared.HasResultsInit(&b)
	return &b
}

func (b *Being) GetData() *dbg.MudBeing {
	return b.data
}

func (b *Being) Offload() {
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
