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

func (self *socketConfig) setFlags() {
	shared.HasLogsInit(self)
	self.idleTimeToDisconnect = flag.Int("socket-dc-time", 1, "Set the idle time to disconnect a socket.")
	self.time_to_check_for_disconnects = flag.Int("socket-dc-check", 1, "Interval to check connections for inactives.")
}

func (self *socketConfig) parseEnv() {
	idle_time := os.Getenv("minutes_to_disconnect_socket")
	if idle_time == "" {
		configWarn("minutes_to_disconnect_socket", *self.idleTimeToDisconnect)
	} else {
		n, err := strconv.Atoi(idle_time)
		if err != nil {
			configWarn("minutes_to_disconnect_socket", idle_time)
		}
		*self.idleTimeToDisconnect = n
	}
	self.Log(fmt.Sprintf("Socket idle disconnect time set to %v from environment.", *self.idleTimeToDisconnect))

	time_check_dcs := os.Getenv("time_to_check_for_idles")
	if idle_time == "" {
		configWarn("time_to_check_for_idles", *self.idleTimeToDisconnect)
	} else {
		n, err := strconv.Atoi(time_check_dcs)
		if err != nil {
			configWarn("time_to_check_for_idles", idle_time)
		}
		*self.time_to_check_for_disconnects = n
	}
	self.Log(fmt.Sprintf("Socket idle check interval set to %v from environment.", *self.idleTimeToDisconnect))
}

// Gets the time before a socket should be disconnected, in minutes.
func (self *socketConfig) IdleTimeout() int {
	return *self.idleTimeToDisconnect
}

// Get the interval that active sockets should be checked to see if they have timed out
func (self *socketConfig) IdleCheckInterval() int {
	return *self.time_to_check_for_disconnects
}
