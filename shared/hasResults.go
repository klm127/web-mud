package shared

type HasResults struct {
	HasErrors
	HasLogs
}

type iHasResults interface {
	iHasErrors
	iHasLogs
}

func HasResultsInit(toInit iHasResults) {
	HasErrorsInit(toInit)
	HasLogsInit(toInit)
}
