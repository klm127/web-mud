package shared

type IInterogator interface {
	StartInterragator(IActor)
	SendIntro(ISendsMessages)
	Finished(IActor, IQuestionResult)
	Cancel(IActor, IQuestionResult)
	Ask(actor IActor, index int)
	Answer(actor IActor, index int, msg string, qr IQuestionResult)
	QuestionsCount() int
}

type IQuestionResult interface {
	Interrogator() IInterogator
	InputReceived(IActor, string)
	AnswerProvided(q_key string, q_answer string)
	Cancel(IActor)
	GetResult() map[string]string
}

type IQuestion interface {
	GetKey() string
	Ask(IActor)
	Answer(IActor, string)
}
