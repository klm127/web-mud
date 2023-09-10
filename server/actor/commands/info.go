package commands

import (
	"fmt"
	"strings"
	"time"

	"github.com/pwsdc/web-mud/server/actor/base"
	"github.com/pwsdc/web-mud/server/actor/message"
	"github.com/pwsdc/web-mud/util/re"
)

func init() {
	createInfoSet()
}

func createInfoSet() {
	cs := base.NewCommandSet("info")
	help_cmd := base.NewCommand("help", "shows information about available commands", []string{"h"}, help)
	time_cmd := base.NewCommand("time", "shows information about time", []string{}, timeCommand)
	cs.RegisterCommand(help_cmd)
	cs.RegisterCommand(time_cmd)
	base.RegisterDefaultCommandSet(cs)
}

func help(actor *base.Actor, msg string) {
	msg = strings.TrimSpace(msg)
	m := message.New()
	if len(msg) == 0 {
		m.Text("Let's get you some help. You can type 'help <command>' to get more information about it.").NewLine(1).Next()
		m.Text("The commands you have available are:").NewLine(1)
		// print all available commands
		for _, cset := range actor.Commands {
			m.Next()
			// print the category name
			m.Text(cset.Name + ":").Classes([]string{"bolded", "right-pad"}).Next()
			cset_commands := cset.GetCommands()
			cat_commands := make([]string, len(cset_commands))
			i := 0
			for _, com := range cset_commands {
				s := com.Name
				if len(com.Alias) > 0 {
					s += fmt.Sprintf(" (%s)", strings.Join(com.Alias, ","))
				}
				cat_commands[i] = s
				i++
			}
			m.Text(strings.Join(cat_commands, ", ")).Indent(10).NewLine(1)
		}
	} else {
		helpParticular(actor, msg)
	}
	actor.Message(m.Bytes())
}

func helpParticular(actor *base.Actor, cmd string) {
	if re.HasPeriod.Match([]byte(cmd)) {
		split := strings.SplitN(cmd, ".", 2)
		if len(split) < 2 {
			actor.Errorf("I need a command to look for.")
		} else {
			cset, ok := actor.Commands[split[0]]
			if !ok {
				actor.Errorf("I couldn't find the command group %s.", split[0])
			} else {
				has_cmd := cset.HasCommandOrAlias(split[1])
				if !has_cmd {
					actor.Errorf("Command group %s doesn't seem to have a command called %s.", split[0], split[1])
				} else {
					m := message.New().Textf("Here's some help for %s.", cmd).NewLine(1).Next()
					m.Indent(10).Textf("It %s.", cset.GetCommand(split[1]).Description)
					actor.Message(m.Bytes())
				}
			}
		}
	} else {
		matches := make([]*base.CommandSet, 0, 1)
		for _, v := range actor.Commands {
			if v.HasCommandOrAlias(cmd) {
				matches = append(matches, v)
			}
		}
		if len(matches) == 1 {
			m := message.New().Textf("Here's some help for %s.", cmd).NewLine(1).Next()
			m.Indent(10).Textf("It %s.", matches[0].GetCommand(cmd).Description)
			actor.Message(m.Bytes())
		} else if len(matches) < 1 {
			actor.Errorf("I couldn't find any commands named %s.", cmd)
		} else {
			cset_names := make([]string, len(matches))
			for i, v := range matches {
				cset_names[i] = v.Name
			}
			joined := strings.Join(cset_names, ", ")

			actor.Errorf("There are multiple commands named %s. You'll need to qualify it with one of the following: %s. For example, try %s.%s.", cmd, joined, cset_names[0], cmd)
		}
	}

}

func timeCommand(actor *base.Actor, msg string) {
	send_time_opened := true
	send_time_last_talk := true
	send_time_now := true
	msg = strings.TrimSpace(msg)
	response := message.New()
	if len(msg) > 0 {
		send_time_opened = false
		send_time_last_talk = false
		send_time_now = false
		if msg == "now" {
			send_time_now = true
		} else if msg == "start" {
			send_time_opened = true
		} else if msg == "last" {
			send_time_last_talk = true
		} else {
			response.Text("I didn't understand. You can say 'time now', 'time start', 'time now', or just 'time'.")
			actor.Message(response.Bytes())
			return
		}
	}
	if send_time_opened {
		s := fmt.Sprintf("You connected at %s.", actor.GetTimeOpened().Format("2006-01-02 15:04:05.000"))
		response.Text(s).NewLine(1).Indent(10).Next()
		ot := time.Now().Sub(actor.GetTimeOpened())
		s2 := fmt.Sprintf("The connection has been open for %.2f minutes.", ot.Minutes())
		response.Text(s2).NewLine(1).Indent(10).Next()
	}
	if send_time_last_talk {
		s := fmt.Sprintf("We last communicated at %s.", actor.GetTimeLastTalked().Format("2006-01-02 15:04:05.000"))
		response.Text(s).NewLine(1).Indent(10).Next()
	}
	if send_time_now {
		s := fmt.Sprintf("The current time is %s.", time.Now().Format("2006-01-02 15:04:05.000"))
		response.Text(s).NewLine(1).Indent(10).Next()
	}
	actor.Message(response.Bytes())

}
