package prompts

import (
	"strings"

	"github.com/pwsdc/web-mud/server/actor/message"
	"github.com/pwsdc/web-mud/shared"
)

type sQuestion struct {
	qkey         string
	questionText string
	validator    func(string) (bool, string)
}

func (sq *sQuestion) Ask(actor shared.IActor) {
	actor.MessageSimple(sq.questionText)
}

func (sq *sQuestion) GetKey() string {
	return sq.qkey
}

func (sq *sQuestion) Answer(actor shared.IActor, msg string, qr shared.IQuestionResult) {
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
	return msg == "cancel"
}
