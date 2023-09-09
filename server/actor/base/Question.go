package base

import (
	"strings"

	"github.com/pwsdc/web-mud/server/actor/message"
)

type Question interface {
	GetKey() string
	Ask(*Actor)
	Answer(*Actor, string, *QuestionResult)
}

type sQuestion struct {
	qkey         string
	questionText string
	validator    func(string) (bool, string)
}

func (sq *sQuestion) Ask(actor *Actor) {
	actor.MessageSimple(sq.questionText)
}

func (sq *sQuestion) GetKey() string {
	return sq.qkey
}

func (sq *sQuestion) Answer(actor *Actor, msg string, qr *QuestionResult) {
	canceled := sq.checkCancel(msg)
	if canceled {
		qr.Cancel(actor)
		return
	}
	if sq.validator != nil {
		good, why := sq.validator(msg)
		if !good {
			m := message.New()
			m.Text(why).NewLine(1).Next()
			m.Text("You can always cancel by typing ").Next()
			m.Text("cancel").Link("cancel").NewLine(1).Next()
			actor.Message(m.Bytes())
			return
		}
	}
	qr.AnswerProvided(sq.qkey, msg)
}

func (sq *sQuestion) checkCancel(msg string) bool {
	msg = strings.ToLower(strings.TrimSpace(msg))
	if msg == "cancel" {
		return true
	}
	return false
}
