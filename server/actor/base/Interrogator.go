package base

type Interragator struct {
	intro_msg  string
	questions  []Question
	onFinished func(*Actor, *QuestionResult)
	onCancel   func(*Actor, *QuestionResult)
}

func (i *Interragator) SendIntro(actor *Actor) {
	actor.MessageSimple(i.intro_msg)
}

func (i *Interragator) StartInterragator(actor *Actor) {
	qr := QuestionResult{i, 0, make(map[string]string)}
	actor.questioning = &qr
	i.SendIntro(actor)
	qr.AskNext(actor)
}
