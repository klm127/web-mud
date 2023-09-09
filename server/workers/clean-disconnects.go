package workers

import (
	"fmt"
	"time"

	"github.com/pwsdc/web-mud/arg"
	"github.com/pwsdc/web-mud/server/actor/base"
)

var idle_cleaner_running bool
var idle_dc_msg string

func StartIdleConnectionCleaner() {
	interval := arg.Config.Socket.IdleCheckInterval()
	timeout := float64(arg.Config.Socket.IdleTimeout())
	idle_cleaner_running = true
	for idle_cleaner_running {
		CleanIdles(timeout)
		time.Sleep(time.Duration(interval) * time.Second)
	}
}

func CleanIdles(idle_time_minutes float64) int {
	bye_msg := fmt.Sprintf("You are being disconnected because you have been idle for more than %v minutes.", idle_time_minutes)
	idles_cleaned := 0
	for _, actor := range base.Actors {
		actor_last_talked := actor.GetTimeSinceLastTalked()
		if actor_last_talked.Minutes() > idle_time_minutes {
			actor.MessageSimple(bye_msg)
			actor.Disconnect()
			idles_cleaned += 1
		}
	}
	return idles_cleaned
}

func StopIdleConnectionCleaner() {
	idle_cleaner_running = false
}
