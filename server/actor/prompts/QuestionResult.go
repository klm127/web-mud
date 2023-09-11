package prompts

import (
	"github.com/pwsdc/web-mud/shared"
)

type QuestionResult struct {
	interragator  shared.IInterogator
	questionIndex int
	result        map[string]string
}

func (qr *QuestionResult) AskNext(actor shared.IActor) {
	qcount := qr.interragator.QuestionsCount()
	if qr.questionIndex >= qcount {
		qr.interragator.Finished(actor, qr)
		return
	}
	qr.interragator.Ask(actor, qr.questionIndex)
}

func (qr *QuestionResult) InputReceived(actor shared.IActor, msg string) {
	qr.interragator.Answer(actor, qr.questionIndex, msg, qr)
}

func (qr *QuestionResult) AnswerProvided(q_key string, q_answer string) {
	qr.result[q_key] = q_answer
	qr.questionIndex += 1
}

func (qr *QuestionResult) Cancel(actor shared.IActor) {
	qr.interragator.Cancel(actor, qr)
}

func (qr *QuestionResult) Interrogator() shared.IInterogator {
	return qr.interragator
}

// func NewQuestioner(interragator *Interrogator) *QuestionResult {
// 	qr := QuestionResult{interragator, 0, make(map[string]string)}
// 	return &qr
// }

func (qr *QuestionResult) GetResult() map[string]string {
	return qr.result
}
