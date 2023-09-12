package shared

import (
	"fmt"
)

type HasLogs struct {
	logs []log
}

type iHasLogs interface {
	hasLogsInit()
	Log(txt string)
	GetLogs() *[]string
	GetLogRaw() *[]log
}

func (hl *HasLogs) hasLogsInit() {
	hl.logs = make([]log, 0, 10)
}

func (hl *HasLogs) Log(txt string) {
	hl.logs = append(hl.logs, newlog(txt))
}

func (logr *HasLogs) Logf(format string, a ...any) {
	logr.logs = append(logr.logs, newlog(fmt.Sprintf(format, a...)))
}

func (hl *HasLogs) GetLogs() *[]string {
	lgs := make([]string, 0, len(hl.logs))
	for _, l := range hl.logs {
		lgs = append(lgs, l.String())
	}
	return &lgs
}

func (hl *HasLogs) GetLogRaw() *[]log {
	return &hl.logs
}

func HasLogsInit(toInit iHasLogs) {
	toInit.hasLogsInit()
}
