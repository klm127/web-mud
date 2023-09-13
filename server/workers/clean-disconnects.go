package workers

import (
	"fmt"
	"time"

	"github.com/pwsdc/web-mud/arg"
	"github.com/pwsdc/web-mud/interfaces/iserver/iactor"
	"github.com/pwsdc/web-mud/server/user/actor"
)

var idle_cleaner_running bool
var idle_dc_msg string

func StartIdleConnectionCleaner() {
	interval := arg.Config.Socket.IdleCheckInterval()
	timeout := float64(arg.Config.Socket.IdleTimeout())
	idle_cleaner_running = true
	traverser := getActorIdleTraverser(timeout)
	for idle_cleaner_running {
		actor.Traverse(traverser, false)
		time.Sleep(time.Duration(interval) * time.Second)
	}
}

func getActorIdleTraverser(timeout_mins float64) func(*map[int64]iactor.IActor) {
	bye_msg := fmt.Sprintf("You are being disconnected because you have been idle for more than %v minutes.", timeout_mins)
	return func(actors *map[int64]iactor.IActor) {
		amap := *actors
		for _, an_actor := range amap {
			time_since := an_actor.GetTimeSinceLastTalked()
			if time_since.Minutes() > timeout_mins {
				an_actor.MessageSimple(bye_msg)
				an_actor.Disconnect()
			}
		}
	}
}

func StopIdleConnectionCleaner() {
	idle_cleaner_running = false
}
