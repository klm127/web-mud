package shared

type HasErrors struct {
	errors []log
}
type iHasErrors interface {
	hasErrorsInit()
	Error(txt string)
	GetErrors() *[]string
}

func (he *HasErrors) hasErrorsInit() {
	he.errors = make([]log, 0, 10)
}

func (he *HasErrors) Error(txt string) {
	he.errors = append(he.errors, newlog(txt))
}

func (he *HasErrors) GetErrors() *[]string {
	ers := make([]string, 0, len(he.errors))
	for _, l := range he.errors {
		ers = append(ers, l.String())
	}
	return &ers
}

func HasErrorsInit(toInit iHasErrors) {
	toInit.hasErrorsInit()
}
