package arg

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/pwsdc/web-mud/shared"
)

type socketConfig struct {
	shared.HasLogs
	idleTimeToDisconnect          *int `env:"minutes_to_disconnect_socket"`
	time_to_check_for_disconnects *int `env:"time_to_check_for_idles"`
}

func (sc *socketConfig) setFlags() {
	shared.HasLogsInit(sc)
	sc.idleTimeToDisconnect = flag.Int("socket-dc-time", 1, "Set the idle time to disconnect a socket.")
	sc.time_to_check_for_disconnects = flag.Int("socket-dc-check", 1, "Interval to check connections for inactives.")
}

func (sc *socketConfig) parseEnv() {
	idle_time := os.Getenv("minutes_to_disconnect_socket")
	if idle_time == "" {
		configWarn("minutes_to_disconnect_socket", *sc.idleTimeToDisconnect)
	} else {
		n, err := strconv.Atoi(idle_time)
		if err != nil {
			configWarn("minutes_to_disconnect_socket", idle_time)
		}
		*sc.idleTimeToDisconnect = n
	}
	sc.Log(fmt.Sprintf("Socket idle disconnect time set to %v from environment.", *sc.idleTimeToDisconnect))

	time_check_dcs := os.Getenv("time_to_check_for_idles")
	if idle_time == "" {
		configWarn("time_to_check_for_idles", *sc.idleTimeToDisconnect)
	} else {
		n, err := strconv.Atoi(time_check_dcs)
		if err != nil {
			configWarn("time_to_check_for_idles", idle_time)
		}
		*sc.time_to_check_for_disconnects = n
	}
	sc.Log(fmt.Sprintf("Socket idle check interval set to %v from environment.", *sc.idleTimeToDisconnect))
}

// Gets the time before a socket should be disconnected, in minutes.
func (sc *socketConfig) IdleTimeout() int {
	return *sc.idleTimeToDisconnect
}

// Get the interval that active sockets should be checked to see if they have timed out
func (sc *socketConfig) IdleCheckInterval() int {
	return *sc.time_to_check_for_disconnects
}
