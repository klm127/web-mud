package prompts

import "github.com/pwsdc/web-mud/server/actor/base"

var TestQuestions *base.Interragator

func init() {
	ib := base.NewInter().Intro("Some test questions.")
	qname := base.NewQuestion().Key("name").Text("What's your name?")
	qcolor := base.NewQuestion().Key("color").Text("What's your favorite color?").MultipleChoice([]string{"red", "blue", "yellow"})
	ib.AddQuestion(qname.Get())
	ib.AddQuestion(qcolor.Get())
	TestQuestions = ib.Get()
}
