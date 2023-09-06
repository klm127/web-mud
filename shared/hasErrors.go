package shared

type HasErrors struct {
	errors []log
}
type iHasErrors interface {
	hasErrorsInit()
	Error(txt string)
	GetErrors() *[]string
}

func (self *HasErrors) hasErrorsInit() {
	self.errors = make([]log, 0, 10)
}

func (self *HasErrors) Error(txt string) {
	self.errors = append(self.errors, newlog(txt))
}

func (self *HasErrors) GetErrors() *[]string {
	ers := make([]string, 0, len(self.errors))
	for _, l := range self.errors {
		ers = append(ers, l.String())
	}
	return &ers
}

func HasErrorsInit(toInit iHasErrors) {
	toInit.hasErrorsInit()
}
