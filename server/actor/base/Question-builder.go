package base

type SimpleQuestionBuilder struct {
	question *sQuestion
}

func NewQuestion() *SimpleQuestionBuilder {
	q := sQuestion{}
	sqb := SimpleQuestionBuilder{&q}
	return &sqb
}

func (sqb *SimpleQuestionBuilder) Get() Question {
	return sqb.question
}

// Set the question key; what answer will be in the result map.
func (sqb *SimpleQuestionBuilder) Key(val string) *SimpleQuestionBuilder {
	sqb.question.qkey = val
	return sqb
}

// Set the question text.
func (sqb *SimpleQuestionBuilder) Text(val string) *SimpleQuestionBuilder {
	sqb.question.questionText = val
	return sqb
}

// Set this question to validate that it's an integer answer.
func (sqb *SimpleQuestionBuilder) Int() *SimpleQuestionBuilder {
	sqb.question.validator = intValidator
	return sqb
}

// Set this question to validate that answer will be a postive integer
func (sqb *SimpleQuestionBuilder) PositiveInt() *SimpleQuestionBuilder {
	sqb.question.validator = positiveIntValidator
	return sqb
}

// Set this question to validate that answer will be an integer between start and end
func (sqb *SimpleQuestionBuilder) IntBetween(start int, end int) *SimpleQuestionBuilder {
	sqb.question.validator = intBetweenValidator(start, end)
	return sqb
}

// Set this question to validate that it's a float number answer.
func (sqb *SimpleQuestionBuilder) Float() *SimpleQuestionBuilder {
	sqb.question.validator = floatValidator
	return sqb
}

// Set this question to validate that answer is a float number between start and end
func (sqb *SimpleQuestionBuilder) FloatBetween(start float32, end float32) *SimpleQuestionBuilder {
	sqb.question.validator = floatBetweenValidator(start, end)
	return sqb
}

// Set this question to validate that it's a yes/no answer.
func (sqb *SimpleQuestionBuilder) YesNo() *SimpleQuestionBuilder {
	sqb.question.validator = yesNoValidator
	return sqb
}

// Set this question to validate that the answer will be on of the choices
func (sqb *SimpleQuestionBuilder) MultipleChoice(choices []string) *SimpleQuestionBuilder {
	sqb.question.validator = multiChoiceValidator(choices)
	return sqb
}
