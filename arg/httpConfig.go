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

func (hc *httpConfig) setFlags() {
	shared.HasLogsInit(hc)
	hc.port = flag.String("http_port", "80", "Set the port for the http server to listen on.")
	hc.docker = flag.Bool("http_docker", false, "Configures whether the http server is in a docker container using a docker network or should use localhost.")
}

func (hc *httpConfig) parseEnv() {

	os_port := os.Getenv("http_port")
	if os_port == "" {
		configWarn("http_port", *hc.port)
	}
	hc.Log("Port set to: " + *hc.port + " from environment.")

	os_docker := os.Getenv("http_docker")
	if os_docker == "" {
		configWarn("http_docker", *hc.docker)
	} else {
		if os_docker == "true" {
			*hc.docker = true
		} else if os_docker == "false" {
			*hc.docker = false
		} else {
			configWarn("http_docker", *hc.docker)
		}
	}
	hc.Log(fmt.Sprintf("Docker mode set to: %v from environment.", *hc.docker))
}

func (hc *httpConfig) Port() string {
	return *hc.port
}

func (hc *httpConfig) DockerMode() bool {
	return *hc.docker
}
