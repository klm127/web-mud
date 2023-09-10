package base

type Interrogator struct {
	intro_msg  string
	questions  []Question
	onFinished func(*Actor, *QuestionResult)
	onCancel   func(*Actor, *QuestionResult)
}

func (i *Interrogator) SendIntro(actor *Actor) {
	actor.MessageSimple(i.intro_msg)
}

func (i *Interrogator) StartInterragator(actor *Actor) {
	qr := QuestionResult{i, 0, make(map[string]string)}
	actor.questioning = &qr
	i.SendIntro(actor)
	qr.AskNext(actor)
}
