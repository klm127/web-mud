package arg

import (
	"flag"
	"fmt"
	"os"

	"github.com/pwsdc/web-mud/shared"
)

// HTTP Server settings
type httpConfig struct {
	shared.HasLogs
	port   *string `env:"http_port"`
	docker *bool   `env:"http_docker"`
}

func (self *httpConfig) setFlags() {
	shared.HasLogsInit(self)
	self.port = flag.String("http_port", "80", "Set the port for the http server to listen on.")
	self.docker = flag.Bool("http_docker", false, "Configures whether the http server is in a docker container using a docker network or should use localhost.")
}

func (self *httpConfig) parseEnv() {

	os_port := os.Getenv("http_port")
	if os_port == "" {
		configWarn("http_port", *self.port)
	}
	self.Log("Port set to: " + *self.port + " from environment.")

	os_docker := os.Getenv("http_docker")
	if os_docker == "" {
		configWarn("http_docker", *self.docker)
	} else {
		if os_docker == "true" {
			*self.docker = true
		} else if os_docker == "false" {
			*self.docker = false
		} else {
			configWarn("http_docker", *self.docker)
		}
	}
	self.Log(fmt.Sprintf("Docker mode set to: %v from environment.", *self.docker))
}

func (self *httpConfig) Port() string {
	return *self.port
}

func (self *httpConfig) DockerMode() bool {
	return *self.docker
}
