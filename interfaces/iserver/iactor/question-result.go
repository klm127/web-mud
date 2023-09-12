package iactor

// Both the result(the users answers) and the references to which questions need to be asked next. Instanced for a specific actor when a new line of questioning begins.
type IQuestionResult interface {
	// Called when input is received.
	InputReceived(input string)
	// Asks the next question. If the last question has been asked, finishes the interrogatory instead.
	AskNext()
	// Called when an answer to a question referenced by key is provided.
	AnswerProvided(key string, answer string)
	// Cancels this line of questions.
	Cancel(actor IActor)
	// Gets all answers as a map of the form question_key->user_answer
	GetResult() *map[string]string
}
