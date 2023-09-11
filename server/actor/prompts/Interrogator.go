package prompts

import "github.com/pwsdc/web-mud/shared"

type Interrogator struct {
	intro_msg  string
	questions  []shared.IQuestion
	onFinished func(shared.IActor, *QuestionResult)
	onCancel   func(shared.IActor, *QuestionResult)
}

func (i *Interrogator) StartInterragator(actor shared.IActor) {
	//actor.StartQuestioning(i)
}

func (i *Interrogator) Answer(actor shared.IActor, index int, msg string, qr shared.IQuestionResult) {

}

func (i *Interrogator) SendIntro(actor shared.ISendsMessages) {
	actor.MessageSimple(i.intro_msg)
}

func (i *Interrogator) Finished(shared.IActor, shared.IQuestionResult) {

}
