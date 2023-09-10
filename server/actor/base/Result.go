package base

type QuestionResult struct {
	interragator  *Interrogator
	questionIndex int
	result        map[string]string
}

func (qr *QuestionResult) AskNext(actor *Actor) {
	if qr.questionIndex >= len(qr.interragator.questions) {
		qr.interragator.onFinished(actor, qr)
		return
	}
	qr.interragator.questions[qr.questionIndex].Ask(actor)
}

func (qr *QuestionResult) InputReceived(actor *Actor, msg string) {
	qr.interragator.questions[qr.questionIndex].Answer(actor, msg, qr)
}

func (qr *QuestionResult) AnswerProvided(q_key string, q_answer string) {
	qr.result[q_key] = q_answer
	qr.questionIndex += 1
}

func (qr *QuestionResult) Cancel(actor *Actor) {
	qr.interragator.onCancel(actor, qr)
}

func NewQuestioner(interragator *Interrogator) *QuestionResult {
	qr := QuestionResult{interragator, 0, make(map[string]string)}
	return &qr
}

func (qr *QuestionResult) GetResult() map[string]string {
	return qr.result
}
