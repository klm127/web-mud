package workers

import (
	"time"

	"github.com/pwsdc/web-mud/world"
)

var room_cleaner_running bool

func StartDirtyRoomClean(interval float64) {
	room_cleaner_running = true
	secs := time.Duration(interval) * time.Second
	for room_cleaner_running {
		world.Rooms.SaveDirty()
		time.Sleep(secs)
	}
}

func StopDirtyRoomClean() {
	room_cleaner_running = false
}
