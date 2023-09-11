package arg

import (
	"flag"
	"os"
	"strconv"

	"github.com/pwsdc/web-mud/shared"
)

type worldConfig struct {
	shared.HasLogs
	startRoom *int64 `env:"start_room"`
}

func (wc *worldConfig) setFlags() {
	shared.HasLogsInit(wc)
	wc.startRoom = flag.Int64("start-room", 0, "Set the ID of the room where new players will start.")
}

func (wc *worldConfig) parseEnv() {
	st_rm := os.Getenv("start_room")
	if st_rm == "" {
		configWarn("start_room", *wc.startRoom)
	} else {
		n, err := strconv.Atoi(st_rm)
		if err != nil {
			configWarn("start_room", st_rm)
		}
		*wc.startRoom = int64(n)
	}
	wc.Logf("Player start room set to %d from environment.", *wc.startRoom)
}

func (wc *worldConfig) StartRoom() int64 {
	return *wc.startRoom
}
