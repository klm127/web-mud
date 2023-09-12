package iactor

// A set of questions. IQuestionResult is closely coupled. An IInterrogatory instances IQuestionResults for each actor using it.
type IInterrogatory interface {
	// Initializes a new question result and returns it. This can be associated with the particular actor using this interrogatory.
	GetQuestioner(actor IActor) IQuestionResult
	// Sends the intro message to this line of questioning to an actor.
	SendIntro(actor IActor)
	// Finishes this line of questioning, calling the callback function.
	Finish(actor IActor, result IQuestionResult)
	// Cancels this line of questioning, calling hte cancel callback function.
	Cancel(actor IActor, result IQuestionResult)
	// Gets the question with a particular index. Returns nil if out of range.
	GetQuestion(index int) IQuestion
	// Gets the total number of questions.
	GetQuestionCount() int
}
