package actor

import (
	"fmt"
	"strings"

	"github.com/pwsdc/web-mud/interfaces/iserver/iactor"
)

func (actor *Actor) commandWithDomain(perstring string, rest string) {
	parts := strings.SplitN(perstring, ".", 2)
	cset := parts[0]
	cmd := parts[1]
	for _, v := range actor.commands {
		if v.GetName() == cset {
			if v.HasCommandOrAlias(cmd) {
				v.Execute(actor, cmd, rest)
			} else {
				actor.ErrorMessage(fmt.Sprintf("I couldn't find a command named '%s' in '%s'. Sorry.", cmd, cset))
			}
			return
		}
	}
	actor.ErrorMessage(fmt.Sprintf("I couldn't find find a command named '%s' in '%s'. Sorry.", cmd, cset))
}

func (actor *Actor) anyMatchingCommand(comname string, rest string) {
	matches := make([]iactor.ICommandGroup, 0, 1)
	for _, v := range actor.commands {
		if v.HasCommandOrAlias(comname) {
			matches = append(matches, v)
		}
	}
	if len(matches) == 1 {
		matches[0].Execute(actor, comname, rest)
	} else if len(matches) < 1 {
		actor.ErrorMessage(fmt.Sprintf("I couldn't find a command named '%s'. Try 'help' to see commands.", comname))
	} else {
		cset_names := make([]string, len(matches))
		for i, v := range matches {
			cset_names[i] = v.GetName()
		}
		joined := strings.Join(cset_names, ", ")

		actor.ErrorMessage(fmt.Sprintf("There are multiple commands named %s. You'll need to qualify it with one of the following: %s. For example, try %s.%s.", comname, joined, cset_names[0], comname))
	}

}
