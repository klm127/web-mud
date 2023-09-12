package iworld

import (
	"github.com/pwsdc/web-mud/db/dbg"
)

type IBeing interface {
	GetData() *dbg.MudBeing
	Offload()
}
