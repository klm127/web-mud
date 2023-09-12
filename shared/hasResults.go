package shared

import (
	"github.com/pwsdc/web-mud/interfaces/iserver/iactor"
	"github.com/pwsdc/web-mud/server/user/actor/message"
)

type HasResults struct {
	HasErrors
	HasLogs
}

type iHasResults interface {
	iHasErrors
	iHasLogs
	MessageResults(iactor.IActor)
}

func HasResultsInit(toInit iHasResults) {
	HasErrorsInit(toInit)
	HasLogsInit(toInit)
}

func (hl HasResults) MessageResults(actor iactor.IActor) {
	mb := message.New()
	for _, v := range hl.logs {
		mb.Textf("%s: %s", v.Time().Format("2006-01-02 15:04:05"), v.String()).NewLine(1).Next()
	}
	for _, v := range hl.errors {
		mb.Textf("%s: %s", v.Time().Format("2006-01-02 15:04:05"), v.String()).NewLine(1).Color("red").Next()
	}
	actor.Message(mb.Bytes())
}
