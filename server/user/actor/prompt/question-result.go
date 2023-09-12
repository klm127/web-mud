package prompt

import "github.com/pwsdc/web-mud/interfaces/iserver/iactor"

// implements iactor.IQuestionResult
type QuestionResult struct {
	interragator  iactor.IInterrogatory
	questionIndex int
	result        map[string]string
	actor         iactor.IActor
}

func NewQuestionResult(actor iactor.IActor, interrogatories iactor.IInterrogatory) *QuestionResult {
	qr := QuestionResult{
		interragator:  interrogatories,
		actor:         actor,
		questionIndex: 0,
		result:        make(map[string]string),
	}
	return &qr
}

func (qr *QuestionResult) InputReceived(input string) {
	q := qr.interragator.GetQuestion(qr.questionIndex)
	q.Answer(qr.actor, input, qr)
}
func (qr *QuestionResult) AskNext() {
	n := qr.interragator.GetQuestionCount()
	if qr.questionIndex >= n {
		qr.interragator.Finish(qr.actor, qr)
		return
	}
	qr.interragator.GetQuestion(qr.questionIndex).Ask(qr.actor)
}
func (qr *QuestionResult) AnswerProvided(key string, answer string) {
	qr.result[key] = answer
	qr.questionIndex += 1
}
func (qr *QuestionResult) Cancel(actor iactor.IActor) {
	qr.interragator.Cancel(actor, qr)
}
func (qr *QuestionResult) GetResult() *map[string]string {
	return &qr.result
}
