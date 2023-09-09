package commands

import (
	"github.com/pwsdc/web-mud/server/actor/base"
	"github.com/pwsdc/web-mud/server/actor/prompts"
)

func init() {
	cs := base.NewCommandSet("test")
	tq_cmd := base.NewCommand("questions", "runs some test questions", nil, startTestQuestions)
	cs.RegisterCommand(tq_cmd)
	base.RegisterDefaultCommandSet(cs)
}

func startTestQuestions(actor *base.Actor, msg string) {
	prompts.TestQuestions.StartInterragator(actor)
}
