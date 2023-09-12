package basecommands

import (
	"fmt"
	"strings"

	"github.com/pwsdc/web-mud/interfaces/iserver/iactor"
	"github.com/pwsdc/web-mud/server/user/actor/command"
	"github.com/pwsdc/web-mud/server/user/actor/message"
	"github.com/pwsdc/web-mud/util/re"
)

/*
	The info set provides general help on other commands, time, and other services.
*/

var UserInfoCommands iactor.ICommandGroup

func init() {
	UserInfoCommands = command.NewCommandGroup("info")
	help_cmd := command.NewCommand().Name("help").Desc("gets help for a command").OnExec(help).Get()
	UserInfoCommands.RegisterCommands(help_cmd)
}

func help(actor iactor.IActor, msg string) {
	msg = strings.TrimSpace(msg)
	m := message.New()
	if len(msg) == 0 {
		m.Text("Let's get you some help. You can type 'help <command>' to get more information about it.").NewLine(1).Next()
		m.Text("The commands you have available are:").NewLine(1)

		command_sets := actor.GetCommandGroups()

		// print all available commands
		for _, cset := range command_sets {
			m.Next()
			// print the category name
			m.Text(cset.GetName() + ":").Classes([]string{"bolded", "right-pad"}).Next()
			cset_commands := *cset.GetCommands()
			cat_commands := make([]string, len(cset_commands))
			i := 0
			for _, command := range cset_commands {
				command_name := command.GetName()
				aliases := *command.GetAliases()
				if len(aliases) > 0 {
					command_name += fmt.Sprintf(" (%s)", strings.Join(aliases, ","))
				}
				cat_commands[i] = command_name
				i++
			}
			m.Text(strings.Join(cat_commands, ", ")).Indent(10).NewLine(1)
		}
	} else {
		helpParticular(actor, msg)
	}
	actor.Message(m.Bytes())
}

func helpParticular(actor iactor.IActor, cmd string) {
	if re.HasPeriod.Match([]byte(cmd)) {
		split := strings.SplitN(cmd, ".", 2)
		if len(split) < 2 {
			actor.Errorf("I need a command to look for.")
		} else {
			cset, ok := actor.GetCommandGroup(split[0])
			if !ok {
				actor.Errorf("I couldn't find the command group %s.", split[0])
			} else {
				has_cmd := cset.HasCommandOrAlias(split[1])
				if !has_cmd {
					actor.Errorf("Command group %s doesn't seem to have a command called %s.", split[0], split[1])
				} else {
					m := message.New().Textf("Here's some help for %s.", cmd).NewLine(1).Next()
					m.Indent(10).Textf("It %s.", cset.GetCommand(split[1]).GetDescription())
					actor.Message(m.Bytes())
				}
			}
		}
	} else {
		all_cmd_groups := actor.GetCommandGroups()
		matches := make([]iactor.ICommandGroup, 0, 1)
		for _, v := range all_cmd_groups {
			if v.HasCommandOrAlias(cmd) {
				matches = append(matches, v)
			}
		}
		if len(matches) == 1 {
			m := message.New().Textf("Here's some help for %s.", cmd).NewLine(1).Next()
			m.Indent(10).Textf("It %s.", matches[0].GetCommand(cmd).GetDescription())
			actor.Message(m.Bytes())
		} else if len(matches) < 1 {
			actor.Errorf("I couldn't find any commands named %s.", cmd)
		} else {
			cset_names := make([]string, len(matches))
			for i, v := range matches {
				cset_names[i] = v.GetName()
			}
			joined := strings.Join(cset_names, ", ")

			actor.Errorf("There are multiple commands named %s. You'll need to qualify it with one of the following: %s. For example, try %s.%s.", cmd, joined, cset_names[0], cmd)
		}
	}
}
