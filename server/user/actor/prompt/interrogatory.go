package prompt

import "github.com/pwsdc/web-mud/interfaces/iserver/iactor"

type Interrogatory struct {
	intro_msg  string
	questions  []iactor.IQuestion
	onFinished func(iactor.IActor, iactor.IQuestionResult)
	onCancel   func(iactor.IActor, iactor.IQuestionResult)
}

func (inter *Interrogatory) GetQuestioner(actor iactor.IActor) iactor.IQuestionResult {
	return NewQuestionResult(actor, inter)
}

func (inter *Interrogatory) SendIntro(actor iactor.IActor) {
	actor.MessageSimplef(inter.intro_msg)
}

func (inter *Interrogatory) Finish(actor iactor.IActor, result iactor.IQuestionResult) {
	inter.onFinished(actor, result)
}

func (inter *Interrogatory) Cancel(actor iactor.IActor, result iactor.IQuestionResult) {
	actor.EndQuestioning()
	inter.onCancel(actor, result)
}

func (inter *Interrogatory) GetQuestion(index int) iactor.IQuestion {
	if index < len(inter.questions) {
		return inter.questions[index]
	}
	return nil
}

func (inter *Interrogatory) GetQuestionCount() int {
	return len(inter.questions)
}
