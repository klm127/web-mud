package prompt

import (
	"fmt"

	"github.com/pwsdc/web-mud/interfaces/iserver/iactor"
)

type questionBuilder struct {
	question *Question
}

func NewQuestion(key string) *questionBuilder {
	return &questionBuilder{
		question: &Question{
			key:        key,
			text:       "I think I should ask you something?",
			validators: make([]func(string) (bool, string), 0),
		},
	}
}

// Any custom validator
func (sqb *questionBuilder) CustomValidator(cb func(string) (bool, string)) *questionBuilder {
	sqb.question.validators = append(sqb.question.validators, cb)
	return sqb
}

func (sqb *questionBuilder) Get() iactor.IQuestion {
	return sqb.question
}

// Set the question key; what answer will be in the result map.
func (sqb *questionBuilder) Key(val string) *questionBuilder {
	sqb.question.key = val
	return sqb
}

// Set the question text.
func (sqb *questionBuilder) Text(val string) *questionBuilder {
	sqb.question.text = val
	return sqb
}

func (sqb *questionBuilder) Textf(format string, a ...any) *questionBuilder {
	sqb.question.text = fmt.Sprintf(format, a...)
	return sqb
}

// Set this question to validate that it's an integer answer.
func (sqb *questionBuilder) Int() *questionBuilder {
	sqb.question.validators = append(sqb.question.validators, intValidator)
	return sqb
}

// Set this question to validate that answer will be a postive integer
func (sqb *questionBuilder) PositiveInt() *questionBuilder {
	sqb.question.validators = append(sqb.question.validators, positiveIntValidator)
	return sqb
}

// Set this question to validate that answer will be an integer between start and end
func (sqb *questionBuilder) IntBetween(start int, end int) *questionBuilder {
	sqb.question.validators = append(sqb.question.validators, intBetweenValidator(start, end))
	return sqb
}

// Set this question to validate that it's a float number answer.
func (sqb *questionBuilder) Float() *questionBuilder {
	sqb.question.validators = append(sqb.question.validators, floatValidator)
	return sqb
}

// Set this question to validate that answer is a float number between start and end
func (sqb *questionBuilder) FloatBetween(start float32, end float32) *questionBuilder {
	sqb.question.validators = append(sqb.question.validators, floatBetweenValidator(start, end))
	return sqb
}

// Set this question to validate that it's a yes/no answer.
func (sqb *questionBuilder) YesNo() *questionBuilder {
	sqb.question.validators = append(sqb.question.validators, yesNoValidator)
	return sqb
}

// Set this question to validate that the answer will be on of the choices
func (sqb *questionBuilder) MultipleChoice(choices []string) *questionBuilder {
	sqb.question.validators = append(sqb.question.validators, multiChoiceValidator(choices))
	return sqb
}

// String length validator
func (sqb *questionBuilder) StringLength(minLength int) *questionBuilder {
	sqb.question.validators = append(sqb.question.validators, strlenValidator(minLength))
	return sqb
}
