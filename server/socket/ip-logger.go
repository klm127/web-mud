package socket

import (
	"github.com/pwsdc/web-mud/shared"
)

type ipLogger struct {
	shared.HasResults
}

var IpLogger ipLogger

func init() {
	IpLogger = ipLogger{}
	shared.HasResultsInit(&IpLogger)
}
