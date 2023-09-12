package prompt

import "github.com/pwsdc/web-mud/interfaces/iserver/iactor"

// An interrogatory builder. Call functions to set fields, then call Get() when you are done to retrieve the IInterrogatory.
type interrogatoryBuilder struct {
	inter *Interrogatory
}

func default_onFinished(actor iactor.IActor, qr iactor.IQuestionResult) {
	actor.MessageSimple("I don't care.")
}
func default_onCancel(actor iactor.IActor, qr iactor.IQuestionResult) {
	actor.MessageSimple("Fine, I'll stop asking.")
}

func NewInterrogatory() *interrogatoryBuilder {
	return &interrogatoryBuilder{
		inter: &Interrogatory{
			intro_msg:  "I have some questions.",
			questions:  make([]iactor.IQuestion, 0),
			onFinished: default_onFinished,
			onCancel:   default_onCancel,
		},
	}
}

func (ib *interrogatoryBuilder) Intro(msg string) *interrogatoryBuilder {
	ib.inter.intro_msg = msg
	return ib
}

func (ib *interrogatoryBuilder) OnFinished(cb func(actor iactor.IActor, qr iactor.IQuestionResult)) *interrogatoryBuilder {
	ib.inter.onFinished = cb
	return ib
}

func (ib *interrogatoryBuilder) OnCancel(cb func(actor iactor.IActor, qr iactor.IQuestionResult)) *interrogatoryBuilder {
	ib.inter.onCancel = cb
	return ib
}

func (ib *interrogatoryBuilder) AddQuestion(q iactor.IQuestion) *interrogatoryBuilder {
	ib.inter.questions = append(ib.inter.questions, q)
	return ib
}

func (ib *interrogatoryBuilder) AddQuestions(q ...iactor.IQuestion) *interrogatoryBuilder {
	for _, v := range q {
		ib.inter.questions = append(ib.inter.questions, v)
	}
	return ib
}

func (ib *interrogatoryBuilder) Get() iactor.IInterrogatory {
	return ib.inter
}
