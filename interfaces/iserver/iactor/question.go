package iactor

type IQuestion interface {
	Answer(IActor, string, IQuestionResult)
	Ask(IActor)
}
