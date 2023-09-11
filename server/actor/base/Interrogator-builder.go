package base

import "fmt"

type InterragatorBuilder struct {
	inter *Interrogator
}

func simpleCancel(actor *Actor, qr *QuestionResult) {
	actor.questioning = nil
	actor.MessageSimple("I'm done with this line of questioning.")
}

func nullopOnDone(actor *Actor, qr *QuestionResult) {
	fmt.Println("An interragator needs a completion handler!", qr)
	actor.MessageSimple("Thanks for answering my questions.")
	actor.questioning = nil
}

// Get a builder for an interragator
func NewInter() *InterragatorBuilder {
	inter := Interrogator{}
	inter.onCancel = simpleCancel
	inter.onFinished = nullopOnDone
	inter.questions = make([]Question, 0)
	ib := InterragatorBuilder{&inter}
	return &ib
}

// set the interragatory intro text
func (ib *InterragatorBuilder) Intro(text string) *InterragatorBuilder {
	ib.inter.intro_msg = text
	return ib
}

// add a question to the chain
func (ib *InterragatorBuilder) AddQuestion(q Question) *InterragatorBuilder {
	ib.inter.questions = append(ib.inter.questions, q)
	return ib
}

// set the function that will be called when the interragatory is finished
func (ib *InterragatorBuilder) OnFinished(cb func(*Actor, *QuestionResult)) *InterragatorBuilder {
	ib.inter.onFinished = cb
	return ib
}

// set the function that will be called when the interragatory is canceled
func (ib *InterragatorBuilder) OnCancel(cb func(*Actor, *QuestionResult)) *InterragatorBuilder {
	ib.inter.onCancel = cb
	return ib
}

func (ib *InterragatorBuilder) Get() *Interrogator {
	return ib.inter
}
