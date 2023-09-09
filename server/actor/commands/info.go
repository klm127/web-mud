package commands

import (
	"fmt"
	"strings"

	"github.com/pwsdc/web-mud/server/actor/base"
	"github.com/pwsdc/web-mud/server/actor/message"
)

func init() {
	createInfoSet()
}

func createInfoSet() {
	cs := base.NewCommandSet("info")
	help_cmd := base.NewCommand("help", "shows information about available commands", []string{"h"}, help)
	cs.RegisterCommand(help_cmd)
	base.RegisterDefaultCommandSet(cs)
}

func help(actor *base.Actor, msg string) {
	msg = strings.TrimSpace(msg)
	m := message.New()
	if len(msg) == 0 {
		m.Text("The commands you have available are:").NewLine(1)
		// print all available commands
		for _, cset := range actor.Commands {
			m.Next()
			cset_commands := cset.GetCommands()
			m.Text(" " + cset.Name).NewLine(1)
			s := ""
			for _, com := range cset_commands {
				s += com.Name
				if len(com.Alias) > 0 {
					s += fmt.Sprintf(" (%s)", strings.Join(com.Alias, ","))
				}
				s += "   "
			}
			m.Text(s).Indent(10)
		}
	}
	actor.Message(m.Bytes())
}
