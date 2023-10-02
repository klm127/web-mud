package workers

import (
	"time"

	"github.com/pwsdc/web-mud/world"
)

var move_processor_running bool

func StartMoveProcesser(interval float64) {
	move_processor_running = true
	secs := time.Duration(interval) * time.Second
	for move_processor_running {
		world.Movements.Process()
		time.Sleep(secs)
	}
}

func StopMoveProcesser() {
	move_processor_running = false
}
