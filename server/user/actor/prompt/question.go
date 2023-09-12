package prompt

import (
	"strings"

	"github.com/pwsdc/web-mud/interfaces/iserver/iactor"
	"github.com/pwsdc/web-mud/server/user/actor/message"
)

type Question struct {
	key        string
	text       string
	validators []func(string) (bool, string)
}

func (q *Question) Ask(actor iactor.IActor) {
	actor.MessageSimple(q.text)
}

func (q *Question) GetKey() string {
	return q.key
}

func (q *Question) Answer(actor iactor.IActor, msg string, qr iactor.IQuestionResult) {
	if q.checkCancel(msg) {
		qr.Cancel(actor)
		return
	}
	if len(q.validators) != 0 {
		for _, v := range q.validators {
			good, why := v(msg)
			if !good {
				m := message.New()
				m.Text(why).NewLine(1).Next()
				m.Text("You can always cancel by typing ").Next()
				m.Text("cancel").Link("cancel").NewLine(1).Next()
				actor.Message(m.Bytes())
				return
			}
		}
	}
	qr.AnswerProvided(q.key, msg)
}

func (q *Question) checkCancel(msg string) bool {
	msg = strings.ToLower(strings.TrimSpace(msg))
	return msg == "cancel"
}
