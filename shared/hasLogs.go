package shared

type HasLogs struct {
	logs []log
}

type iHasLogs interface {
	hasLogsInit()
	Log(txt string)
	GetLogs() *[]string
}

func (self *HasLogs) hasLogsInit() {
	self.logs = make([]log, 0, 10)
}

func (self *HasLogs) Log(txt string) {
	self.logs = append(self.logs, newlog(txt))
}

func (self *HasLogs) GetLogs() *[]string {
	lgs := make([]string, 0, len(self.logs))
	for _, l := range self.logs {
		lgs = append(lgs, l.String())
	}
	return &lgs
}

func HasLogsInit(toInit iHasLogs) {
	toInit.hasLogsInit()
}
