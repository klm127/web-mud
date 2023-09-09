package commands

import (
	"fmt"
	"strings"
	"time"

	"github.com/pwsdc/web-mud/server/actor/base"
	"github.com/pwsdc/web-mud/server/actor/message"
)

func init() {
	createInfoSet()
}

func createInfoSet() {
	cs := base.NewCommandSet("info")
	help_cmd := base.NewCommand("help", "shows information about available commands", []string{"h"}, help)
	time_cmd := base.NewCommand("time", "shows information about time.", []string{}, timeCommand)
	cs.RegisterCommand(help_cmd)
	cs.RegisterCommand(time_cmd)
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
					s += fmt.Sprintf("(%s)", strings.Join(com.Alias, ","))
				}
				s += "   "
			}
			m.Text(s).Indent(10)
		}
	}
	actor.Message(m.Bytes())
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
